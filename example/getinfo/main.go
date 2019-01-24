package main

import (
	"flag"

	"github.com/fanyang1988/force-go/config"

	"github.com/fanyang1988/force-go/common"

	"encoding/json"
	"fmt"

	"github.com/cihub/seelog"
)

var configPath = flag.String("cfg", "../config.json", "confg file path")

func main() {
	defer seelog.Flush()
	flag.Parse()

	cfg, err := config.LoadCfgFromFile(*configPath)
	if err != nil {
		seelog.Errorf("load cfg err by %s", err.Error())
		return
	}

	api := common.NewAPI(cfg)
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
}
