package main

import (
	"fmt"

	"github.com/avalanche-pwn/utd/pkg/srv"
)

func main() {
	fmt.Println("Hello, World!")
	var client srv.ClientSrv
	client.Initialize("127.0.0.1:44453")
	client.InitializeDAP()
	client.Test()
	defer client.Close()
}
