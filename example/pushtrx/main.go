package main

import (
	"flag"

	"github.com/eosforce/goeosforce/ecc"

	"github.com/cihub/seelog"
	eos "github.com/eosforce/goeosforce"
	force "github.com/fanyang1988/force-go"
)

var configPath = flag.String("cfg", "../config.json", "confg file path")

type transfer struct {
	From     eos.AccountName `json:"from"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}

func init() {
	ecc.PublicKeyPrefixCompat = "FOSC"
}

func main() {
	defer seelog.Flush()
	flag.Parse()

	client, err := force.NewClientFromFile(*configPath)

	q, err := eos.NewAsset("1000.0000 SYS")
	if err != nil {
		seelog.Errorf("asset err by %s", err.Error())
		return
	}

	_, err = client.PushActions(&eos.Action{
		Account: eos.AN("force.token"),
		Name:    eos.ActN("transfer"),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AN("eosforce"), Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(transfer{
			From:     eos.AN("eosforce"),
			To:       eos.AN("testc"),
			Quantity: q,
			Memo:     "test transfer",
		}),
	})

	if err != nil {
		seelog.Errorf("push action err by %s", err.Error())
		return
	}
}
