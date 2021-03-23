package main

import (
	"flag"
	"fmt"
	"log"

	"crawler/rpcsupport"
	"crawler/worker"
)

var port = flag.Int("listen_port", 0, "The port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Printf("Worker start serve and listen on %d\n", *port)
	err := rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{})
	if err != nil {
		log.Printf("Worker start serve fail on %d: %v\n", *port, err)
	}
}
