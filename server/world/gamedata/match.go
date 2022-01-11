package gamedata

import "container/list"

type SMatchCtrl struct {
}

var MatchCtrl SMatchCtrl = SMatchCtrl{}

//获取全部房间
func (this *SMatchCtrl) room_all(body interface{}) {

}

type MatchData struct {
	roomId int
	list   list.List
	number int
}

//房间加入匹配队列
func (this *SMatchCtrl) room2Match(body *MatchData) {
	var room_id = body.roomId
	var list = body.list
	GMatch.add2TeamMatchQueue(room_id, list)
}

//房间匹配
func (this *SMatchCtrl) room_match(body interface{}) {
}

//快速匹配
func (this *SMatchCtrl) quickMatch() {
}

//退出匹配
func (this *SMatchCtrl) quit_match(room_id int) {
	GMatch.quit_match(room_id)
}

func (this *SMatchCtrl) quit_match_user(user string) {
	GMatch.quit_match_user(user)
}

//离开房间
func (this *SMatchCtrl) room_leave(body interface{}) {
}

//房间重命名
func (this *SMatchCtrl) room_rename(body interface{}) {
}

//房间修改密码
func (this *SMatchCtrl) room_change_pwd(body interface{}) {
}

//踢人
func (this *SMatchCtrl) room_kick_off(body interface{}) {
}
