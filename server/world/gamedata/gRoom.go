package gamedata

import (
	"container/list"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/world/cache"
	"gonet/server/world/data"
	"gonet/server/world/datafnc"
	"math/rand"
	"time"
)

const Room_Max = 1000

//GlobalData room 房间数据逻辑

type RoomData struct {
	room_name  string
	pwd        string
	status     int32
	battle_id  int64
	room_id    int64
	owner_id   string
	list       []*data.TmpRoomPlayerData
	Code       int32
	user_count int32
	max        int32
}
type (
	ServerRoom struct {
		User         map[string]interface{}
		Room         map[int]*RoomData
		RoomIdx      interface{}
		RoomCount    int
		next_room_id int64
		RoomOwner    []string
	}

	IServerRoom interface {
	}
)

var GRoom *ServerRoom = &ServerRoom{}

func (this *ServerRoom) constructor() {
	this.next_room_id = 0
}

func (this *ServerRoom) getNextRoomId() int64 {
	this.next_room_id += 1
	return this.next_room_id
}

func (this *ServerRoom) initPlayer() *cache.PlayerInfo {
	return &cache.PlayerInfo{
		Id:       "address",
		NickName: "noName",
		Email:    "email",
		Dvt:      0,
		CreateTs: 0,
		UpdateTs: 0,
	}
}

//全部房间
func (this *ServerRoom) room_all() map[int]*RoomData {
	if this.Room != nil {
		return this.Room
	}
	return nil
}

func (this *ServerRoom) get_room(roomId int) interface{} {
	if this.Room != nil && this.Room[roomId] != nil {
		return this.Room[roomId]
	}
	return nil
}

func (this *ServerRoom) rename(roomId int, name string) {
	this.Room[roomId].room_name = name
}

func (this *ServerRoom) changePwd(roomId int, pwd string) {
	this.Room[roomId].pwd = pwd
}
func (this *ServerRoom) startGame(roomId int) {
	this.Room[roomId].status = common.RoomSatus.Game_start
}

//删除房间
func (this *ServerRoom) delRoom(roomId int) {
	delete(this.Room, roomId)
	this.RoomCount -= 1
}

func (this *ServerRoom) creatRoom() *RoomData {
	return &RoomData{
		battle_id: 0,
		room_id:   0,
		owner_id:  "",
		room_name: "",
		pwd:       "",
		status:    0,
		list:      make([]*data.TmpRoomPlayerData, 0),
	}
}

func (this *ServerRoom) CheckRoom() {
	var delList = make([]int, 0)
	for key, data := range this.Room {
		if data != nil {
			userList := data.list
			//判断有没有用户在线
			online := CheckRoomOnline(data.room_id, userList)
			if !online {
				delList = append(delList, key)
			}
		}
	}

	for _, roomId := range delList {
		//产出没用户的房间
		delete(this.Room, roomId)
	}
}

func (this *ServerRoom) addNewRoom(room *RoomData) {
	if this.Room[int(room.room_id)] != nil {
		return
	}
	this.Room[int(room.room_id)] = room
	//房间
	this.RoomCount += 1
	// this.next_room_id += 1
}

func (this *ServerRoom) newRoomPlayer(fromUser *cache.PlayerData, hero_id string, role int32) *data.TmpRoomPlayerData {
	name := fromUser.Player.NickName
	user := fromUser.Player.Id
	//更具heroid 获取英雄属性和装备
	// let role: number = fromUser.hero[hero_id].heroType
	var equipList map[int]int
	heroRole := role
	if fromUser.Hero[hero_id] != nil {
		for _, equip := range fromUser.Hero[hero_id].Equips {
			id := equip.ItemId
			equipList[id] = id
		}

		// let list = fromUser.hero[hero_id].equips.map(equip => {
		//     equip.itemId
		// })

		heroRole = fromUser.Hero[hero_id].HeroType
	}
	//equipList.PushBack(2001)
	//equipList.PushBack(2002)
	//equipList.PushBack(2003)
	//equipList.PushBack(2004)
	//equipList.PushBack(2021)

	return &data.TmpRoomPlayerData{
		User:    user,
		Name:    name,
		Role:    heroRole,
		Hero_id: hero_id,
		Equips:  equipList,
	}
}

func (this *ServerRoom) createNewPlayer(user string) *cache.PlayerInfo {
	newPlayer := &cache.PlayerInfo{}
	nowTs := time.Now().Unix()

	newPlayer.Id = user
	newPlayer.NickName = user
	newPlayer.Email = "111@gmail.com"
	newPlayer.CreateTs = int(nowTs)
	newPlayer.UpdateTs = int(nowTs)
	newPlayer.Dvt = 99999

	return newPlayer
}

