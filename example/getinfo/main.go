package main

import (
	"github.com/fanyang1988/force-go/common"

	"encoding/json"
	"fmt"

	"github.com/cihub/seelog"
)

func main() {
	defer seelog.Flush()
	api := common.NewAPI("http://127.0.0.1:8001")
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
