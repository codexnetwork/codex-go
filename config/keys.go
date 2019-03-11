package config

import (
	eos "github.com/eosforce/goforceio"
	"github.com/eosforce/goforceio/ecc"
)

// AccountKeyData pub and pri key for account config info
type accountKeyData struct {
	Name   string `json:"name"`
	PriKey string `json:"key"`
}

type accountKey struct {
	Name   eos.AccountName
	PubKey ecc.PublicKey
	PriKey ecc.PrivateKey
}
