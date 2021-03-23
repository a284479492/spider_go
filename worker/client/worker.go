package client

import (
	"fmt"
	"net/rpc"

	"crawler/config"
	"crawler/engine"
	"crawler/types"
	"crawler/worker"
)

func CreateProcessor(clients chan *rpc.Client) engine.Processor {
	return func(req types.Request) (types.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var result worker.ParseResult
		c := <-clients
		err := c.Call(config.CrawlServiceRpc, sReq, &result)
		if err != nil {
			fmt.Printf("Failed with request :%v\n", req)
			return types.ParseResult{}, err
		}
		return worker.DeserializeParseResult(result), nil
	}
}
