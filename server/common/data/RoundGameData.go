package data

type RoundGameData struct {
	Round        int
	StartTs      int64
	EndTs        int64
	FlagOwner    int32
	FlagUser     string
	FlagUpdateTs int64
	Part1score   float64
	Part1Ts      int64
	Part2Ts      int64
	Part2score   float64
}
