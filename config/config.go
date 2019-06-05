package config

import (
	"github.com/pkg/errors"
)

// ConfigData JSON struct define to config
type ConfigData struct {
	ChainID  string           `json:"chainId"`
	StartNum uint32           `json:"startNum"`
	URL      string           `json:"url"`
	Keys     []accountKeyData `json:"keys"`
	PriKeys  []string         `json:"pri"`
}

// LoadCfgFromFile load cfg from file
func LoadCfgFromFile(path string) (*ConfigData, error) {
	var data ConfigData
	err := LoadJSONFile(path, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "load %s err", path)
	}

	return &data, nil
}
