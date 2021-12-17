package table

type StrengthenTableCfg struct {
	ID        int   `json:"ID"`
	Class     int   `json:"Class"`
	Part      int   `json:"Part"`
	Level     int   `json:"Level"`
	Attribute []int `json:"Attribute"`
}

type StrengthenTableData map[string]StrengthenTableCfg
