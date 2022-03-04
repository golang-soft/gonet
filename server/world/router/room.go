package router

import (
	"gonet/base/logger"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/world/datafnc"
	"gonet/server/world/gamedata"
	"gonet/server/world/param"
	"gonet/server/world/sender"
	"gonet/server/world/socket"
)

func Includes(datas []int, role int32) bool {
	for _, data := range datas {
		if role == int32(data) {
			return true
		}
	}
	return false
}

func ROOM_ALL(socket socket.Socket, param param.RoomParam) {
	var mode = common.Mode.Room
	funcName := USER_EVENT.ROOM.ROOM_ALL
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		User:     param.User,
	})
}

func ROOM_CTEATE(socket socket.Socket, param param.RoomParam) {
	//data: { user: string, battle_id: number, name: string, pwd: string, hero_id: any, role: number }
	var user = param.User
	var battle_id = param.BattleId
	var name = param.Name
	var pwd = param.Pwd
	var hero_id = param.HeroId
	var role = param.Role

	if socket.Data.RoomId != 0 {
		return
	}

	if name == "" {
		logger.Errorf("error error_name")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Error_name)})
		return
	}
	if !Includes(datafnc.Hero_Role_Type, role) {
		logger.Errorf("error role")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_hero)})
		return
	}

	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_CTEATE
	sender.AddRoomTask(data.RoomData{
		Mode:      mode,
		FuncName:  funcName,
		User:      user,
		Battle_id: battle_id,
		Name:      name,
		Pwd:       pwd,
		Hero_id:   hero_id,
		Role:      role,
	})
}

func ROOM_JOIN(socket socket.Socket, param param.RoomParam) {

	//data: { room_id: number, user: string, hero_id: any, pwd: string, role: number }
	var user = param.User
	var room_id = param.RoomId
	var pwd = param.Pwd
	var role = param.Role
	var hero_id = param.HeroId

	if socket.Data.RoomId != 0 {
		logger.Errorf("has room")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Already_in)})
		return
	}

	//const { user, room_id, hero_id, pwd, role } = data
	if !Includes(datafnc.Hero_Role_Type, role) {
		logger.Errorf("error role")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_hero)})
		return
	}

	var roomData = gamedata.RoomCtrl.GetRoomById(room_id)
	if roomData == nil || len(roomData.List) >= datafnc.Camp_Player_Amount {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.No_room)})
		return
	}

	for _, data := range roomData.List {
		if data.User == user {
			socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Already_in)})
			return
		}
	}

	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_JOIN
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		User:     user,
		RoomId:   room_id,
		Hero_id:  hero_id,
		Pwd:      pwd,
		Role:     role,
	})
}

func ROOM_JOIN_QUICK(socket socket.Socket, param param.RoomParam) {

	//data: { user: string, hero_id: any, pwd: string, role: number }
	var user = param.User
	var hero_id = param.HeroId
	//var pwd = param.Pwd

	var role = param.Role

	if socket.Data.RoomId != 0 {
		logger.Errorf("has room")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Already_in)})
		return
	}
	if !Includes(datafnc.Hero_Role_Type, role) {
		logger.Errorf("error role")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_hero)})
		return
	}
	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_JOIN_QUICK
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		User:     user,
		Hero_id:  hero_id,
		Role:     role,
	})
}

func ROOM_LEAVE(socket socket.Socket, param param.RoomParam) {

	//data: { room_id: number, user: string }
	var room_id = param.RoomId
	var user = param.User

	if socket.Data.RoomId == 0 {
		logger.Errorf("no room")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.No_room)})
		return
	}
	var roomData = gamedata.RoomCtrl.GetRoomById(room_id)
	if roomData == nil {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.No_room)})
		return
	}

	var inRoom = false
	for _, data := range roomData.List {
		if data.User == user {
			inRoom = true
			break
		}
	}
	if !inRoom {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.No_room)})
		return
	}

	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_LEAVE
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		RoomId:   room_id,
		User:     user,
	})
}

