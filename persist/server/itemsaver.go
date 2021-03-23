package main

import (
	"flag"
	"fmt"

	"crawler/config"
	"crawler/persist"
	"crawler/rpcsupport"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("listen_port", 0, "The port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	err := serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	fmt.Println("Itemsaver start serve and listen on ", *port)
	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client: client,
		Index:  index,
	})
}
