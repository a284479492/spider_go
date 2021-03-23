package client

import (
	"log"

	"crawler/config"
	"crawler/rpcsupport"
	"crawler/types"
)

func ItemSaver(host string) (chan types.Item, error) {
	rpcClient, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	out := make(chan types.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item saver: get item #%d:%v\n", itemCount, item)
			itemCount++

			var result string
			err = rpcClient.Call(config.ItemSaveRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("ItemSaver ERROR saving item %v: %v\n", item, err)
			}
		}
	}()
	return out, nil
}
