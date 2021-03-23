package main

import (
	"fmt"
	"testing"
	"time"

	"../../config"
	"../../rpcsupport"
	"../../worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"

	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	var result worker.ParseResult
	args := worker.Request{
		URL: "http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		Parser: worker.SerializedParser{
			FunctionName: "ParseCity",
			Args:         "",
		},
	}
	err = client.Call(config.CrawlServiceRpc, args, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("Get result:", result)
	}
}
