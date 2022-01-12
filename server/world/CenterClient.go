package world

import (
	"context"
	"fmt"
	"gonet/actor"
	"gonet/network"
	"gonet/server/cmessage"
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
		SERVER.M_Log.Debugf("head[%v]", head)
	})

	this.RegisterCall("AttackReq", func(ctx context.Context, packet *cmessage.AttackReq) {
		head := this.GetRpcHead(ctx)

		fmt.Printf("AttackReq  %v", head)
	})

	this.Actor.Start()
}
