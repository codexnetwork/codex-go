package force

import (
	"github.com/fanyang1988/force-go/common"
	"github.com/fanyang1988/force-go/config"
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

// Client client to forceio chain
type Client struct {
	api *common.API
	typ ClientType
}

// NewClient create client by config data
func NewClient(typ ClientType, cfg *config.Config) (*Client, error) {
	api, err := common.NewAPI(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		api: api,
		typ: typ,
	}, nil
}

// NewClientFromFile create client by config file
func NewClientFromFile(typ ClientType, path string) (*Client, error) {
	cfg, err := config.LoadCfgFromFile(path)
	if err != nil {
		return nil, err
	}
	return NewClient(typ, cfg)
}
