package gamedata

import (
	"fmt"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/world/datafnc"
)

//房间
type (
	SRoomCtrl struct {
	}

	ISRoomCtrl interface {
	}
)

var (
	RoomCtrl *SRoomCtrl = &SRoomCtrl{}
)

func (this *SRoomCtrl) GetRoomById(roomId int) *RoomData {
	return GRoom.get_room(roomId)
}

//获取全部房间
//func (this* SRoomCtrl) Room_all(user string) []*RoomData {
func (this *SRoomCtrl) Room_all(message data.RoomData) []*RoomData {
	user := message.User
	rooms := GRoom.room_all()
	var resData []*RoomData
	for key, room := range rooms {
		userLen := len(room.List)
		pwd := "false"
		if room.pwd != "" {
			pwd = "true"
		}
		resData = append(resData, &RoomData{
			room_id:    int64(key),
			room_name:  room.room_name,
			user_count: int32(userLen),
			battle_id:  room.battle_id,
			status:     room.status,
			pwd:        pwd,
			max:        int32(datafnc.Camp_Player_Amount),
		})
	}
	//broadcastToSelf(USER_EVENT.ROOM.ROOM_ALL, user, { room: resData })
	var roomdatas []*cmessage.RoomData = make([]*cmessage.RoomData, 0)
	for _, data := range resData {
		var list = make([]*cmessage.RoleData, 0)
		for _, dd := range data.List {
			var equips = make([]int32, 0)
			for _, id := range dd.Equips {
				equips = append(equips, int32(id))
			}

			eee := &cmessage.RoleData{
				User:   dd.User,
				Name:   dd.Name,
				Role:   dd.Role,
				HeroId: dd.Hero_id,
				Equip:  equips,
			}
			list = append(list, eee)
		}
		d := &cmessage.RoomData{
			RoomId:    data.room_id,
			RoomName:  data.room_name,
			UserCount: data.user_count,
			BattleId:  data.battle_id,
			Status:    data.status,
			Pwd:       data.pwd,
			Max:       data.max,
			List:      list,
		}
		roomdatas = append(roomdatas, d)
	}

	BroadcastToSelf("ROOM_ALL", user,
		&cmessage.GetRoomAllDataResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_GetRoomAllDataResp), 0),
			Room:       roomdatas,
		},
	)
	return resData
}

//创建房间
//func (this *SRoomCtrl) room_create(user string, battle_id int64, name string, pwd string, hero_id string, role int32) *RoomData {
func (this *SRoomCtrl) room_create(message data.RoomData) *RoomData {
	user := message.User
	battle_id := message.Battle_id
	name := message.Name
	pwd := message.Pwd
	hero_id := message.Hero_id
	role := message.Role

	room := GRoom.Room_create(user, battle_id, name, pwd, hero_id, role)
	fmt.Sprintf("error room_create")
	BroadcastToSelf("ERROR", user,
		&cmessage.ErrorResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
			Code:       room.Code,
		},
	)

	roomData := this.forMartRoomData(room)

	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}
	BroadcastAll("ROOM_CTEATE",
		&cmessage.CreateRoomResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_CreateRoomResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)
	return room
}

//加入房间
//func (this* SRoomCtrl)room_join(room_id int, user string, hero_id string, pwd string, role int32) *ReturnData {
func (this *SRoomCtrl) room_join(message data.RoomData) *ReturnData {
	room_id := message.RoomId
	user := message.User
	hero_id := message.Hero_id
	pwd := message.Pwd
	role := message.Role

	room := GRoom.room_join(room_id, user, hero_id, pwd, role)
	if room.code > 0 {
		BroadcastToSelf("ERROR", user,
			&cmessage.ErrorResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
				Code:       int32(room.code),
			},
		)
		return nil
	}

	//gvgRoomBroadcastAll(USER_EVENT.ROOM.ROOM_JOIN, room.room_id, { room: this.forMartRoomData(room.data) })
	roomData := this.forMartRoomData(room.data)

	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}
	GvgRoomBroadcastAll("ROOM_JOIN", int(room.data.room_id),
		&cmessage.JoinRoomResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_JoinRoomResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)

	return room
}

