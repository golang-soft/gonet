package center

import (
	"context"
	"gonet/actor"
	"gonet/server/message"
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

	this.RegisterCall("ReqServerVerify", func(ctx context.Context, packet *message.ReqServerVerify) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)
		SERVER.m_pServerManager.UniqueAdd(packet.Info)
		SERVER.m_pServerManager.DebugServerList()
	})

	this.RegisterCall("PlayerData", func(ctx context.Context, packet *message.PlayerData) {
		SERVER.m_Log.Debugf("head[%v]", packet)
	})

	this.Actor.Start()
}
