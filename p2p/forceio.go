package p2p

import (
	"encoding/hex"
	"fmt"
	"runtime/debug"
	"time"

	eos "github.com/eosforce/goforceio"
	"github.com/eosforce/goforceio/p2p"
	"github.com/fanyang1988/force-go/types"
	"go.uber.org/zap"
)

// p2pForceioClient a manager for peers to diff p2p node
type p2pForceioClient struct {
	*p2pClientImp
}

// NewP2PPeers new p2p peers from cfg
func NewP2PClient4Forceio(name string, chainID string, startBlock *P2PSyncData, peers []string, logger *zap.Logger) *p2pForceioClient {
	p := &p2pForceioClient{
		&p2pClientImp{},
	}

	p.init(name, types.FORCEIO, chainID, startBlock, peers, logger)
	p.setHandlerImp(p)

	cID, err := hex.DecodeString(chainID)
	if err != nil {
		p.logger.Error("decode chain id err", zap.Error(err))
		panic(err)
	}

	var startBlockNum uint32 = 1
	var startBlockId eos.Checksum256
	var startBlockTime time.Time
	var irrBlockNum uint32 = 0
	var irrBlockId eos.Checksum256
	if startBlock != nil {
		startBlockId = eos.Checksum256(startBlock.HeadBlockID)
		startBlockNum = startBlock.HeadBlockNum
		startBlockTime = startBlock.HeadBlockTime
		irrBlockNum = startBlock.LastIrreversibleBlockNum
		irrBlockId = eos.Checksum256(startBlock.LastIrreversibleBlockID)
	}
	for idx, peer := range peers {
		p.logger.Debug("new peer client", zap.Int("idx", idx), zap.String("peer", peer))
		client := p2p.NewClient(
			p2p.NewOutgoingPeer(peer, fmt.Sprintf("%s-%02d", name, idx), &p2p.HandshakeInfo{
				ChainID:                  cID,
				HeadBlockNum:             startBlockNum,
				HeadBlockID:              startBlockId,
				HeadBlockTime:            startBlockTime,
				LastIrreversibleBlockNum: irrBlockNum,
				LastIrreversibleBlockID:  irrBlockId,
			}),
			true,
		)
		client.RegisterHandler(p)
		p.clients = append(p.clients, client)
	}

	return p
}

func (p *p2pForceioClient) handleImp(m p2pClientMsg) {
	peer := m.peer
	pkg, ok := m.msg.(*eos.Packet)
	if !ok {
		p.logger.Error("packet type err")
		return
	}

	for _, h := range p.handlers {
		func(hh types.P2PHandler) {
			defer func() {
				if err := recover(); err != nil {
					p.logger.Error("handler process ev panic",
						zap.String("err", fmt.Sprintf("err:%s", err)),
						zap.String("stack", string(debug.Stack())))
				}
			}()

			var err error
			switch pkg.Type {
			case eos.GoAwayMessageType:
				m, ok := pkg.P2PMessage.(*eos.GoAwayMessage)
				if !ok {
					p.logger.Error("msg type err by go away")
					return
				}
				p.logger.Info("peer goaway",
					zap.String("peer", peer),
					zap.String("reason", m.Reason.String()),
					zap.String("nodeid", m.NodeID.String()))
				err = hh.OnGoAway(peer, uint8(m.Reason), types.Checksum256(m.NodeID))
			case eos.SignedBlockType:
				m, ok := pkg.P2PMessage.(*eos.SignedBlock)
				if !ok {
					p.logger.Error("msg type err by go away")
					return
				}
				p.logger.Debug("on signed block",
					zap.String("peer", peer),
					zap.String("block", m.String()))
				msg, err := p.switcher.BlockToCommon(m)
				if err == nil {
					err = hh.OnBlock(peer, msg)
				} else {
					p.logger.Error("handle msg err", zap.Error(err))
				}
			}

			if err != nil {
				p.logger.Error("handle msg err", zap.Error(err))
			}

		}(h)
	}
}

// Handle handler for p2p clients
func (p *p2pForceioClient) Handle(envelope *p2p.Envelope) {
	p.onMsg(p2pClientMsg{
		peer: envelope.Sender.Address,
		msg:  envelope.Packet,
	})
}
