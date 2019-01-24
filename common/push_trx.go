package common

import (
	"strings"
	"time"

	eos "github.com/eosforce/goeosforce"
	"github.com/pkg/errors"
)

func permissionToPermissionLevel(in string) (out eos.PermissionLevel, err error) {
	return eos.NewPermissionLevel(in)
}

func permissionsToPermissionLevels(in []string) (out []eos.PermissionLevel, err error) {
	// loop all parameters
	for _, singleArg := range in {

		// if they specified "account@active,account2", handle that too..
		for _, val := range strings.Split(singleArg, ",") {
			level, err := permissionToPermissionLevel(strings.TrimSpace(val))
			if err != nil {
				return out, err
			}

			out = append(out, level)
		}
	}

	return
}

// PushActions push action to chain
func PushActions(api *API, actions ...*eos.Action) (*eos.PushTransactionFullResp, error) {
	return pushEOSCActionsAndContextFreeActions(api, nil, actions)
}

func pushEOSCActionsAndContextFreeActions(api *API, contextFreeActions []*eos.Action, actions []*eos.Action) (*eos.PushTransactionFullResp, error) {
	for _, act := range contextFreeActions {
		act.Authorization = nil
	}

	opts := &eos.TxOptions{
		ChainID: api.Cfg.ChainID,
	}

	//opts.DelaySecs = 0

	if err := opts.FillFromChain(api.API); err != nil {
		return nil, err
	}

	tx := eos.NewTransaction(actions, opts)
	if len(contextFreeActions) > 0 {
		tx.ContextFreeActions = contextFreeActions
	}

	tx.SetExpiration(30 * time.Second) // TODO use params

	signedTx, packedTx, err := signTransaction(tx, opts.ChainID, api)
	if err != nil {
		return nil, err
	}

	return pushTransaction(signedTx, packedTx, opts.ChainID, api)
}

func signTransaction(tx *eos.Transaction, chainID eos.SHA256Bytes, api *API) (*eos.SignedTransaction, *eos.PackedTransaction, error) {
	signedTx, packedTx, err := api.SignTransaction(tx, chainID, eos.CompressionNone)
	if err != nil {
		return nil, nil, errors.Wrap(err, "signing transaction")
	}
	return signedTx, packedTx, nil
}

// PushTransaction push trx to chain
func pushTransaction(signedTx *eos.SignedTransaction, packedTx *eos.PackedTransaction, chainID eos.SHA256Bytes, api *API) (*eos.PushTransactionFullResp, error) {
	if packedTx == nil {
		return nil, errors.New("A signed transaction is required if you want to broadcast it")
	}

	return api.PushTransaction(packedTx)
}
