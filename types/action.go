package types

import (
	eosio "github.com/eoscanada/eos-go"
	forceio "github.com/eosforce/goforceio"
)

/*
	actToCommit := &eos.Action{
		Account: eos.AN(cfg.GetRelayCfg().RelayContract),
		Name:    eos.ActN("out"),
		Authorization: []eos.PermissionLevel{
			committer.SideAccount,
		},
		ActionData: eos.NewActionData(OutAction{
			Committer: eos.Name(committer.RelayAccount.Actor),
			Num:       num,
			To:        eos.Name(act.From),
			Chain:     act.Chain,
			Contract:  eos.Name("force.token"),
			Quantity:  act.Quantity,
			Memo:      act.Memo,
		}),
	}

*/

type Action struct {
	Account       string
	Name          string
	Authorization []PermissionLevel
	Data          interface{}
}

type PermissionLevel struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

func (a *Action) FromForceio(act *forceio.Action) error {
	a.Account = string(act.Account)
	a.Name = string(act.Name)

	a.Authorization = make([]PermissionLevel, 0, len(act.Authorization))
	for _, au := range act.Authorization {
		a.Authorization = append(a.Authorization, PermissionLevel{
			Actor:      string(au.Actor),
			Permission: string(au.Permission),
		})
	}

	a.Data = act.Data

	return nil
}

func (a *Action) FromEOSIO(act *eosio.Action) error {
	a.Account = string(act.Account)
	a.Name = string(act.Name)

	a.Authorization = make([]PermissionLevel, 0, len(act.Authorization))
	for _, au := range act.Authorization {
		a.Authorization = append(a.Authorization, PermissionLevel{
			Actor:      string(au.Actor),
			Permission: string(au.Permission),
		})
	}

	a.Data = act.Data

	return nil
}

func (a *Action) ToForceio() (*forceio.Action, error) {
	auth := make([]forceio.PermissionLevel, 0, len(a.Authorization))
	for _, au := range a.Authorization {
		auth = append(auth, forceio.PermissionLevel{
			Actor:      forceio.AN(au.Actor),
			Permission: forceio.PermissionName(au.Permission),
		})
	}
	return &forceio.Action{
		Account:       forceio.AN(a.Account),
		Name:          forceio.ActN(a.Name),
		Authorization: auth,
		ActionData:    forceio.NewActionData(a.Data),
	}, nil
}

func (a *Action) ToEOSIO() (*eosio.Action, error) {
	auth := make([]eosio.PermissionLevel, 0, len(a.Authorization))
	for _, au := range a.Authorization {
		auth = append(auth, eosio.PermissionLevel{
			Actor:      eosio.AN(au.Actor),
			Permission: eosio.PermissionName(au.Permission),
		})
	}
	return &eosio.Action{
		Account:       eosio.AN(a.Account),
		Name:          eosio.ActN(a.Name),
		Authorization: auth,
		ActionData:    eosio.NewActionData(a.Data),
	}, nil
}
