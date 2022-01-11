package socket

import (
	"github.com/golang/protobuf/proto"
	"gonet/network"
)

type SocketData struct {
	User   string
	RoomId int64
	Round  int
	Part   int32
}
type Room struct {
	Name string
}
type SocketId string

type Socket struct {
	id     SocketId
	Data   *SocketData
	Room   *Room
	Client *network.ClientSocket
}

func (s *Socket) Emit(env string, message proto.Message) {
	//world.SendToClient(s.GetGateClusterId(), message)
}

func (s *Socket) Emitstr(env string, msg string) {
}

func (s *Socket) GetGateClusterId() uint32 {
	return 0
}

func (s *Socket) Join(battleName string) {
	s.Room.Name = battleName
}

func (s *Socket) Leave(sprintf string) {
	s.Room.Name = ""
}

func (s *Socket) Disconnect(b bool) {
	if s.Client != nil {
		s.Client.Close()
	}
}
