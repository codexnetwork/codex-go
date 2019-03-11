package force

import (
	"github.com/eosforce/goforceio"
	"github.com/fanyang1988/force-go/common"
)

// PushAction push action to chain
func (c *Client) PushAction(action *eos.Action) (*eos.PushTransactionFullResp, error) {
	return common.PushActions(c.api, action)
}

// PushActions push actions to chain
func (c *Client) PushActions(actions ...*eos.Action) (*eos.PushTransactionFullResp, error) {
	return common.PushActions(c.api, actions...)
}

// GetInfo get info
func (c *Client) GetInfo() (out *eos.InfoResp, err error) {
	return c.api.GetInfo()
}

// GetBlockByID get block by id
func (c *Client) GetBlockByID(id string) (out *eos.BlockResp, err error) {
	return c.api.GetBlockByID(id)
}

// GetBlockByNum get block by num
func (c *Client) GetBlockByNum(num uint32) (out *eos.BlockResp, err error) {
	return c.api.GetBlockByNum(num)
}

// GetTableRows get table
func (c *Client) GetTableRows(params eos.GetTableRowsRequest) (out *eos.GetTableRowsResp, err error) {
	return c.api.GetTableRows(params)
}
