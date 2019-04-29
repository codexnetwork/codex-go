package force

import (
	"github.com/fanyang1988/force-go/config"
	eosforceapi "github.com/fanyang1988/force-go/eosforce"
	eosioapi "github.com/fanyang1988/force-go/eosio"
	forceioapi "github.com/fanyang1988/force-go/forceio"
	"github.com/fanyang1988/force-go/types"
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

// Client client to forceio chain
type Client struct {
	api types.ClientInterface
	typ ClientType
}

func NewClientAPI(typ ClientType, cfg *config.ConfigData) (types.ClientInterface, error) {
	var res types.ClientInterface
	switch typ {
	case FORCEIO:
		res = &forceioapi.API{}
	case EOSIO:
		res = &eosioapi.API{}
	case EOSForce:
		res = &eosforceapi.API{}
	default:
		return nil, errors.New("unsupported api type")
	}

	err := res.Init(cfg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// NewClient create client by config data
func NewClient(typ ClientType, cfg *config.ConfigData) (types.ClientInterface, error) {
	return NewClientAPI(typ, cfg)
}

// NewClientFromFile create client by config file
func NewClientFromFile(typ ClientType, path string) (types.ClientInterface, error) {
	cfg, err := config.LoadCfgFromFile(path)
	if err != nil {
		return nil, err
	}
	return NewClient(typ, cfg)
}
