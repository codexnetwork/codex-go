package eosio

import (
	"encoding/hex"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/pkg/errors"

	"github.com/codexnetwork/codex-go/config"
)

// Config config to codex-go
type Config struct {
	ChainID eos.SHA256Bytes
	URL     string
	Keys    map[string]accountKey
	Prikeys []ecc.PrivateKey
	IsDebug bool
}

type accountKey struct {
	Name   eos.AccountName
	PubKey ecc.PublicKey
	PriKey ecc.PrivateKey
}

// Parse parse cfg from data
func (c *Config) Parse(data *config.ConfigData) error {
	// keys
	c.Keys = make(map[string]accountKey, 64)
	for _, k := range data.Keys {
		pk, err := ecc.NewPrivateKey(k.PriKey)
		if err != nil {
			return errors.Wrapf(err, "parse account pri keys map err")
		}
		n := accountKey{
			Name:   eos.AN(k.Name),
			PriKey: *pk,
		}
		n.PubKey = n.PriKey.PublicKey()

		c.Keys[k.PriKey] = n
		//seelog.Debugf("load account key %s -> %s", n.Name, n.PubKey)
	}

	c.Prikeys = make([]ecc.PrivateKey, 0, len(data.PriKeys)+64)
	for _, k := range data.PriKeys {
		pk, err := ecc.NewPrivateKey(k)
		if err != nil {
			return errors.Wrapf(err, "parse prikey err")
		}
		c.Prikeys = append(c.Prikeys, *pk)
		//seelog.Debugf("load key %s", pk.PublicKey())
	}

	//chainID
	if data.ChainID != "" {
		id, err := ToSHA256Bytes(data.ChainID)
		if err != nil {
			return errors.Wrapf(err, "parse chainid err")
		}
		c.ChainID = id
	} else {
		c.ChainID = nil
	}

	c.URL = data.URL

	return nil
}

// ToSHA256Bytes from string to sha256
func ToSHA256Bytes(in string) (eos.SHA256Bytes, error) {
	if len(in) != 64 {
		return nil, errors.New("should be 64 hexadecimal characters")
	}

	bytes, err := hex.DecodeString(in)
	if err != nil {
		return nil, err
	}

	return eos.SHA256Bytes(bytes), nil
}
