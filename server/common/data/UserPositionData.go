package data

type UserPositionData struct {
	X             float64
	Y             float64
	Speed         float64
	ReduceSpeedTs int64
	Dizzy         int32
	DizzyTs       int64
	Barrier       int32
	Direction     float64
	PosUpdateTs   int64
	StopMove      int32
	StopMoveTs    int64
	SkillId       int32
}

type Pos struct {
	X float64
	Y float64
}
