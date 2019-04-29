package main

import (
	"flag"

	"github.com/eosforce/goforceio/ecc"

	"github.com/cihub/seelog"
	eos "github.com/eosforce/goforceio"
	force "github.com/fanyang1988/force-go"
	"github.com/fanyang1988/force-go/action"
)

var configPath = flag.String("cfg", "../config.json", "confg file path")

func init() {
	ecc.PublicKeyPrefixCompat = "CDX"
}

func main() {
	defer seelog.Flush()
	flag.Parse()

	client, err := force.NewClientFromFile(force.FORCEIO, *configPath)
	if err != nil {
		seelog.Errorf("new client err by %s", err.Error())
		return
	}

	q, err := eos.NewAsset("1000.0000 CDX")
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
