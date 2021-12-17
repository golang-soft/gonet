package table

type RandNameTableCfg struct {
	ID       int    `json:"ID"`
	UserName string `json:"UserName"`
}

type RandNameTableData map[string]RandNameTableCfg
