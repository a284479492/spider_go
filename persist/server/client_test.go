package main

import (
	"testing"
	"time"

	"../../../crawler/engine"
	"../../../crawler/model"
	"../../config"
	"../../rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const (
		host  = ":1234"
		index = "dating_test"
	)
	//start server
	go serveRpc(host, index)
	time.Sleep(time.Second)
	//start client
	conn, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//call server
	item := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/6466817866957127625",
		Id:   "6466817866957127625",
		Type: "zhenai",
		PayLoad: model.Profile{
			Name:       "孤者何惧考莎啦",
			Gender:     "男",
			Age:        45,
			Weight:     266,
			Height:     181,
			Income:     "5001-8000元",
			Marriage:   "未婚",
			Education:  "博士及以上",
			Occupation: "人事/行政",
			Hokou:      "宁波市",
			Xinzuo:     "处女座",
			House:      "租房",
			Car:        "无车",
		},
	}
	var result string
	err = conn.Call(config.ItemSaveRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s, error: %s\n", result, err)
	}
}
