package grpcserver

import (
	"context"
	"gonet/actor"
	"gonet/server/cmessage"
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

	this.RegisterCall("W_C_Test", func(ctx context.Context, packet *cmessage.W_C_Test) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)

	})

	this.RegisterCall("PlayerData", func(ctx context.Context, packet *cmessage.PlayerData) {
		SERVER.m_Log.Debugf("head[%v]", packet)
	})

	this.Actor.Start()
}
