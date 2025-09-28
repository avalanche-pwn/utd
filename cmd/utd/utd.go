package main

import (
	"fmt"
	"os"

	"github.com/avalanche-pwn/utd/pkg/srv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: utd PORT")
		return
	}
	port := os.Args[1]

	var client srv.ClientSrv
	err := client.Initialize(fmt.Sprintf("127.0.0.1:%s", port))
	if err != nil {
		fmt.Println("Couldn't connect")
		return
	}
	defer client.Close()

	client.InitializeDAP()
	client.Test()
}
