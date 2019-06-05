package main

import (
	"flag"

	"github.com/eosforce/goforceio/ecc"
	"github.com/cihub/seelog"

	gocodex "github.com/codexnetwork/codex-go"
	"github.com/codexnetwork/codex-go/types"
)

var configPath = flag.String("cfg", "../config.json", "confg file path")

func init() {
	ecc.PublicKeyPrefixCompat = "CDX"
}

func main() {
	defer seelog.Flush()
	flag.Parse()

	client, err := gocodex.NewClientFromFile(types.FORCEIO, *configPath)
	if err != nil {
		seelog.Errorf("new client err by %s", err.Error())
		return
	}

	b, err := client.GetBlockDataByNum(1)
	if err != nil {
		seelog.Errorf("err by %s", err.Error())
		return
	}

	seelog.Infof("get block %v", b)

	b2, err := client.GetBlockDataByID(b.ID.String())
	if err != nil {
		seelog.Errorf("err by %s", err.Error())
		return
	}

	seelog.Infof("get block %v", b2)
}
