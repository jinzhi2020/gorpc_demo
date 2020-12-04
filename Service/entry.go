package main

import (
	"demo/handler"
	"demo/server_proxy"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)


func main() {
	listener, _ := net.Listen("tcp", ":1234")
	_ = server_proxy.RegisterTimeService(&handler.TimeService{})
	for {
		conn, _ := listener.Accept()
		// 使用 jsonrpc.NewServerCodec 包装 conn
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
