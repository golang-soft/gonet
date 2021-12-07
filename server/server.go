package main

import (
	"fmt"
	"gonet/base"
	"gonet/server/account"
	"gonet/server/login"
	"gonet/server/netgate"
	"gonet/server/world"
	"gonet/server/worlddb"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	args := os.Args
	if args[1] == "account" {
		account.SERVER.Init()
	} else if args[1] == "netgate" {
		netgate.SERVER.Init()
	} else if args[1] == "world" {
		world.SERVER.Init()
	} else if args[1] == "login" {
		login.SERVER.Init()
	} else if args[1] == "worlddb" {
		worlddb.SERVER.Init()
	}

	base.SEVERNAME = args[1]

	InitMgr(args[1])

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c

	ExitMgr(args[1])

	fmt.Printf("server【%s】 exit ------- signal:[%v]", args[1], s)
}
