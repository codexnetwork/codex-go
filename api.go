package force

import (
	eos "github.com/eosforce/goeosforce"
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
