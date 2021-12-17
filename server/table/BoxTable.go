package table

type ItemCfg []int

type BoxTableCfg struct {
	BoxID int       `json:"BoxID"`
	Type  int       `json:"Type"`
	Item  []ItemCfg `json:"Item"`
	Icon  int       `json:"Icon"`
}

type BoxTableData map[string]BoxTableCfg