//func (this* SRoomCtrl)room_join_quick(user string, hero_id string, pwd string, role int32){
func (this *SRoomCtrl) room_join_quick(message data.RoomData) {
	user := message.User
	hero_id := message.Hero_id
	role := message.Role

	room := GRoom.room_join_quick(user, hero_id, role)
	if room.code > 0 {
		//console.log("error room_join_quick");
		BroadcastToSelf("ERROR", user,
			&cmessage.ErrorResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
				Code:       int32(room.code),
			},
		)
		return
	}
	roomData := this.forMartRoomData(room.data)

	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}
	GvgRoomBroadcastAll("ROOM_JOIN_QUICK", int(room.data.room_id),
		&cmessage.JoinRoomQuickResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_JoinRoomQuickResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)
}

//离开房间
//func (this* SRoomCtrl)room_leave(room_id int, user string) {
func (this *SRoomCtrl) room_leave(message data.RoomData) {
	room_id := message.RoomId
	user := message.User

	room := GRoom.room_leave(room_id, user)
	if room.code > 0 {
		//console.log("error room_leave");
		BroadcastToSelf("ERROR", user,
			&cmessage.ErrorResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
				Code:       int32(room.code),
			},
		)
		return
	}
	//gvgRoomBroadcastAll(USER_EVENT.ROOM.ROOM_LEAVE, room.room_id, { room: this.forMartRoomData(room.data) })
	roomData := this.forMartRoomData(room.data)

	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}
	GvgRoomBroadcastAll("ROOM_LEAVE", int(room.data.room_id),
		&cmessage.LeaveRoomResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_LeaveRoomResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)
}

//查找房间
func (this *SRoomCtrl) searchRoom() int {
	rooms := GRoom.room_all()
	roomId := 0
	for index, room := range rooms {
		//人数小于阵营最大值
		if len(room.List) < datafnc.Camp_Player_Amount {
			roomId = index
		}
	}

	return roomId
}

//房间匹配，
//func (this* SRoomCtrl)room_match(room_id int, user string) {
func (this *SRoomCtrl) room_match(message data.RoomData) {
	room_id := message.RoomId
	user := message.User

	room := GRoom.room_match(room_id, user)
	if room.code > 0 {
		//console.log("error room_match");
		BroadcastToSelf("ERROR", user,
			&cmessage.ErrorResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
				Code:       int32(room.code),
			},
		)
		return
	}

	//gvgRoomBroadcastAll(USER_EVENT.ROOM.ROOM_MATCH, body.room_id, { room: this.forMartRoomData(room) })

	roomData := this.forMartRoomData(room.data)

	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}
	GvgRoomBroadcastAll("ROOM_MATCH", int(room.data.room_id),
		&cmessage.MatchRoomResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_MatchRoomResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)
}

//房间重命名
//func (this * SRoomCtrl)room_rename(room_id int, user string, name string) {
func (this *SRoomCtrl) room_rename(message data.RoomData) {
	room_id := message.RoomId
	user := message.User
	name := message.Name

	room := GRoom.room_rename(room_id, user, name)
	if room.code > 0 {
		//console.log("error room_rename");
		BroadcastToSelf("ERROR", user,
			&cmessage.ErrorResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
				Code:       int32(room.code),
			},
		)
		return
	}
	//gvgRoomBroadcastAll(USER_EVENT.ROOM.ROOM_RENAME, body.room_id, { room: this.forMartRoomData(room) })

	roomData := this.forMartRoomData(room.data)

	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}
	GvgRoomBroadcastAll("ROOM_RENAME", int(room.data.room_id),
		&cmessage.RenameRoomResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_RenameRoomResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)

}

