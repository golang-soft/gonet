package gamedata

import (
	"container/list"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/world/datafnc"
	"time"
)

//GlobalData room 房间数据逻辑
type ServerMatch struct {
	Match interface{}
	//队伍匹配
	TeamMatch map[int]*GMatchData
	//快速匹配
	QuickMatch map[int]interface{}
}

var GMatch ServerMatch = ServerMatch{}

type GMatchData struct {
	roomId int
	list   list.List
	start  int64
}

func (this *ServerMatch) constructor() {
	this.Match = new(interface{})
	this.TeamMatch = make(map[int]*GMatchData)
	this.QuickMatch = make(map[int]interface{})
}

//全部匹配数据
func (this *ServerMatch) getAllRoom() interface{} {
	if this.Match != nil {
		return this.Match
	}
	return nil
}

//验证匹配房间
func (this *ServerMatch) CheckMatch() {
	this.teamMatch()
	this.quickMatch()
}

//加入匹配队列
func (this *ServerMatch) add2TeamMatchQueue(roomId int, userList list.List) {
	if this.TeamMatch[roomId] != nil {
		return
	}

	start := time.Now().Unix()
	logData := data.LogData{
		Mode: common.LogMode.Match,
		//Userid: userList[0].user,
		Start: start,
	}
	AddLogTask(logData)
	this.TeamMatch[roomId] = &GMatchData{roomId, userList, start}
}

//退出p匹配
func (this *ServerMatch) quit_match(room_id int) {
	delete(this.TeamMatch, room_id)
}
func (this *ServerMatch) quit_match_user(user string) {

}

func (this *ServerMatch) teamMatch() {
	//队伍匹配
	for key, _ := range this.TeamMatch {
		if this.TeamMatch[key] != nil && this.TeamMatch[key+1] != nil {
			red := this.TeamMatch[key].list
			blue := this.TeamMatch[key+1].list
			//房间匹配
			GameCtrl.matchGameStart(red, blue)
			//删除房间数据
			delete(this.TeamMatch, this.TeamMatch[key].roomId)
			delete(this.TeamMatch, this.TeamMatch[key+1].roomId)
		} else {
			break
		}
	}
}

//退出匹配队列
func (this *ServerMatch) quickMatchQueue(roomId string, user interface{}) {

}

func (this *ServerMatch) getTeam() {

}

func (this *ServerMatch) quickMatch() {
	//队伍匹配
	allMatch := make([]string, 0)
	for index := 0; index < len(this.QuickMatch); index++ {
		if len(allMatch) == datafnc.Room_Player_Max_Size {
			this.battleStart()
		}
	}
}

func (this *ServerMatch) battleStart() {

}

//房间加入匹配
func (this *ServerMatch) room2match(roomId string) {

}

//退出匹配
func (this *ServerMatch) quitMatch() {

}

func (this *ServerMatch) matchRoom() {

}
