package gamedata

import (
	"errors"
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/golang/protobuf/proto"
	"gonet/server/common/data"
	"gonet/server/world/datafnc"
	"gonet/server/world/helper"
	"gonet/server/world/param"
	"gonet/server/world/socket"
	"math"
	"reflect"
	"strings"
	"time"
)

var (
	unknonw_err = errors.New("unknown error")
)

type Position struct {
	Min float32
	Max float32
}
type PositionConfig struct {
	SPEED Position
}

var POSITION_CONFIG = &PositionConfig{
	SPEED: Position{
		Min: 0,
		Max: 8,
	},
}

func handleUnknownError(err error) {
	fmt.Sprintf("%v", unknonw_err)
}

//获取房间用户
func getRoomUsers(round int) []string {
	var userList []string = make([]string, 0)
	sockets := socket.FetchSockets()
	for _, socket := range sockets {
		userList = append(userList, socket.Data.User)
	}
	return userList
}

func Includes(user []*param.TmpRoomPlayerData, u string) bool {
	for _, data := range user {
		if u == data.User {
			return true
		}
	}
	return false
}

func IncludeUsers(user []string, u string) bool {
	for _, data := range user {
		if u == data {
			return true
		}
	}
	return false
}

//获取同阵营在线用户
func CheckRoomOnline(roomId int64, users []*param.TmpRoomPlayerData) bool {
	var online = false
	sockets := socket.FetchSockets()
	for _, socket := range sockets {
		if Includes(users, socket.Data.User) && socket.Data.RoomId == roomId {
			online = true
			break
		}
	}
	return online
}

func GetRoomUsersByPart(round int, part int32) []string {
	userList := make([]string, 0)
	sockets := socket.FetchSockets()
	for _, socket := range sockets {
		if socket.Data.Round == round && socket.Data.Part == part {
			userList = append(userList, socket.Data.User)
		}
	}

	return userList
}

func getConnSockets() map[int]*socket.Socket {
	return socket.FetchSockets()
}

func BroadcastAll(env string, data proto.Message) {
	socket.Emit(env, data)
}

func BroadcastToSelf(env string, user string, message proto.Message) {

	sockets := socket.FetchSockets()
	for _, socket := range sockets {
		if socket.Data.User == user {
			socket.Emit(env, message)
		}
	}
}

//加入战场
func SocketJoin2BattleRoom(users []string, round int64, userPart map[string]int32) {
	sockets := socket.FetchSockets()
	for _, socket := range sockets {
		if IncludeUsers(users, socket.Data.User) {
			//添加战场round
			socket.Data.Round = int(round)
			//添加阵容id
			if userPart[socket.Data.User] != 0 {
				socket.Data.Part = userPart[socket.Data.User]
			}
			socket.Join(fmt.Sprintf("Battle%d", round))
		}
	}
}

func SocketLeaveBattleRoom(users []string, round int) {
	sockets := socket.FetchSockets()
	var delId int
	for _, socket := range sockets {
		if socket.Data.Round == round {
			socket.Leave(fmt.Sprintf("Battle%d", round))
			delId = socket.Data.Round
			break
		}
	}

	delete(sockets, delId)
}

func In(sockets map[int]*socket.Socket, battleName string) *socket.Socket {
	if sockets != nil {
		for _, socket := range sockets {
			if socket.Room.Name == battleName {
				return socket
			}
		}
	}

	return nil
}

//战场房间广播，所有房间内
func GvgBattleBroadcastAll(env string, round int, data proto.Message) {
	sockets := socket.FetchSockets()
	if sockets != nil {
		battleName := fmt.Sprintf("Battle%d", round)
		socket := In(sockets, battleName)
		if socket != nil {
			socket.Emit(env, data)
		}
	}
}

func SocketJoin2GvgRoom(users []string, roomId int) {
	sockets := socket.FetchSockets()
	if sockets != nil {
		for _, socket := range sockets {
			if IncludeUsers(users, socket.Data.User) {
				socket.Data.RoomId = int64(roomId)
				battleName := fmt.Sprintf("Battle%d", roomId)
				socket.Join(battleName)
			}
		}
	}
}

func GvgRoomBroadcastAll(env string, roomId int, data proto.Message) {
	sockets := socket.FetchSockets()
	if sockets != nil {
		battleName := fmt.Sprintf("room%d", roomId)
		socket := In(sockets, battleName)
		if socket != nil {
			socket.Emit(env, data)
		}
	}
}

//房间广播，自己
func gvgRoomBroadcastToSelf(env string, roomId int64, user string, data proto.Message) {
	sockets := socket.FetchSockets()
	if sockets != nil {
		for _, socket := range sockets {
			if socket.Data.User == user && socket.Data.RoomId == roomId {
				socket.Emit(env, data)
			}
		}
	}
}

