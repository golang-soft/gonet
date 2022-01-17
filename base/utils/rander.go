package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand1 := &Rander{}
	num := rand1.RandInt(1, 100)

	fmt.Sprintf("num: %d", num)
}

type Rander struct {
}

func (this *Rander) Init() {
	rand.Seed(time.Now().UnixNano())
}

func (this *Rander) randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		temp = string(this.RandInt(65, 90))
		result.WriteString(temp)
		i++

	}
	return result.String()
}

func (this *Rander) RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
