package common

import "strconv"

func GetRoundKey(user string, round int) string {
	return strconv.Itoa(round) + ":" + user
}
