package config

import (
	"github.com/cihub/seelog"
	eos "github.com/eosforce/goforceio"
	"github.com/eosforce/goforceio/ecc"
	"github.com/pkg/errors"
)

// ConfigData JSON struct define to config
type ConfigData struct {
	ChainID string           `json:"chainid"`
	URL     string           `json:"url"`
	Keys    []accountKeyData `json:"keys"`
	PriKeys []string         `json:"pri"`
}

// Config config to force-go
type Config struct {
	ChainID eos.SHA256Bytes
	URL     string
	Keys    map[string]accountKey
	Prikeys []ecc.PrivateKey
	IsDebug bool
}

// LoadCfgFromFile load cfg from file
func LoadCfgFromFile(path string) (*Config, error) {
	var data ConfigData
	err := LoadJSONFile(path, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "load %s err", path)
	}

	res := &Config{
		URL: data.URL,
	}

	err = res.Parse(&data)
	if err != nil {
		return nil, errors.Wrapf(err, "parse %s err", path)
	}

	return res, nil
}

// Parse parse cfg from data
func (c *Config) Parse(data *ConfigData) error {
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
		seelog.Debugf("load account key %s -> %s", n.Name, n.PubKey)
	}

	c.Prikeys = make([]ecc.PrivateKey, 0, len(data.PriKeys)+64)
	for _, k := range data.PriKeys {
		pk, err := ecc.NewPrivateKey(k)
		if err != nil {
			return errors.Wrapf(err, "parse prikey err")
		}
		c.Prikeys = append(c.Prikeys, *pk)
		seelog.Debugf("load key %s", pk.PublicKey())
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
