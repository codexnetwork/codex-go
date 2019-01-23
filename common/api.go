package common

import (
	eos "github.com/eosforce/goeosforce"
)

// NewAPI create api to apiURL
func NewAPI(apiURL string) *eos.API {
	api := eos.New(apiURL)
	return api
}
