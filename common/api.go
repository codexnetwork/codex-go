package common

import (
	eos "github.com/eosforce/goeosforce"
	"github.com/fanyang1988/force-go/config"
	"github.com/pkg/errors"
)

// API client api to forceio chain
type API struct {
	*eos.API
	Cfg config.Config
}

// NewAPI create api to apiURL
func NewAPI(cfg *config.Config) (*API, error) {
	api := eos.New(cfg.URL)
	res := &API{
		api,
		*cfg,
	}
	err := checkChainID(res)
	if err != nil {
		return nil, err
	}
	res.SetSigner(NewConfigSigner(res))
	return res, nil
}

func checkChainID(api *API) error {
	res, err := api.GetInfo()
	if err != nil {
		return errors.Wrapf(err, "check chain id to get err")
	}

	if api.Cfg.ChainID != nil && api.Cfg.ChainID.String() != res.ChainID.String() {
		return errors.Errorf("chain id diff from cfg by %s <-> %s",
			api.Cfg.ChainID.String(), res.ChainID.String())
	}

	api.Cfg.ChainID = res.ChainID
	return nil
}
