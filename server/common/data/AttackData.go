package data

type AttackData struct {
	Round   int
	From    string
	To      string
	Skill   int32
	SkillId int32
	X       float64
	Y       float64
	Z       float64
	Msg     string
	Mode    int
	ItemId  int
	Part    int
	Count   int
	Cd      int64
	Role    string
}
