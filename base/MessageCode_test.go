package base

import (
	"fmt"
	"hash/crc32"
	"testing"
)

func TestIEEE(t *testing.T) {
	ieee := crc32.ChecksumIEEE([]byte("c_a_loginrequest"))

	fmt.Printf("%v", ieee)
}
