package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fanyang1988/force-block-ev/log"
	"github.com/fanyang1988/force-go/p2p"
	"github.com/fanyang1988/force-go/types"
	"go.uber.org/zap"
)

var chainID = flag.String("chain-id", "66b03fd7b1fa2f86afa0bdb408e1261494001b08a3ba16d5093f8d1c3d44f385", "net chainID to connect to")
var showLog = flag.Bool("v", false, "show detail log")
var startNum = flag.Int("num", 1, "start block num to sync")
var p2pAddress = flag.String("p2p", "", "p2p address")

// Wait wait for term signal, then stop the server
func Wait() {
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGQUIT,
		syscall.SIGUSR1)
	<-stopSignalChan
}

type handlerImp struct {
}

func (h *handlerImp) OnBlock(peer string, msg *types.BlockGeneralInfo) error {
	log.Logger().Info("on checked block", zap.String("peer", peer), zap.Uint32("num", msg.BlockNum))
	return nil
}
func (h *handlerImp) OnGoAway(peer string, reason uint8, nodeID types.Checksum256) error {
	return nil
}

func main() {
	flag.Parse()

	if *showLog {
		log.EnableLogging(false)
	}

	// from 9001 - 9020
	const maxNumListen int = 1
	peers := make([]string, 0, maxNumListen+1)

	if *p2pAddress == "" {
		for i := 0; i < maxNumListen; i++ {
			peers = append(peers, fmt.Sprintf("127.0.0.1:%d", 8101+i))
		}
	} else {
		peers = append(peers, *p2pAddress)
	}

	var stratBlock *p2p.P2PSyncData
	if *startNum != 0 {
		stratBlock = &p2p.P2PSyncData{
			HeadBlockNum: uint32(*startNum),
		}
		log.Logger().Sugar().Infof("start %v", *stratBlock)
	}

	p2pPeers := p2p.NewP2PClient(types.FORCEIO, p2p.P2PInitParams{
		Name:       "testNode",
		ClientID:   *chainID,
		StartBlock: stratBlock,
		Peers:      peers[:],
		Logger:     log.Logger(),
	})

	p2pPeers.RegHandler(&handlerImp{})
	err := p2pPeers.Start()

	if err != nil {
		log.Logger().Error("start err", zap.Error(err))
	}

	Wait()

	err = p2pPeers.CloseConnection()
	if err != nil {
		log.Logger().Error("start err", zap.Error(err))
	}
}
