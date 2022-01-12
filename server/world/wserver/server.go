package wserver

import (
	"gonet/server/world/gamedata"
	"gonet/server/world/router"
	"gonet/server/world/socket"
)

type Server struct {
	funcs map[string]interface{}
}

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

		if gamedata.GameCtrl.CheckUserEmpty(user) {
			return disConn("user account already connect")
		}

		socket.Data.User = user
		gamedata.GameCtrl.AddUser(user, socket.Data)
		socket.Emit("login", nil)
		return true
	})
}

func (this *Server) Start() {
	//初始化grpc
	LoadMiddleware(*this)
	this.funcs = make(map[string]interface{}, 0)
	this.On("connection", func(socket socket.Socket) bool {
		handleConn(socket)
		//处理user的协议
		router.HandleUser(socket)
		//处理helper协议
		router.HandleHelper(socket)
		//处理room协议
		router.HandleRoom(socket)
		return true
	})
}

func handleConn(socket1 socket.Socket) bool {
	//所有在线用户
	onlineUser := socket.GetOnlineUsers()
	//在线用户去重
	disConn := func(msg string) bool {
		socket1.Emitstr("conn_error", msg)
		account := socket1.Data.User
		gamedata.GGame.RemoveUser(account)
		socket1.Disconnect(true)
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
