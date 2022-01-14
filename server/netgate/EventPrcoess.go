package netgate

import (
	"context"
	"gonet/actor"
	"gonet/server/cmessage"
	"gonet/server/rpc"
	"log"
)

type (
	EventProcess struct {
		actor.Actor
	}

	IEventProcess interface {
		actor.IActor
	}
)

func (this *EventProcess) Init() {
	this.Actor.Init()

	this.RegisterCall("A_G_Account_Login", func(ctx context.Context, clusterInfo rpc.PlayerClusterInfo) {
		head := this.GetRpcHead(ctx)
		SERVER.GetPlayerMgr().SendMsg(head, "ADD_ACCOUNT", head.SocketId, clusterInfo)
	})
	this.RegisterCall("W_C_Test", func(ctx context.Context, packet *cmessage.W_C_Test) {
		log.Printf("W_C_Test")
	})
	this.Actor.Start()
}
