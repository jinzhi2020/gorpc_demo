package main

import (
	"demo/client_proxy"
	"fmt"
)

func main() {
	client := client_proxy.NewTimeServiceClient("tcp", "localhost:1234")
	var reply string
	_ = client.Now("", &reply)
	fmt.Println(reply)
}
