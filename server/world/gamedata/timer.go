package gamedata

import (
	"gonet/actor"
	"gonet/server/common/data"
	glogger2 "gonet/server/glogger"
	"gonet/server/world/sender"
	"gonet/server/world/wcluster"
	"time"
)

type (
	SOnloadTimer struct {
		actor.Actor
	}

	ISOnloadTimer interface {
		actor.IActor
	}
)

var OnloadTimer = &SOnloadTimer{}

func (this *SOnloadTimer) Init() {
	glogger2.M_Log.Debugf("数据库初始化 ...................................")
	this.Actor.Init()
	this.RegisterTimer(100*time.Second, this.OnloadGameCheckTimer) //定时器
	this.Actor.Start()
}

func (this *SOnloadTimer) OnloadGameCheckTimer() {
	glogger2.M_Log.Debugf("触发定时器 ...................................")

	GameCtrl.CheckGame()
	GMatch.CheckMatch()
	GRoom.CheckRoom()

	wcluster.GetCluster().DebugService()

	sender.AddRoomTask(data.RoomData{
		FuncName: "Room_all",
		User:     "111111",
	})
	sender.AddLogTask(data.LogData{
		Mode:   1,
		Userid: "11111111",
		Start:  1,
		Ip:     "111.00.22.33",
		Mac:    "111.00.22.33",
	})
	SaveCtrl.SaveRound(1)
	Consume()
}
