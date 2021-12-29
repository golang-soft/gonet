package worlddb

import (
	"context"
	"database/sql"
	"gonet/actor"
	"gonet/server/cmessage"
)

type (
	EventProcess struct {
		actor.Actor
		m_db *sql.DB
	}

	IEventProcess interface {
		actor.IActor
	}
)

func (this *EventProcess) Init() {
	this.Actor.Init()
	this.m_db = SERVER.GetDB()

	this.RegisterCall("W_C_Test", func(ctx context.Context, packet *cmessage.W_C_Test) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)
	})
	this.Actor.Start()
}
