package data

type UserGameAttrData struct {
	User          string
	UpdateTs      int64
	DefPercent    int     //护甲增加或者减少的buff
	Hp            float64 // 血条
	Part          int32   // 阵营
	Itype         int     // 职业类型
	DieTs         int64
	Shield        float64
	ShieldTs      int64
	Immune        float64
	ImmuneTs      int64
	Thorns        float64
	ThornsTs      int64
	Stopmove      int32
	StopmoveTs    int64
	AddDef        int
	AddDefTs      int64
	AddAtk        float64
	AllAttr       int
	Dvt           int32
	GetDvt        int
	desDvt        int
	Kill          int
	Die           int
	Dps           float64
	skill_1101_cd int
	skill_1102_cd int
	skill_1103_cd int
	skill_1104_cd int
	skill_1201_cd int
	skill_1202_cd int
	skill_1203_cd int
	skill_1204_cd int
	skill_1301_cd int
	skill_1302_cd int
	skill_1303_cd int
	skill_1304_cd int
	skill_1401_cd int
	skill_1402_cd int
	skill_1403_cd int
	skill_1404_cd int
	skill_1501_cd int
	skill_1502_cd int
	skill_1503_cd int
	skill_1504_cd int
	X             float64
	Y             float64
	PosUpdateTs   int64
	ReduceSpeedTs int64
	Dizzy         int32
	DizzyTs       int64
	Speed         float32
	Item          map[int]ItemData
	Equip         map[int]int
	Direction     float64
	Barrier       int32
	Equip_3       int
	Equip_9       float64
	Equip_1       int
	Equip_2       float64
	Equip_4       float64
	Equip_5       float64
	Equip_6       float64
	Round         int
	Hid           string
	DesDvt        int
	Skill_1101Cd  int32
	Skill_1102Cd  int32
	Skill_1103Cd  int32
	Skill_1104Cd  int32
}
