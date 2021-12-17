package table

type HelpTableCfg struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	StringID int    `json:"StringID"`
}

type HelpTableData map[string]HelpTableCfg