//房间修改密码
//func (this * SRoomCtrl)room_change_pwd(room_id int, user string, pwd string) {
func (this *SRoomCtrl) room_change_pwd(message data.RoomData) {
	room_id := message.RoomId
	user := message.User
	pwd := message.Pwd

	room := GRoom.room_change_pwd(room_id, user, pwd)
	if room.code > 0 {
		//console.log("error room_change_pwd");
		BroadcastToSelf("ERROR", user,
			&cmessage.ErrorResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
				Code:       int32(room.code),
			},
		)
		return
	}
	//gvgRoomBroadcastToSelf(USER_EVENT.ROOM.ROOM_CHANGE_PWD, body.room_id, body.user, { room: this.forMartRoomData(room) })
	roomData := this.forMartRoomData(room.data)

	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}
	GvgRoomBroadcastAll("ROOM_CHANGE_PWD", int(room.data.room_id),
		&cmessage.ChangepwdRoomResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ChangepwdRoomResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)
}

//踢人
//func (this* SRoomCtrl)room_kick_off(room_id int, user string, kick_user string) {
func (this *SRoomCtrl) room_kick_off(message data.RoomData) {
	room_id := message.RoomId
	user := message.User
	kick_user := message.KickUser

	room := GRoom.room_kick_off(room_id, user, kick_user)
	if room.code > 0 {
		//console.log("error room_kick_off");
		BroadcastToSelf("ERROR", user,
			&cmessage.ErrorResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
				Code:       int32(room.code),
			},
		)
		return
	}
	roomData := this.forMartRoomData(room.data)
	//gvgRoomBroadcastAll(USER_EVENT.ROOM.ROOM_KICK_OFF, room_id, { room: this.forMartRoomData(room) })
	roledatas := make([]*cmessage.RoleData, 0)
	for _, data := range roomData.List {
		list := make([]int32, 0)
		for _, id := range data.Equips {
			list = append(list, int32(id))
		}

		roledata := &cmessage.RoleData{
			User:   data.User,
			Name:   data.Name,
			Role:   data.Role,
			HeroId: data.Hero_id,
			Equip:  list,
		}

		roledatas = append(roledatas, roledata)
	}

	GvgRoomBroadcastAll("ROOM_KICK_OFF", room_id,
		&cmessage.RoomKickOffResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ErrorResp), 0),
			Room: &cmessage.RoomData{
				RoomId:    roomData.room_id,
				RoomName:  roomData.room_name,
				UserCount: roomData.user_count,
				BattleId:  roomData.battle_id,
				Status:    roomData.status,
				Pwd:       roomData.pwd,
				Max:       roomData.max,
				List:      roledatas,
			},
		},
	)
}

//func (this* SRoomCtrl)room_quick_Join(room_id int, user string, role int32, hero_id string) {
func (this *SRoomCtrl) room_quick_Join(message data.RoomData) {
	room_id := message.RoomId
	user := message.User
	role := message.Role
	hero_id := message.Hero_id

	roomId := this.searchRoom()
	var room *ReturnData
	if roomId != 0 {
		room = this.room_join(
			data.RoomData{
				RoomId:  room_id,
				User:    user,
				Hero_id: "16391290684511",
				Pwd:     "1",
				Role:    role,
			})
		// gvgRoomBroadcastAll(USER_EVENT.ROOM.ROOM_JOIN, body.room_id, { room })
	} else {
		room.data = RoomCtrl.room_create(data.RoomData{
			User:      user,
			Battle_id: 1,
			Name:      "0000",
			Pwd:       "1",
			Hero_id:   hero_id,
			Role:      role,
		})
		// gvgRoomBroadcastAll(USER_EVENT.ROOM.ROOM_CTEATE, body.room_id, { room })
	}
}

func (this *SRoomCtrl) forMartRoomData(room *RoomData) *RoomData {
	if room.pwd != "" {
		room.pwd = "true"
	} else {
		room.pwd = "false"
	}
	room.max = int32(datafnc.Camp_Player_Amount)
	return room
}
