package action

import (
	eos "github.com/eosforce/goforceio"
	"github.com/fanyang1988/force-go/types"
)

// TokenAccountName account name for token contact to chain
var TokenAccountName = eos.AN("force.token")

// Transfer action data for force.token::transfer action
type Transfer struct {
	From     eos.AccountName `json:"from"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}

// NewTransfer new transfer action by TokenAccountName contract
func NewTransfer(from, to eos.AccountName, asset eos.Asset, memo string) *types.Action {
	return &types.Action{
		Account: string(TokenAccountName),
		Name:    "transfer",
		Authorization: []types.PermissionLevel{
			{Actor: string(from), Permission: "active"},
		},
		Data: Transfer{
			From:     from,
			To:       to,
			Quantity: asset,
			Memo:     memo,
		},
	}
}
