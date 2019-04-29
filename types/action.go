package types

import (
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
	HexData       []byte
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
	a.HexData = act.HexData[:]

	return nil
}
