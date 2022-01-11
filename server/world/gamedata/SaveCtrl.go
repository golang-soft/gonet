package gamedata

import (
	"gonet/server/common/data"
	"gonet/server/smessage"
	"gonet/server/world/wcluster"
	"time"
)

type (
	SSaveCtrl struct {
	}

	ISSaveCtrl interface {
	}
)

var SaveCtrl = &SSaveCtrl{}

func (this *SSaveCtrl) SaveRound(round int) {
	msg := smessage.SaveRound{
		Round: int32(round),
	}

	wcluster.SendToDB("SaveRound", msg)
}

func (this *SSaveCtrl) SaveLoginLog(message data.LogData) {
	msg := smessage.LoginLogData{
		UserId:    message.Userid,
		Ip:        message.Ip,
		Mac:       message.Mac,
		Lastlogin: message.LastLogin,
	}

	wcluster.SendToDB("LoginLogData", msg)
}

func (this *SSaveCtrl) SaveBattleLog(message data.LogData) {
	msg := smessage.BattleLogData{
		Round: 1,
		Day:   time.Now().Format("2006-01-02"),
	}
	wcluster.SendToDB("BattleLogData", msg)
}

func (this *SSaveCtrl) SaveMatchLog(message data.LogData) {
	msg := smessage.MatchLogData{
		UserId: message.Userid,
		Start:  message.Start,
		End:    message.End,
	}

	wcluster.SendToDB("MatchLogData", msg)
}
