package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"crawler/engine"
	itemSaver "crawler/persist/client"
	"crawler/rpcsupport"
	"crawler/scheduler"
	"crawler/types"
	worker "crawler/worker/client"
	"crawler/zhenai/parser"
)

var (
	itemSaverHost = flag.String("itemSaver_host", ":30000", "Host of item saver")
	workerHost    = flag.String("worker_host", ":40000", "Host of worker (comma separated)")
)

func main() {
	flag.Parse()
	if *itemSaverHost == "" {
		log.Fatal("must specify a host for item saver")
	}
	if *workerHost == "" {
		log.Fatal("must specify a or more host for worker")
	}
	workerHosts := strings.Split(*workerHost, ",")

	shanghaiurl := "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai"
	itemSaver, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	// type ProcessChan chan engine.Processor
	pools := createProcessorPool(workerHosts)
	processor := worker.CreateProcessor(pools)

	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueuedScheduler{},
		WorkerCount:    10,
		ItemChan:       itemSaver,
		ProcessRequest: processor,
	}

	e.Run(types.Request{
		URL:    shanghaiurl,
		Parser: types.NewFuncParser(parser.ParseCity, "ParseCity"),
	})
}

func createProcessorPool(hosts []string) chan *rpc.Client {
	clients := []*rpc.Client{}
	for _, host := range hosts {
		rpcClient, err := rpcsupport.NewClient(host)
		if err != nil {
			log.Printf("Failed to create rcp connection with client:%v\n", host)
		}
		clients = append(clients, rpcClient)
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
