package table

type HeroAttrTableCfg struct {
	ID             int     `json:"ID"`
	Class          int     `json:"Class"`
	String_ID      int     `json:"String_ID"`
	Attack         float32 `json:"Attack"`
	Defend         float32 `json:"Defend"`
	HP             int     `json:"HP"`
	Crit           int     `json:"Crit"`
	Crit_Damage    int     `json:"Crit_Damage"`
	Dodge          int     `json:"Dodge"`
	Speed          float32 `json:"Speed"`
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
