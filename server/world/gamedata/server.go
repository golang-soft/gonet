package gamedata

import (
	"github.com/golang/protobuf/proto"
	"gonet/server/world/socket"
)

var (
	sockets map[int]*socket.Socket //socketid -- socket
)

func FetchSockets() map[int]*socket.Socket {
	return sockets
}

func Emit(env string, data proto.Message) {
	sockets := FetchSockets()
	for _, socket := range sockets {
		socket.Emit(env, data)
	}
}

type MyFunc func(socket socket.Socket) bool
type Server struct {
	funcs map[string]interface{}
}

//var io = &Server{}

func LoadMiddleware(io Server) {
	io.Use(func(socket socket.Socket) bool {

		disConn := func(msg string) bool {
			//socket.Emit("conn_error", msg)
			socket.Disconnect(true)
			return true
		}

		var user = socket.Data.User
		if user != "" {
			return disConn("user account not found")
		}

		if GameCtrl.CheckUserEmpty(user) {
			return disConn("user account already connect")
		}

		socket.Data.User = user
		GameCtrl.AddUser(user, socket.Data)
		socket.Emit("login", nil)
		return true
	})
}

func (this *Server) Start() {
	//初始化grpc
	LoadMiddleware(*this)
	//加载定时器
	OnloadTimer.OnloadGameCheckTimer()
	this.funcs = make(map[string]interface{}, 0)
	this.On("connection", func(socket socket.Socket) bool {
		handleConn(socket)
		//处理user的协议
		//handleUser(socket)
		//处理helper协议
		//handleHelper(socket)
		//处理room协议
		//handleRoom(socket, io)
		return true
	})
}

func handleConn(socket socket.Socket) bool {
	//所有在线用户
	onlineUser := GetOnlineUsers()
	////在线用户去重
	disConn := func(msg string) bool {
		socket.Emitstr("conn_error", msg)
		account := socket.Data.User
		GGame.RemoveUser(account)
		socket.Disconnect(true)
		return true
	}
	if len(onlineUser) != len(onlineUser) && len(onlineUser) > 0 {
		return disConn("user account have exist")
	}
	return false
}

func (this *Server) Use(f func(socket socket.Socket) bool) {
	if this.funcs == nil {
		this.funcs = make(map[string]interface{})
	}
	this.funcs["1"] = f
}

func (this *Server) On(s string, f func(socketdd socket.Socket) bool) {
	this.funcs[s] = f
}
