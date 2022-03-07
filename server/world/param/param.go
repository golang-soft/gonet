package param

import "gonet/server/common/data"

type Param struct {
}

type RoomParam struct {
	Param
	User     string
	BattleId int64
	Name     string
	Pwd      string
	HeroId   string
	Role     int32
	RoomId   int
	To       string
}

type UserParam struct {
	Param
	User      string
	BattleId  int64
	Name      string
	Pwd       string
	HeroId    string
	Role      int32
	RoomId    int
	To        string
	Msg       string
	Mode      int
	ItemId    int
	Count     int
	Round     int
	Skill     int32
	SkillId   int32
	X         float64
	Y         float64
	Z         float64
	Speed     float64
	Direction float64
	Barrier   int32
	Pos       data.Pos
}
