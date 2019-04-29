package types

import (
	"encoding/json"
	"time"

	eosio "github.com/eoscanada/eos-go"
	eosforce "github.com/eosforce/goeosforce"
	forceio "github.com/eosforce/goforceio"
	"github.com/fanyang1988/force-go/config"
)

// ClientInterface client interface for all client
type ClientInterface interface {
	Init(cfg *config.ConfigData) error
	PushActions(actions ...*Action) (*PushTransactionFullResp, error)
	GetInfoData() (*InfoResp, error)
	GetBlockDataByID(id string) (*BlockGeneralInfo, error)
	GetBlockDataByNum(num uint32) (*BlockGeneralInfo, error)
	Name(n string) interface{}
}

// PushTransactionFullResp
type PushTransactionFullResp struct {
	StatusCode     string
	TransactionID  string `json:"transaction_id"`
	ProcessedDatas []byte `json:"processed"` // WARN: is an `fc::variant` in server..
	BlockID        string `json:"block_id"`
	BlockNum       uint32 `json:"block_num"`
}

func (p *PushTransactionFullResp) FillProcessedDatas(data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	p.ProcessedDatas = d
	return nil
}

func (p *PushTransactionFullResp) FromForceio(rsp *forceio.PushTransactionFullResp) error {
	p.StatusCode = rsp.StatusCode
	p.TransactionID = rsp.TransactionID
	p.BlockID = rsp.BlockID
	p.BlockNum = rsp.BlockNum
	return p.FillProcessedDatas(rsp.Processed)
}

func (p *PushTransactionFullResp) FromEOSIO(rsp *eosio.PushTransactionFullResp) error {
	p.StatusCode = rsp.StatusCode
	p.TransactionID = rsp.TransactionID
	p.BlockID = rsp.BlockID
	p.BlockNum = rsp.BlockNum
	return p.FillProcessedDatas(rsp.Processed)
}

func (p *PushTransactionFullResp) FromEOSForce(rsp *eosforce.PushTransactionFullResp) error {
	p.StatusCode = rsp.StatusCode
	p.TransactionID = rsp.TransactionID
	p.BlockID = rsp.BlockID
	p.BlockNum = rsp.BlockNum
	return p.FillProcessedDatas(rsp.Processed)
}

type InfoResp struct {
	ServerVersion            string      `json:"server_version"`
	ChainID                  Checksum256 `json:"chain_id"`
	HeadBlockNum             uint32      `json:"head_block_num"`
	LastIrreversibleBlockNum uint32      `json:"last_irreversible_block_num"`
	LastIrreversibleBlockID  Checksum256 `json:"last_irreversible_block_id"`
	HeadBlockID              Checksum256 `json:"head_block_id"`
	HeadBlockTime            time.Time   `json:"head_block_time"`
	HeadBlockProducer        string      `json:"head_block_producer"`
	VirtualBlockCPULimit     int64       `json:"virtual_block_cpu_limit"`
	VirtualBlockNetLimit     int64       `json:"virtual_block_net_limit"`
	BlockCPULimit            int64       `json:"block_cpu_limit"`
	BlockNetLimit            int64       `json:"block_net_limit"`
	ServerVersionString      string      `json:"server_version_string"`
}

func (i *InfoResp) FromForceio(info *forceio.InfoResp) error {
	i.ServerVersion = info.ServerVersion
	i.ChainID = Checksum256(info.ChainID)
	i.HeadBlockNum = info.HeadBlockNum
	i.LastIrreversibleBlockNum = info.LastIrreversibleBlockNum
	i.LastIrreversibleBlockID = Checksum256(info.LastIrreversibleBlockID)
	i.HeadBlockID = Checksum256(info.HeadBlockID)
	i.HeadBlockTime = info.HeadBlockTime.Time
	i.HeadBlockProducer = string(info.HeadBlockProducer)
	i.VirtualBlockCPULimit = int64(info.VirtualBlockCPULimit)
	i.VirtualBlockNetLimit = int64(info.VirtualBlockNetLimit)
	i.BlockCPULimit = int64(info.BlockCPULimit)
	i.BlockNetLimit = int64(info.BlockNetLimit)
	i.ServerVersionString = info.ServerVersionString

	return nil
}

func (i *InfoResp) FromEOSIO(info *eosio.InfoResp) error {
	i.ServerVersion = info.ServerVersion
	i.ChainID = Checksum256(info.ChainID)
	i.HeadBlockNum = info.HeadBlockNum
	i.LastIrreversibleBlockNum = info.LastIrreversibleBlockNum
	i.LastIrreversibleBlockID = Checksum256(info.LastIrreversibleBlockID)
	i.HeadBlockID = Checksum256(info.HeadBlockID)
	i.HeadBlockTime = info.HeadBlockTime.Time
	i.HeadBlockProducer = string(info.HeadBlockProducer)
	i.VirtualBlockCPULimit = int64(info.VirtualBlockCPULimit)
	i.VirtualBlockNetLimit = int64(info.VirtualBlockNetLimit)
	i.BlockCPULimit = int64(info.BlockCPULimit)
	i.BlockNetLimit = int64(info.BlockNetLimit)
	i.ServerVersionString = info.ServerVersionString

	return nil
}

func (i *InfoResp) FromEOSForce(info *eosforce.InfoResp) error {
	i.ServerVersion = info.ServerVersion
	i.ChainID = Checksum256(info.ChainID)
	i.HeadBlockNum = info.HeadBlockNum
	i.LastIrreversibleBlockNum = info.LastIrreversibleBlockNum
	i.LastIrreversibleBlockID = Checksum256(info.LastIrreversibleBlockID)
	i.HeadBlockID = Checksum256(info.HeadBlockID)
	i.HeadBlockTime = info.HeadBlockTime.Time
	i.HeadBlockProducer = string(info.HeadBlockProducer)
	i.VirtualBlockCPULimit = int64(info.VirtualBlockCPULimit)
	i.VirtualBlockNetLimit = int64(info.VirtualBlockNetLimit)
	i.BlockCPULimit = int64(info.BlockCPULimit)
	i.BlockNetLimit = int64(info.BlockNetLimit)
	i.ServerVersionString = info.ServerVersionString

	return nil
}
