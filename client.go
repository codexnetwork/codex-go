package force

import (
	"github.com/pkg/errors"

	"github.com/fanyang1988/force-go/config"
	eosforceApi "github.com/fanyang1988/force-go/eosforce"
	eosioApi "github.com/fanyang1988/force-go/eosio"
	forceioApi "github.com/fanyang1988/force-go/forceio"
	"github.com/fanyang1988/force-go/types"
)

// Client client to forceio chain
type Client struct {
	api types.ClientInterface
	typ types.ClientType
}

// NewClientAPI create client to chain by typ
func NewClientAPI(typ types.ClientType, cfg *config.ConfigData) (types.ClientInterface, error) {
	var res types.ClientInterface
	switch typ {
	case types.FORCEIO:
		res = &forceioApi.API{}
	case types.EOSIO:
		res = &eosioApi.API{}
	case types.EOSForce:
		res = &eosforceApi.API{}
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
