package table

type MapCfg struct {
	ID        int   `json:"ID"`
	Map_ID    int   `json:"Map_ID"`
	Map_X     int   `json:"Map_X"`
	Map_Y     int   `json:"Map_Y"`
	Flag      []int `json:"Flag"`
	Red       []int `json:"Red"`
	Blue      []int `json:"Blue"`
	Score     int   `json:"Score"`
	Max_Score int   `json:"Max_Score"`
	Race_Time int   `json:"Race_Time"`
	Duration  int   `json:"Duration"`
}

type MapData map[int]*MapCfg
