package socket

import (
	"github.com/golang/protobuf/proto"
	"gonet/network"
	"gonet/server/glogger"
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
	sockets   map[int64]*Socket //accountid -- socket
	socketids map[int64]*Socket //socketid -- socket
)

func Init() {
	sockets = make(map[int64]*Socket)
	socketids = make(map[int64]*Socket)
}

func FetchSockets() map[int64]*Socket {
	return sockets
}

func Emit(env string, data proto.Message) {
	sockets := FetchSockets()
	for _, socket := range sockets {
		socket.Emit(env, data)
	}
}

type Socket struct {
	id        uint32
	AccountId int64
	Data      *SocketData
	Room      *Room
	Client    *network.ClientSocket
	handlers  map[string]interface{}
}

func GetOnlineUsers() []string {
	var list []string = make([]string, 0)
	sockets := FetchSockets()

	for _, socket := range sockets {
		list = append(list, socket.Data.User)
	}
	return list
}

func (s *Socket) Emit(env string, message proto.Message) bool {
	//var para interface{}
	fc := s.handlers[env]
	if fc == nil {
		return false
	}

	wcluster.SendToClient(s.GetGateClusterId(), message, s.id)
	return false
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
	if event != "" {
		s.handlers[event] = fc
	}
}

func (s *Socket) Handle(event string, sock Socket, para interface{}) {
	fc := s.handlers[event]
	if fc != nil {
		switch fc.(type) {
		case func(param.RoomParam):
			fc.(func(param.RoomParam))(para.(param.RoomParam))
		case func(param param.Param):
			fc.(func(param.Param))(para.(param.Param))
		case func(param param.UserParam):
			fc.(func(param.UserParam))(para.(param.UserParam))
		case func(sock Socket, userParam param.UserParam):
			fc.(func(sock Socket, userParam param.UserParam))(sock, para.(param.UserParam))
		default:
			{
				glogger.M_Log.Debugf("未知的函数  %v", fc)
			}
		}
	}
}

func AddOne(socketid uint32, accountid int64, data *SocketData, Room *Room, Client *network.ClientSocket) {
	s := &Socket{
		id:        socketid,
		AccountId: accountid,
		Data:      data,
		Room:      Room,
		Client:    Client,
		handlers:  make(map[string]interface{}),
	}

	sockets[accountid] = s
	socketids[int64(socketid)] = s
}

func GetOne(accountid int64) *Socket {
	return sockets[accountid]
}

func GetOneBySocketid(socketid uint32) *Socket {
	return socketids[int64(socketid)]
}

func RemoveOne(accountid int64) bool {
	socket, ok := sockets[accountid]
	if ok {
		if socket != nil {
			delete(sockets, accountid)
			delete(socketids, int64(socket.id))
			return true
		}
	}
	return false
}
