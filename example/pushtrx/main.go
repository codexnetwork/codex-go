package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/cihub/seelog"
	eos "github.com/eosforce/goeosforce"
	"github.com/fanyang1988/force-go/common"
	"github.com/fanyang1988/force-go/config"
)

var configPath = flag.String("cfg", "../config.json", "confg file path")

type transfer struct {
	From     eos.AccountName `json:"from"`
	To       eos.AccountName `json:"to"`
	Quantity eos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}

func main() {
	defer seelog.Flush()
	flag.Parse()

	config.Load(*configPath)

	api := common.NewAPI(config.Data.URL)
	info, err := api.GetInfo()
	if err != nil {
		seelog.Errorf("get api err by %s", err.Error())
		return
	}

	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		seelog.Errorf("MarshalIndent err by %s", err.Error())
		return
	}

	fmt.Println(string(data))

	q, err := eos.NewAsset("1000.0000 SYS")
	if err != nil {
		seelog.Errorf("asset err by %s", err.Error())
		return
	}

	testAction := &eos.Action{
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
	}

	err = common.PushActions(api, testAction)
	if err != nil {
		seelog.Errorf("push action err by %s", err.Error())
		return
	}
}
