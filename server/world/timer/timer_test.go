package timer

import (
	"gonet/server/world/gamedata"
	"testing"
)

func main() {

}

func TestTimer(t *testing.T) {
	s := gamedata.OnloadTimer
	s.Init()
	s.OnloadGameCheckTimer()
}
