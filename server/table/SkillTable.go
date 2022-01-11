package table

type SkillCfg struct {
	ID                int `json:"ID"`
	Name_String       int `json:"Name_String"`
	Desc_String       int
	Type              int
	Rate              float64
	Range             float64
	Center            int
	Shape             int
	Radius            []float64
	Buff              []float64
	Buff_Target       int
	Target_Num        int
	Last_Time         int
	CD                float32
	Need_Target       int
	Bullet_speed      float64
	CastingEffect     string
	CastingEffectSize int
	FlightEffect      string
	FlightEffectSize  int
	ShowEffec         string
	ShowEffectSize    int
}
type SkillTableData map[int32]*SkillCfg
