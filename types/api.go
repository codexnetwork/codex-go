package types

import (
	"encoding/json"
	"time"

	"github.com/fanyang1988/force-go/config"
	"github.com/pkg/errors"
)

type ClientType uint8

const (
	ClientTypeNil = ClientType(iota)
	FORCEIO
	EOSForce
	Codex
	EOSIO
	ENU     // no support now
	BOS     // no support now
	TLOS    // no support now
	MEETONE // no support now
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

// SwitcherInterface a interface for diff chain type transfer to common
type SwitcherInterface interface {
	Type() ClientType
	NameFromCommon(n string) interface{}
	Checksum256FromCommon(c Checksum256) interface{}
	PushTransactionFullRespToCommon(r interface{}) (*PushTransactionFullResp, error)
	InfoRespToCommon(r interface{}) (*InfoResp, error)
	ActionToCommon(d interface{}) (*Action, error)
	ActionFromCommon(d *Action) (interface{}, error)
	TransactionToCommon(r interface{}) (*TransactionGeneralInfo, error)
	BlockToCommon(r interface{}) (*BlockGeneralInfo, error)
}

// NewSwitcherInterface create SwitcherInterface by typ
func NewSwitcherInterface(typ ClientType) SwitcherInterface {
	switch typ {
	case FORCEIO:
		return &switcher2FORCEIO{}
	case EOSIO:
		return &switcher2EOSIO{}
	case EOSForce:
		return &switcher2EOSForce{}
	default:
		panic(ErrNoSupportChain)
	}
}

var (
	ErrTypeErrToChain = errors.New("ErrTypeErrToChain")
	ErrNoSupportChain = errors.New("ErrNoSupportChain")
)

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
