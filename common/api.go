package common

import (
	eos "github.com/eosforce/goeosforce"
	"github.com/fanyang1988/force-go/config"
)

// API client api to forceio chain
type API struct {
	*eos.API
	Cfg config.Config
}

// NewAPI create api to apiURL
func NewAPI(apiURL string) *API {
	api := eos.New(apiURL)
	return &API{
		api,
		config.Cfg,
	}
}
