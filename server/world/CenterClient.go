package world

import (
	"context"
	"gonet/actor"
	"gonet/network"
	"gonet/server/cmessage"
	"gonet/server/glogger"
)

type (
	CenterProcess struct {
		actor.Actor

		Client *network.ClientSocket
	}

	ICenterProcess interface {
		actor.IActor
	}
)

func (this *CenterProcess) Init() {
	this.Actor.Init()

	this.RegisterCall("W_C_Test", func(ctx context.Context, packet *cmessage.W_C_Test) {
		head := this.GetRpcHead(ctx)
		glogger.M_Log.Debugf("head[%v]", head)
	})

	this.Actor.Start()
}
