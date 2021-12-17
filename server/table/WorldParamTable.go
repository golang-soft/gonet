package table

type WorldParamCfg struct {
	ID    int    `json:"ID"`
	Key   string `json:"Name_String"`
	Value string `json:"Value"`
	Desc  string `json:"Desc"`
}

type WorldParamCfgTableData map[string]WorldParamCfg
