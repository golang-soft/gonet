package netgate

import (
	"gonet/actor"
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

	//this.RegisterCall("A_G_Account_Login", func(ctx context.Context, socketId uint32, clusterInfo rpc.PlayerClusterInfo) {
	//	SERVER.GetPlayerMgr().SendMsg(rpc.RpcHead{}, "ADD_ACCOUNT", socketId, clusterInfo)
	//})
	//this.RegisterCall("W_C_Test", func(ctx context.Context, packet *message.W_C_Test) {
	//	log.Printf("W_C_Test")
	//})
	this.Actor.Start()
}
