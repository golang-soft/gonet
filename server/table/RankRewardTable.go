package table

type RankRewardTableCfg struct {
	ID   int `json:"ID"`
	Type int `json:"Type"`
	Rank int `json:"Rank"`
	DVT  int `json:"DVT"`
}

type RankRewardTableData map[string]RankRewardTableCfg
