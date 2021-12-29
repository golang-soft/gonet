package center

import (
	"context"
	"gonet/actor"
	"gonet/server/cmessage"
	"gonet/server/smessage"
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

	this.RegisterCall("ReqServerVerify", func(ctx context.Context, packet *smessage.ReqServerVerify) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)
		if SERVER.m_pServerManager.UniqueAdd(packet.Info) {
			SERVER.m_Log.Debugf("添加失败，请检查是否重复 ")
		}
		SERVER.m_pServerManager.DebugServerList()
	})

	this.RegisterCall("PlayerData", func(ctx context.Context, packet *cmessage.PlayerData) {
		SERVER.m_Log.Debugf("head[%v]", packet)
	})

	this.Actor.Start()
}
