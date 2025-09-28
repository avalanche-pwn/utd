package main

import (
	"fmt"
	"log/slog"

	"github.com/avalanche-pwn/utd/pkg/cmdline"
	"github.com/avalanche-pwn/utd/pkg/srv"
)

func main() {

	var pf cmdline.Flags
	pf.ParseFlags()

	if pf.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	var client srv.ClientSrv
	err := client.Initialize(fmt.Sprintf("%s:%d", pf.Host, pf.Port))
	if err != nil {
		fmt.Println("Couldn't connect")
		return
	}
	defer client.Close()

	client.InitializeDAP()
	client.ServeBlocking()
}
