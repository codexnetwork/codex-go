package types

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
	Account     string
	Name        string
	Permissions []PermissionLevel
	Data        interface{}
}

type PermissionLevel struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
