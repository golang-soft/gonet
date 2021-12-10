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

	this.RegisterCall("W_C_Test", func(ctx context.Context, packet *message.W_C_Test) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)

	})

	this.RegisterCall("PlayerData", func(ctx context.Context, packet *message.PlayerData) {
		SERVER.m_Log.Debugf("head[%v]", packet)
	})

	this.Actor.Start()
}