func ROOM_MATCH(socket socket.Socket, param param.RoomParam) {
	//data: { room_id: number, user: string }
	var room_id = param.RoomId
	var user = param.User
	//房间匹配
	//const { user, room_id } = data
	var roomData = gamedata.RoomCtrl.GetRoomById(room_id)
	if socket.Data.RoomId == 0 || roomData == nil || roomData.OwnerId != user {
		logger.Errorf("no room")
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_owner)})
		return
	}
	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_MATCH
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		RoomId:   room_id,
		User:     user,
	})
}

func ROOM_KICK_OFF(socket socket.Socket, param param.RoomParam) {
	//data: { room_id: number, user: string }
	var room_id = param.RoomId
	var user = param.User
	var to = param.To

	//房间踢人
	var roomData = gamedata.RoomCtrl.GetRoomById(room_id)
	if socket.Data.RoomId == 0 || roomData == nil || roomData.OwnerId != user {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_owner)})
		return
	}
	var inRoom = false
	for _, data := range roomData.List {
		if data.User == to {
			inRoom = true
			break
		}
	}
	if !inRoom {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_owner)})
		return
	}

	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_KICK_OFF
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		RoomId:   room_id,
		User:     user,
		To:       to,
	})
}

func ROOM_RENAME(socket socket.Socket, param param.RoomParam) {
	//async function (data: { room_id: number, user: string, name: string }) {
	//const { room_id, user, name } = data
	var room_id = param.RoomId
	var user = param.User
	var name = param.Name

	var roomData = gamedata.RoomCtrl.GetRoomById(room_id)

	if socket.Data.RoomId == 0 || roomData == nil || roomData.OwnerId != user {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_owner)})
		return
	}
	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_RENAME
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		RoomId:   room_id,
		User:     user,
		Name:     name,
	})

}

func ROOM_CHANGE_PWD(socket socket.Socket, param param.RoomParam) {
	//data: { room_id: number, user: string, pwd: string }
	var room_id = param.RoomId
	var user = param.User
	var pwd = param.Pwd

	if socket.Data.RoomId == 0 || int64(room_id) != socket.Data.RoomId {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_owner)})
		return
	}

	var roomData = gamedata.RoomCtrl.GetRoomById(room_id)

	if roomData == nil || roomData.OwnerId != user {
		socket.Emit(USER_EVENT.ERROR.ERROR, &cmessage.ErrorResp{Code: int32(common.ErrorCode.Not_owner)})
		return
	}
	var mode = common.Mode.Room
	var funcName = USER_EVENT.ROOM.ROOM_CHANGE_PWD
	sender.AddRoomTask(data.RoomData{
		Mode:     mode,
		FuncName: funcName,
		RoomId:   room_id,
		User:     user,
		Pwd:      pwd,
	})

}
func HandleRoom(socket socket.Socket) {

	//全部房间
	socket.On(USER_EVENT.ROOM.ROOM_ALL, ROOM_ALL)

	//房间创建
	socket.On(USER_EVENT.ROOM.ROOM_CTEATE, ROOM_CTEATE)

	//房间加入
	socket.On(USER_EVENT.ROOM.ROOM_JOIN, ROOM_JOIN)

	//快速加入
	socket.On(USER_EVENT.ROOM.ROOM_JOIN_QUICK, ROOM_JOIN_QUICK)

	//房间离开
	socket.On(USER_EVENT.ROOM.ROOM_LEAVE, ROOM_LEAVE)

	//房间匹配
	socket.On(USER_EVENT.ROOM.ROOM_MATCH, ROOM_MATCH)

	//房间踢人
	socket.On(USER_EVENT.ROOM.ROOM_KICK_OFF, ROOM_KICK_OFF)
	//房间更名
	socket.On(USER_EVENT.ROOM.ROOM_RENAME, ROOM_RENAME)

	//房间修改密码
	socket.On(USER_EVENT.ROOM.ROOM_CHANGE_PWD, ROOM_CHANGE_PWD)
}
