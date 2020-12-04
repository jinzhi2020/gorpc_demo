package client_proxy

import (
	"demo/handler"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type TimeServiceStub struct {
	*rpc.Client
}

func NewTimeServiceClient(protocol, address string) TimeServiceStub {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		panic("connect error!")
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return TimeServiceStub{client}
}

func (c *TimeServiceStub)Now(request string, reply *string) error {
	err := c.Call(handler.TimeServiceName + ".Now", request, reply)
	if err != nil {
		return err
	}
	return nil
}
