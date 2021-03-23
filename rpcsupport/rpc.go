package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listenner, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Printf("Accept err: %v\n", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	return jsonrpc.NewClient(conn), nil
}
