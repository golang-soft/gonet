package base

import (
	"fmt"
	"testing"
)

func TestAAA(t *testing.T) {
	var dh Dh
	dh.Init()

	fmt.Printf("%v", dh.ShareKey())
	fmt.Printf("%v", dh.PubKey())
}
