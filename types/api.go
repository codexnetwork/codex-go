package types

import (
	"encoding/json"
	"time"
)

// ClientInterface client interface for all client
type ClientInterface interface {
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
