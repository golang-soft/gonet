package table

type HeroAttrTableCfg struct {
	ID             int     `json:"ID"`
	Class          int     `json:"Class"`
	String_ID      int     `json:"String_ID"`
	Attack         float64 `json:"Attack"`
	Defend         float64 `json:"Defend"`
	HP             float64 `json:"HP"`
	Crit           float64 `json:"Crit"`
	Crit_Damage    float64 `json:"Crit_Damage"`
	Dodge          float64 `json:"Dodge"`
	Speed          float64 `json:"Speed"`
	DMG            int     `json:"DMG"`
	DEF            int     `json:"DEF"`
	SPD            int     `json:"SPD"`
	CC             int     `json:"CC"`
	SUP            int     `json:"SUP"`
	Skill          []int   `json:"Skill"`
	Hp_Height      float32 `json:"Hp_Height"`
	Face_Height    float32 `json:"Face_Height"`
	Jump_Speed     int     `json:"Jump_Speed"`
	Story_StringID int     `json:"Story_StringID"`
	Model          string  `json:"Model"`
}

type HeroAttrTableData map[string]HeroAttrTableCfg
