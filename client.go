package force

import (
	"github.com/fanyang1988/force-go/common"
	"github.com/fanyang1988/force-go/config"
)

// Client client to forceio chain
type Client struct {
	api *common.API
}

// NewClient create client by config data
func NewClient(cfg *config.Config) (*Client, error) {
	api, err := common.NewAPI(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		api: api,
	}, nil
}

// NewClientFromFile create client by config file
func NewClientFromFile(path string) (*Client, error) {
	cfg, err := config.LoadCfgFromFile(path)
	if err != nil {
		return nil, err
	}
	return NewClient(cfg)
}
