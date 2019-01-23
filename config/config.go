package config

import (
	"github.com/cihub/seelog"
	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/ecc"
)

// Config config to force-go
type configDatas struct {
	URL  string           `json:"url"`
	Keys []accountKeyData `json:"keys"`
}

// Data config data for force-go
var Data configDatas

var keys map[string]accountKey

func parse() {
	// keys
	for _, k := range Data.Keys {
		pk, err := ecc.NewPrivateKey(k.PriKey)
		if err != nil {
			panic(err)
		}
		n := accountKey{
			Name:   eos.AN(k.Name),
			PriKey: *pk,
		}
		n.PubKey = n.PriKey.PublicKey()

		seelog.Debugf("load account key %s -> %s", n.Name, n.PubKey)
	}
}

// Load load cfg to Cfg
func Load(path string) {
	err := LoadJSONFile(path, &Data)
	if err != nil {
		seelog.Errorf("load %s cfg err by %s", path, err.Error())
		panic(err)
	}
	parse()
}
