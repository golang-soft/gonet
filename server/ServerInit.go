package main

import (
	"gonet/actor"
	"gonet/common"
	"gonet/server/world/chat"
	"gonet/server/world/cmd"
	"gonet/server/world/mail"
	"gonet/server/world/param"
	"gonet/server/world/player"
	"gonet/server/world/social"
	"gonet/server/world/toprank"
	"gonet/server/world/wcluster"
)

func InitMgr(serverName string) {
	//一些共有数据量初始化
	common.Init()
	if serverName == "account" {
	} else if serverName == "netgate" {
		actor.MGR.InitActorHandle(wcluster.GetCluster())
	} else if serverName == "world" {
		cmd.Init()
		param.InitRepository()
		player.MGR.Init()
		chat.MGR.Init()
		mail.MGR.Init()
		toprank.MGR().Init()
		player.SIMPLEMGR.Init()
		social.MGR().Init()
		actor.MGR.InitActorHandle(wcluster.GetCluster())
	} else if serverName == "center" {

	} else if serverName == "world" {

	} else if serverName == "login" {

	} else if serverName == "worlddb" {

	} else if serverName == "zone" {

	} else if serverName == "grpcserver" {

	} else if serverName == "all" {
		//actor.MGR.InitActorHandle(netgate.SERVER.GetCluster())
		cmd.Init()
		param.InitRepository()
		player.MGR.Init()
		chat.MGR.Init()
		mail.MGR.Init()
		toprank.MGR().Init()
		player.SIMPLEMGR.Init()
		social.MGR().Init()
		actor.MGR.InitActorHandle(wcluster.GetCluster())
	}
}

//程序退出后执行
func ExitMgr(serverName string) {
	if serverName == "account" {
	} else if serverName == "netgate" {
	} else if serverName == "world" {
	}
}
