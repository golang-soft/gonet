package players

var Players map[int64]string

func Init() {
	if Players == nil {
		Players = make(map[int64]string)
	}
}
func GetPlayerByAccountId(accountId int64) string {
	p, ok := Players[accountId]
	if ok {
		return p
	}
	return ""
}
