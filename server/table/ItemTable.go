package table

type ItemTableCfg struct {
	ID               int   `json:"ID"`
	Name_String      int   `json:"Name_String"`
	Type             int   `json:"Type"`
	Effect_String    int   `json:"Effect_String"`
	Desc_String      int   `json:"Desc_String"`
	Class            int   `json:"Class"`
	Part             int   `json:"Part"`
	Part_Name_String int   `json:"Part_Name_String"`
	Quality          int   `json:"Quality"`
	Attribute        []int `json:"Attribute"`
	Last_time        int   `json:"Last_time"`
	CD               int   `json:"CD"`
	Flag             int   `json:"Flag"`
	Battle_id        int   `json:"Battle_id"`
	Model            int   `json:"Model"`
	Icon             int   `json:"Icon"`
}

type ItemTableData map[string]ItemTableCfg
