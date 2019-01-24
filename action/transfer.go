package action

import (
	eos "github.com/eosforce/goeosforce"
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
func NewTransfer(from, to eos.AccountName, asset eos.Asset, memo string) *eos.Action {
	return &eos.Action{
		Account: TokenAccountName,
		Name:    eos.ActN("transfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: asset,
			Memo:     memo,
		}),
	}
}