func SocketLeaveGvgRoom(users []string, roomId int) {
	var delId int = 0
	sockets := socket.FetchSockets()
	if sockets != nil {
		for _, socket := range sockets {
			if IncludeUsers(users, socket.Data.User) {
				battleName := fmt.Sprintf("Battle%d", roomId)
				socket.Leave(battleName)
				delId = socket.Data.Round
			}
		}
	}
	delete(sockets, delId)
}

//房间广播，自己
func roomBroadcastToSelf(env string, round int, user string, data proto.Message) {
	sockets := socket.FetchSockets()
	if sockets != nil {
		for _, socket := range sockets {
			if socket.Data.User == user && socket.Data.Round == round {
				socket.Emit(env, data)
			}
		}
	}
}

func InitUserBasic(round int, user string, hid string, role int32, part int32, equip map[int]int, attrAll int, battle_id int) *data.UserGameAttrData {
	itype := role
	var skillAttr map[string]int
	for _, skillId := range helper.USER_BASIC_INFO[int(itype)].Skill {
		key := fmt.Sprintf("Skill_%d_cd", skillId)
		skillAttr[key] = 0
	}

	itemAttr := datafnc.GetBattleItemsById(battle_id)
	var equipAttr map[string]int
	for _, eData := range equip {
		equipCfg := datafnc.GetEquipCfg(int32(eData))
		if equipCfg != nil {
			//attr 叠加
			equipKey := fmt.Sprintf("Equip_%d", equipCfg.Attribute[0])
			if equipAttr[equipKey] > 0 {
				equipAttr[equipKey] += equipCfg.Attribute[1]
			} else {
				equipAttr[equipKey] = equipCfg.Attribute[1]
			}
		}
	}

	/*
	  1-攻击
	    2-防御
	    3-血量
	    4-暴击
	    5-爆伤
	    6-闪避
	    9-全属性
	*/
	//道具全属性
	var addAttrAll float64 = 0
	if attrAll != 0 {
		addAttrAll = datafnc.All_Attr_Percent
	}

	var hp float64 = helper.USER_BASIC_INFO[int(itype)].Hp
	if equipAttr["Equip_3"] != 0 {
		hp += float64(equipAttr["Equip_3"])
	}

	if equipAttr["Equip_9"] != 0 {
		addPercent := float64(equipAttr["Equip_9"])/100 + addAttrAll
		hp += math.Floor(hp * addPercent)
	}

	bornData := datafnc.GetBornPoint(part, 1)

	ret := &data.UserGameAttrData{
		User:          user,
		Round:         round,
		Part:          part,
		Hid:           hid,
		Itype:         int(itype),
		DefPercent:    0,
		Hp:            hp,
		UpdateTs:      0,
		X:             bornData.X,
		Y:             bornData.Y,
		Speed:         0,
		ReduceSpeedTs: 0,
		Direction:     0,
		Barrier:       0,
		Dizzy:         0,
		DizzyTs:       0,
		Shield:        0,
		ShieldTs:      0,
		Immune:        0,
		ImmuneTs:      0,
		Thorns:        0,
		ThornsTs:      0,
		StopmoveTs:    0,
		Stopmove:      0,
		AddDef:        0,
		AddDefTs:      0,
		AddAtk:        0,
		PosUpdateTs:   time.Now().Unix(),
		DieTs:         0,
		AllAttr:       attrAll,
		Dvt:           0,
		GetDvt:        0,
		DesDvt:        0,
		Kill:          0,
		Die:           0,
		Dps:           0,
	}
	var data0 data.UserGameAttrData
	ConvertToUserGameAttrData(skillAttr, data0, ret, "Skill_")
	ConvertToUserGameAttrData(itemAttr, data0, ret, "item_")
	ConvertToUserGameAttrData(equipAttr, data0, ret, "Equip_")

	return ret
}

func ConvertToUserGameAttrData(dataAttr map[string]int, data0 data.UserGameAttrData, ret *data.UserGameAttrData, prefix string) {
	err := mapstructure.Decode(dataAttr, data0)
	if err != nil {
		fmt.Println(err)
	}
	vt := reflect.TypeOf(*ret)
	vv := reflect.ValueOf(*ret)

	for i := 0; i < vt.NumField(); i++ {
		fieldName := vt.Field(i).Name
		field := reflect.ValueOf(ret).Elem().FieldByName(fieldName)
		if strings.HasPrefix(fieldName, prefix) {
			if vv.FieldByName(fieldName).Int() > 0 {
				field.SetInt(vv.FieldByName(fieldName).Int())
			}
		}
	}
}
