package config

import (
	"github.com/cihub/seelog"
	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/ecc"
)

type configDatas struct {
	ChainID string           `json:"chainid"`
	URL     string           `json:"url"`
	Keys    []accountKeyData `json:"keys"`
}

// Config config to force-go
type Config struct {
	ChainID eos.SHA256Bytes
	URL     string
	keys    map[string]accountKey
}

// Data config data for force-go
var Data configDatas

// Cfg parsed Data to cfg
var Cfg Config

// Parse parse cfg from data
func (c *Config) Parse(data *configDatas) {
	// keys
	c.keys = make(map[string]accountKey, 64)
	for _, k := range data.Keys {
		pk, err := ecc.NewPrivateKey(k.PriKey)
		if err != nil {
			panic(err)
		}
		n := accountKey{
			Name:   eos.AN(k.Name),
			PriKey: *pk,
		}
		n.PubKey = n.PriKey.PublicKey()

		c.keys[k.PriKey] = n
		seelog.Debugf("load account key %s -> %s", n.Name, n.PubKey)
	}

	//chainID
	id, err := ToSHA256Bytes(data.ChainID)
	if err != nil {
		panic(err)
	}
	c.ChainID = id
}

// Load load cfg to Cfg
func Load(path string) {
	err := LoadJSONFile(path, &Data)
	if err != nil {
		seelog.Errorf("load %s cfg err by %s", path, err.Error())
		panic(err)
	}
	Cfg.Parse(&Data)
}