func (this *ServerRoom) Room_create(user string, battle_id int64, name string, pwd string, hero_id string, role int32) *RoomData {
	room_id := this.getNextRoomId()
	//验证用户合法化，拉取最新数据
	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		fromUser.Player = this.createNewPlayer(user)
	}

	if fromUser.Hero[hero_id] != nil {

	}

	if fromUser.Player.Dvt < datafnc.Entry_Fee {
		//console.log("dvt_not_enough");
	}

	//判断dvt
	newRoom := this.creatRoom()
	newRoom.room_id = room_id
	newRoom.room_name = name
	newRoom.pwd = pwd
	newRoom.battle_id = battle_id
	newRoom.owner_id = user
	newRoom.list = append(newRoom.list, this.newRoomPlayer(fromUser, hero_id, role))

	this.addNewRoom(newRoom)
	//TODO
	//await socketJoin2GvgRoom([user], room_id)
	return newRoom
}

func (this *ServerRoom) isPlayerInRoom(list []*data.TmpRoomPlayerData, user string) bool {
	for _, player := range list {
		if player.User == user {
			return true
		}
	}
	return false
}

//加入房间
func (this *ServerRoom) room_join(room_id int, user string, hero_id string, pwd string, role int32) *ReturnData {
	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		//用户不合法
		//return &ReturnData{code: common.ErrorCode.No_player,}
		fromUser.Player = this.createNewPlayer(user)
	}
	if fromUser.Hero[hero_id] == nil {
		//没有英雄
	}
	if fromUser.Player.Dvt < datafnc.Entry_Fee {
		//dvt 不足
	}
	if this.Room[room_id] == nil {
		//房间不存在
		return &ReturnData{code: common.ErrorCode.No_room}
	}
	if this.Room[room_id].status == common.RoomSatus.Game_start {
		//游戏开始
		return &ReturnData{code: common.ErrorCode.Already_start}
	}
	if this.Room[room_id].pwd != pwd {
		//密码不同
		return &ReturnData{code: common.ErrorCode.Pwd_error}
	}

	if this.isPlayerInRoom(this.Room[room_id].list, user) {
		//人已在房间
		return &ReturnData{
			code: common.ErrorCode.Already_in,
		}
	}

	// let userList = Object.keys(this.Room[room_id].user)

	if len(this.Room[room_id].list) >= datafnc.Camp_Player_Amount {
		//人数上限
		return &ReturnData{
			code: common.ErrorCode.Play_max_limit,
		}
	}
	//添加用户数据到房间
	this.userJoinRoom(room_id, fromUser, hero_id, role)
	SocketJoin2GvgRoom([]string{user}, room_id)
	return &ReturnData{data: this.Room[room_id]}
}

//加入房间
func (this *ServerRoom) room_join_quick(user string, hero_id string, role int32) *ReturnData {
	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		//用户不合法
		//return &ReturnData{code: common.ErrorCode.No_player,}
		fromUser.Player = this.createNewPlayer(user)
	}
	if fromUser.Hero[hero_id] == nil {
		//没有英雄
	}
	if fromUser.Player.Dvt < datafnc.Entry_Fee {
		//dvt 不足
	}

	var room_id int = -1
	for key, roomData := range this.Room {
		if roomData.pwd != "" {
			continue
		}
		if len(roomData.list) >= datafnc.Camp_Player_Amount {
			continue
		}
		room_id = key
		break
	}

	if room_id == -1 {
		//没有空闲房间
		return &ReturnData{
			code: common.ErrorCode.No_empty_room,
		}
	}

	this.userJoinRoom(room_id, fromUser, hero_id, role)
	//添加用户数据到房间
	SocketJoin2GvgRoom([]string{user}, room_id)
	//检查房间人数
	return &ReturnData{data: this.Room[room_id]}
}

func (this *ServerRoom) userJoinRoom(roomId int, play *cache.PlayerData, hero_id string, role int32) *data.TmpRoomPlayerData {
	newUser := this.newRoomPlayer(play, hero_id, role)
	this.Room[roomId].list = append(this.Room[roomId].list, newUser)
	return newUser
}

//移除房间用户
func (this *ServerRoom) removeUser(roomId int, user string) *RoomData {
	var index int = 0
	for _, data := range this.Room[roomId].list {
		if data.User == user {
			break
		}
		index++
	}

	if index != -1 {
		this.Room[roomId].list = append(this.Room[roomId].list[:index], this.Room[roomId].list[index+1:]...)
	}

	return this.Room[roomId]
}

