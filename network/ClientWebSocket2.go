package network

import (
	"fmt"
	"gonet/base"
	"gonet/grpc"
	"gonet/server/rpc"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type IClientWebSocket2 interface {
	ISocket
	OnConnected()
}

type ClientWebSocket2 struct {
	Socket
	m_nMaxClients int
	m_nMinClients int

	Derived IClientWebSocket2
}

func (this *ClientWebSocket2) Init(ip string, port int, params ...OpOption) bool {
	if this.m_nPort == port || this.m_sIP == ip {
		return false
	}

	this.Socket.Init(ip, port, params...)
	this.m_sIP = ip
	this.m_nPort = port
	fmt.Println(ip, port)
	return true
}

func (this *ClientWebSocket2) Start() bool {
	if this.m_sIP == "" {
		this.m_sIP = "127.0.0.1"
	}
	//var strRemote = fmt.Sprintf("%s:%d", this.m_sIP, this.m_nPort)
	if this.Connect() {
		go this.Run()

		if this.Derived != nil {
			this.Derived.OnConnected()
		}
		return true
	}
	//延迟，监听关闭
	//defer ln.Close()
	return false
}

//func (this *ClientWebSocket2) Start2(host string) bool {
//	if this.m_sIP == "" {
//		this.m_sIP = "127.0.0.1"
//	}
//
//	if this.Connect() {
//		go this.Run()
//
//		if this.Derived != nil {
//			this.Derived.OnConnected()
//		}
//		return true
//	}
//	//延迟，监听关闭
//	//defer ln.Close()
//	return false
//}

func (this *ClientWebSocket2) SendMsg(head rpc.RpcHead, funcName string, params ...interface{}) {
	buff := grpc.Marshal(head, funcName, params...)
	this.Send(head, buff)
}

func (this *ClientWebSocket2) Send(head rpc.RpcHead, buff []byte) int {
	defer func() {
		if err := recover(); err != nil {
			base.TraceCode(err)
		}
	}()

	if this.m_Conn == nil {
		return 0
	}

	n, err := this.m_Conn.Write(this.m_PacketParser.Write(buff))
	handleError(err)
	if n > 0 {
		return n
	}
	//this.m_Writer.Flush()
	return 0
}

func (this *ClientWebSocket2) Restart() bool {
	return true
}

func (this *ClientWebSocket2) Connect() bool {
	var host = fmt.Sprintf("%s:%d", this.m_sIP, this.m_nPort)
	wshost := "ws://" + host + "/ws"
	httphost := "http://" + host + "/"
	ws, err := websocket.Dial(wshost, "", httphost)
	if err != nil {
		log.Fatal(err)
	}
	this.SetConn(ws)
	return true
}

func (this *ClientWebSocket2) OnDisconnect() {
}

func (this *ClientWebSocket2) OnNetFail(int) {
	this.Stop()
	this.CallMsg("DISCONNECT", this.m_ClientId)
}

func (this *ClientWebSocket2) Run() bool {
	this.SetState(SSF_RUN)
	var buff = make([]byte, this.m_ReceiveBufferSize)
	loop := func() bool {
		defer func() {
			if err := recover(); err != nil {
				base.TraceCode(err)
			}
		}()

		if this.m_Conn == nil {
			return false
		}

		n, err := this.m_Conn.Read(buff)
		if err == io.EOF {
			fmt.Printf("远程链接：%s已经关闭！\n", this.m_Conn.RemoteAddr().String())
			this.OnNetFail(0)
			return false
		}
		if err != nil {
			handleError(err)
			this.OnNetFail(0)
			return false
		}
		if n > 0 {
			this.m_PacketParser.Read(buff[:n])
		}
		return true
	}

	for {
		if !loop() {
			break
		}
	}

	this.Close()
	return true
}
func (this *ClientWebSocket2) Close() {
	this.Clear()
	//this.GetConn().Close()
}
