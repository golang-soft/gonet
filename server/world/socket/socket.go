package socket

import (
	"github.com/golang/protobuf/proto"
	"gonet/network"
	"gonet/server/world/param"
	"gonet/server/world/wcluster"
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
type SocketId uint32

//type HandleEvent func(param router.Param)

var (
	sockets map[int]*Socket //socketid -- socket
)

func FetchSockets() map[int]*Socket {
	return sockets
}

func Emit(env string, data proto.Message) {
	sockets := FetchSockets()
	for _, socket := range sockets {
		socket.Emit(env, data)
	}
}

type Socket struct {
	id       uint32
	Data     *SocketData
	Room     *Room
	Client   *network.ClientSocket
	handlers map[string]interface{}
}

func (s *Socket) Emit(env string, message proto.Message) {
	wcluster.SendToClient(s.GetGateClusterId(), message, s.id)
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

func (s *Socket) On(event string, fc interface{}) {
	var para interface{}

	switch fc.(type) {
	case func(param.RoomParam):
		fc.(func(param.RoomParam))(para.(param.RoomParam))
	case func(param param.Param):
		fc.(func(param.Param))(para.(param.Param))
	case func(param param.UserParam):
		fc.(func(param.UserParam))(para.(param.UserParam))
	}
}

func GetOnlineUsers() []string {
	var list []string = make([]string, 0)
	sockets := FetchSockets()

	for _, socket := range sockets {
		list = append(list, socket.Data.User)
	}
	return list
}
