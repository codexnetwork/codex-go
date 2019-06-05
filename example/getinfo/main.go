package main

import (
	"encoding/json"
	"flag"

	"github.com/cihub/seelog"

	gocodex "github.com/codexnetwork/codex-go"
	"github.com/codexnetwork/codex-go/config"
	"github.com/codexnetwork/codex-go/types"
)

var configPath = flag.String("cfg", "../config.json", "config file path")

func main() {
	defer seelog.Flush()
	flag.Parse()

	cfg, err := config.LoadCfgFromFile(*configPath)
	if err != nil {
		seelog.Errorf("load cfg err by %s", err.Error())
		return
	}

	api, err := gocodex.NewClientAPI(types.EOSIO, cfg)
	if err != nil {
		seelog.Errorf("create api err by %s", err.Error())
		return
	}

	info, err := api.GetInfoData()
	if err != nil {
		seelog.Errorf("get api err by %s", err.Error())
		return
	}

	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		seelog.Errorf("MarshalIndent err by %s", err.Error())
		return
	}

	seelog.Infof("res info data : %v", string(data))
}
