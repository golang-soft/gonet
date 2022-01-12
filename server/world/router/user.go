package router

import (
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/world/gamedata"
	"gonet/server/world/helper"
	"gonet/server/world/param"
	"gonet/server/world/sender"
	"gonet/server/world/socket"
)

func LOGIN(socket socket.Socket, param param.UserParam) {
	//speed: string; direction: string; barrier: string, skillId: number
	var user = param.User
	var heroId = param.HeroId
	//logger.Debug("LOGIN: ", JSON.stringify(data))
	socket.Emit(USER_EVENT.USER.LOGIN, &cmessage.LoginResp{
		User:   user,
		HeroId: heroId,
	})
}

func MOVING(socket socket.Socket, param param.UserParam) {
	//data: { speed: string; direction: string; barrier: string, skillId: number, pos: string }
	var speed = param.Speed
	var direction = param.Direction
	var barrier = param.Barrier
	var skillId = param.SkillId
	//var pos = param.Pos
	var user = param.User
	var round = param.Round
	var x = param.X
	var y = param.Y

	//const { user, round, role, part, roomId } = socket.data
	////TODO:判断比赛时长
	////TODO:倒计时不能行动
	var funcName = USER_EVENT.USER.MOVING
	var userInfo = gamedata.GGame.GetUserById(round, user)
	if userInfo == nil {
		return
	}
	if userInfo.Hp <= 0 {
		return
	}

	var play = helper.USER_BASIC_INFO[userInfo.Itype]
	var roleSpeed = play.Speed
	var realSpeed = speed
	if realSpeed > roleSpeed {
		realSpeed = roleSpeed
	}

	var newPos = gamedata.PositionCtrl.NewUpdatePosition(user, round, data.UserPositionData{
		X:         x,
		Y:         y,
		Speed:     realSpeed,
		Barrier:   barrier,
		Direction: direction,
		SkillId:   skillId,
	})
	gamedata.GvgBattleBroadcastAll(USER_EVENT.USER.MOVING,
		round, &cmessage.MovingResp{
			Speed:         realSpeed,
			ReduceSpeedTs: 0,
			Direction:     direction,
			Barrier:       barrier,
			X:             newPos.X,
			Y:             newPos.Y,
			StopMove:      0,
			StopMoveTs:    0,
			PosUpdateTs:   0,
			User:          user,
			FuncName:      funcName,
		})
}

func JUMP(socket socket.Socket, param param.UserParam) {
	var user = param.User
	var round = param.Round
	gamedata.GvgBattleBroadcastAll(USER_EVENT.USER.JUMP, round, &cmessage.JumpResp{
		User:       user,
		PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_JumpResp), 0),
	})

}
func ATTACK(socket socket.Socket, param param.UserParam) {
	var user = param.User
	var msg = param.Msg
	var itemId = param.ItemId
	var count = param.Count
	var round = param.Round
	var to = param.To
	var skill = param.Skill
	var x = param.X
	var y = param.Y
	var z = param.Z
	var skillId = param.SkillId

	var mode = common.Mode.Skill
	var funcName = USER_EVENT.USER.ATTACK

	sender.AddAttackTask(data.AttackData{
		From:     user,
		To:       to,
		Msg:      msg,
		Mode:     mode,
		FuncName: funcName,
		ItemId:   itemId,
		Count:    count,
		Round:    round,
		Skill:    skill,
		SkillId:  skillId,
		X:        x,
		Y:        y,
		Z:        z,
	})

}
func FLAG(socket socket.Socket, param param.UserParam) {
	var user = param.User
	var round = param.Round
	var itemId = param.ItemId
	var mode = common.Mode.Map
	var funcName = USER_EVENT.USER.FLAG

	sender.AddAttackTask(data.AttackData{
		From:     user,
		Mode:     mode,
		ItemId:   itemId,
		FuncName: funcName,
		Round:    round,
	})
}
func RELIVE(socket socket.Socket, param param.UserParam) {
	var user = param.User
	var msg = param.Msg
	var itemId = param.ItemId
	var count = param.Count
	var round = param.Round
	var to = param.To
	var skill = param.Skill
	var x = param.X
	var y = param.Y
	var z = param.Z
	var skillId = param.SkillId

	var mode = common.Mode.Item
	sender.AddAttackTask(data.AttackData{
		From:    user,
		To:      to,
		Msg:     msg,
		Mode:    mode,
		ItemId:  itemId,
		Count:   count,
		Round:   round,
		Skill:   skill,
		SkillId: skillId,
		X:       x,
		Y:       y,
		Z:       z,
	})

}
func USE_ITEM(socket socket.Socket, param param.UserParam) {
	var user = param.User
	var msg = param.Msg
	var itemId = param.ItemId
	var count = param.Count
	var round = param.Round

	var mode = common.Mode.Item
	var funcName = USER_EVENT.USER.USE_ITEM
	sender.AddAttackTask(data.AttackData{
		From:     user,
		Msg:      msg,
		Mode:     mode,
		FuncName: funcName,
		ItemId:   itemId,
		Count:    count,
		Round:    round,
	})
}

func GAME_END(socket socket.Socket, param param.UserParam) {
	var user = param.User
	var msg = param.Msg
	var itemId = param.ItemId
	var count = param.Count
	var round = param.Round

	var mode = common.Mode.Item
	var funcName = USER_EVENT.USER.USE_ITEM
	sender.AddAttackTask(data.AttackData{
		From:     user,
		Msg:      msg,
		Mode:     mode,
		FuncName: funcName,
		ItemId:   itemId,
		Count:    count,
		Round:    round,
	})
}

func HandleUser(socket socket.Socket) {

	socket.On(USER_EVENT.USER.LOGIN, LOGIN)

	socket.On(USER_EVENT.USER.MOVING, MOVING)

	socket.On(USER_EVENT.USER.JUMP, JUMP)

	socket.On(USER_EVENT.USER.ATTACK, ATTACK)

	socket.On(USER_EVENT.USER.FLAG, FLAG)

	socket.On(USER_EVENT.USER.RELIVE, RELIVE)

	socket.On(USER_EVENT.USER.USE_ITEM, USE_ITEM)

	socket.On(USER_EVENT.GLOBAL.GAME_END, GAME_END)
}
