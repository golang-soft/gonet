package worlddb

import (
	"context"
	"database/sql"
	"gonet/actor"
	"gonet/server/cmessage"
	"gonet/server/smessage"
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

	this.RegisterCall("SaveRound", func(ctx context.Context, packet *smessage.SaveRound) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)
		SERVER.m_pDataWriterMgr.SaveRound(int(packet.Round))
	})

	this.RegisterCall("LoginLogData", func(ctx context.Context, packet *smessage.LoginLogData) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)

	})

	this.RegisterCall("BattleLogData", func(ctx context.Context, packet *smessage.BattleLogData) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)

	})

	this.RegisterCall("MatchLogData", func(ctx context.Context, packet *smessage.MatchLogData) {
		head := this.GetRpcHead(ctx)
		SERVER.m_Log.Debugf("head[%v]", head)

	})

	this.Actor.Start()
}
