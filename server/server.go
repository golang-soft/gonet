package main

import (
	"fmt"
	"gonet/base"
	"gonet/server/account"
	"gonet/server/center"
	"gonet/server/grpcserver"
	"gonet/server/login"
	"gonet/server/netgate"
	"gonet/server/world"
	"gonet/server/worlddb"
	"gonet/server/zone"
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
	} else if args[1] == "center" {
		center.SERVER.Init()
	} else if args[1] == "zone" {
		zone.SERVER.Init()
	} else if args[1] == "grpcserver" {
		grpcserver.SERVER.Init()
	} else {
		center.SERVER.Init()
		netgate.SERVER.Init()
		account.SERVER.Init()
		world.SERVER.Init()
		worlddb.SERVER.Init()
		zone.SERVER.Init()
		grpcserver.SERVER.Init()
		login.SERVER.Init()
	}

	base.SEVERNAME = args[1]

	InitMgr(args[1])

	startServer(args[1])

	ExitMgr(args[1])
}

func startServer(args string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c

	fmt.Printf("server【%s】 exit ------- signal:[%v]", args, s)
}
