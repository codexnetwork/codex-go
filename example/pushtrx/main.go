package main

import (
	"flag"

	"github.com/eosforce/goeosforce/ecc"

	"github.com/cihub/seelog"
	eos "github.com/eosforce/goeosforce"
	force "github.com/fanyang1988/force-go"
	"github.com/fanyang1988/force-go/action"
)

var configPath = flag.String("cfg", "../config.json", "confg file path")

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

	_, err = client.PushActions(
		action.NewTransfer(eos.AN("eosforce"), eos.AN("testc"), q, "for test"))

	if err != nil {
		seelog.Errorf("push action err by %s", err.Error())
		return
	}
}