//离开房间
func (this *ServerRoom) room_leave(room_id int, user string) *ReturnData {
	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		//用户不合法
		return &ReturnData{code: common.ErrorCode.No_player}
	}
	if this.Room[room_id] == nil {
		//房间不存在
		return &ReturnData{code: common.ErrorCode.No_room}
	}
	if len(this.Room[room_id].list) == 0 {
		//人不在房间
		return &ReturnData{code: common.ErrorCode.No_player}
	}

	//删除用户
	this.removeUser(room_id, user)
	if len(this.Room[room_id].list) > 0 {
		if this.Room[room_id].owner_id == user {
			next_id := this.Room[room_id].list[0].User
			this.Room[room_id].owner_id = next_id
		}
		SocketLeaveGvgRoom([]string{user}, room_id)
		return &ReturnData{data: this.Room[room_id]}
	} else {
		//房间销毁 房间数减一
		SocketLeaveGvgRoom([]string{user}, room_id)
		this.delRoom(room_id)
		//TODO:退出房间匹配
		// MatchCtrl.quit_match(room_id)
		//broadcastAll(USER_EVENT.ROOM.ROOM_DEL, { room_id: room_id })

		BroadcastAll("ROOM_DEL",
			&cmessage.DelRoomResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_DelRoomResp), 0),
				RoomId:     int64(room_id),
			},
		)

		return nil
	}
	return nil
}

//踢人
func (this *ServerRoom) room_kick_off(room_id int, user string, kick_user string) *ReturnData {

	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		return &ReturnData{code: common.ErrorCode.No_player}
	}
	if this.Room[room_id] == nil {
		//房间不存在
		return &ReturnData{code: common.ErrorCode.No_room}
	}
	if this.Room[room_id].owner_id != user {
		//判断房主
		return &ReturnData{code: common.ErrorCode.Not_owner}
	}

	//删除用户
	this.removeUser(room_id, kick_user)
	userList := this.Room[room_id].list
	if len(userList) > 0 {
		if this.Room[room_id].owner_id == user {
			//房主转移
			Shuffle(userList)
			this.Room[room_id].owner_id = userList[0].User
		}
		//TODO
		//socketLeaveGvgRoom([kick_user], room_id)
		return &ReturnData{data: this.Room[room_id]}
	} else {
		//房间销毁 房间数减一
		//TODO
		//socketLeaveGvgRoom([kick_user], room_id)
		this.delRoom(room_id)
		return nil
	}
	//广播房间用户
	return nil
}

func Shuffle(slice []*data.TmpRoomPlayerData) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

//房间匹配，
func (this *ServerRoom) room_match(room_id int, user string) *ReturnData {
	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		return &ReturnData{code: common.ErrorCode.No_player}
	}
	if this.Room[room_id] == nil {
		return &ReturnData{code: common.ErrorCode.No_room}
	}
	if this.Room[room_id].owner_id != user {
		return &ReturnData{code: common.ErrorCode.Not_owner}
	}
	if len(this.Room[room_id].list) != datafnc.Camp_Player_Amount {
		return &ReturnData{code: common.ErrorCode.Not_owner}
	}
	if this.Room[room_id].status == common.RoomSatus.Room_wait {
		return &ReturnData{code: common.ErrorCode.Already_start}
	}
	//房间状态游戏开始
	this.startGame(room_id)
	var userList list.List

	for _, data := range this.Room[room_id].list {
		userList.PushBack(data)
	}

	matchReq := &MatchData{
		roomId: room_id,
		list:   userList,
	}

	MatchCtrl.room2Match(matchReq)
	return &ReturnData{data: this.Room[room_id]}
}

//房间重命名
func (this *ServerRoom) room_rename(room_id int, user string, name string) *ReturnData {
	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		return &ReturnData{code: common.ErrorCode.No_player}
	}
	if this.Room[room_id] == nil {
		return &ReturnData{code: common.ErrorCode.No_room}
	}
	if this.Room[room_id].owner_id != user {
		return &ReturnData{code: common.ErrorCode.Not_owner}
	}
	if this.Room[room_id].status == common.RoomSatus.Game_start {
		return &ReturnData{code: common.ErrorCode.Already_start}
	}

	this.rename(room_id, name)
	return &ReturnData{data: this.Room[room_id]}
}

type ReturnData struct {
	code int
	data *RoomData
}

//房间修改密码
func (this *ServerRoom) room_change_pwd(room_id int, user string, pwd string) *ReturnData {
	fromUser := cache.GameCache.GetPlayer(user)
	if fromUser.Player == nil {
		return &ReturnData{
			code: common.ErrorCode.No_player,
		}
	}
	if this.Room[room_id] == nil {
		return &ReturnData{
			code: common.ErrorCode.No_room,
		}
	}
	if this.Room[room_id].owner_id != user {
		return &ReturnData{
			code: common.ErrorCode.Not_owner,
		}
	}
	if this.Room[room_id].status == common.RoomSatus.Game_start {
		return &ReturnData{code: common.ErrorCode.Already_start}
	}

	this.changePwd(room_id, pwd)
	return &ReturnData{data: this.Room[room_id]}
}
