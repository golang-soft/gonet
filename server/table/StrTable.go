package table

type StrTableCfg struct {
	ID     interface{} `json:"ID"`
	Module interface{} `json:"Module"`
	CN     interface{} `json:"CN"`
	EN     interface{} `json:"EN"`
}

type StrTableData map[string]StrTableCfg
