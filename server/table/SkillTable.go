package table

type SkillCfg struct {
	ID                int `json:"ID"`
	Name_String       int `json:"Name_String"`
	Desc_String       int
	Type              int
	Rate              int
	Range             float32
	Center            int
	Shape             int
	Radius            []float32
	Buff              []int
	Buff_Target       int
	Target_Num        int
	Last_Time         int
	CD                float32
	Need_Target       int
	Bullet_speed      int
	CastingEffect     string
	CastingEffectSize int
	FlightEffect      string
	FlightEffectSize  int
	ShowEffec         string
	ShowEffectSize    int
}
type SkillTableData map[string]SkillCfg
