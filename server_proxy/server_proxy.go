package server_proxy

import (
	"demo/handler"
	"net/rpc"
)

type TimeServicer interface {
	Now(request string, reply *string) error
}

func RegisterTimeService(srv TimeServicer) error {
	return 	rpc.RegisterName(handler.TimeServiceName, srv)
}
