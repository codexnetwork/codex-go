package force

import (
	"github.com/fanyang1988/force-go/config"
	eosforceapi "github.com/fanyang1988/force-go/eosforce"
	eosioapi "github.com/fanyang1988/force-go/eosio"
	forceioapi "github.com/fanyang1988/force-go/forceio"
	"github.com/fanyang1988/force-go/types"
	"github.com/pkg/errors"
)

// Client client to forceio chain
type Client struct {
	api types.ClientInterface
	typ types.ClientType
}

func NewClientAPI(typ types.ClientType, cfg *config.ConfigData) (types.ClientInterface, error) {
	var res types.ClientInterface
	switch typ {
	case types.FORCEIO:
		res = &forceioapi.API{}
	case types.EOSIO:
		res = &eosioapi.API{}
	case types.EOSForce:
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
func NewClient(typ types.ClientType, cfg *config.ConfigData) (types.ClientInterface, error) {
	return NewClientAPI(typ, cfg)
}

// NewClientFromFile create client by config file
func NewClientFromFile(typ types.ClientType, path string) (types.ClientInterface, error) {
	cfg, err := config.LoadCfgFromFile(path)
	if err != nil {
		return nil, err
	}
	return NewClient(typ, cfg)
}
