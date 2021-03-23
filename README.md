1. 启动web服务
    `go run mockserver/main.go`
2. 启动itemSaver服务
    `go run persist/server/itemsaver.go -listen_port 30000`
3. 启动worker服务
    `go run worker/server/worker.go -listen_port 40000`
4. 启动main
    `go run main/main.go -itemSaver_host ":30000" -worker_host ":40000"`