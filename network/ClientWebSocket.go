package network

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gonet/base"
	"gonet/rpc"
	"io"
	"net/url"
)

type IClientWebSocket interface {
	ISocket
}

type ClientWebSocket struct {
	Socket
	m_nMaxClients int
	m_nMinClients int
}

func (this *ClientWebSocket) Init(ip string, port int, params ...OpOption) bool {
	if this.m_nPort == port || this.m_sIP == ip {
		return false
	}

	this.Socket.Init(ip, port, params...)
	this.m_sIP = ip
	this.m_nPort = port
	fmt.Println(ip, port)
	return true
}

func (this *ClientWebSocket) Start() bool {
	if this.m_sIP == "" {
		this.m_sIP = "127.0.0.1"
	}

	if this.Connect() {
		go this.Run()
	}
	//延迟，监听关闭
	//defer ln.Close()
	return true
}

func (this *ClientWebSocket) SendMsg(head rpc.RpcHead, funcName string, params ...interface{}) {
	buff := rpc.Marshal(head, funcName, params...)
	this.Send(head, buff)
}

func (this *ClientWebSocket) Send(head rpc.RpcHead, buff []byte) int {
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

func (this *ClientWebSocket) Restart() bool {
	return true
}

func (this *ClientWebSocket) Connect() bool {
	connectStr := "ws"

	server := "127.0.0.1:31700"
	u := url.URL{Scheme: connectStr, Host: server, Path: "/ws"}
	d := websocket.DefaultDialer
	//d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c, _, err := d.Dial(u.String(), nil)
	if err != nil {
		fmt.Printf("连接失败：%s \n", err)
		return false
	}
	this.SetConn(c.UnderlyingConn())
	fmt.Printf("%s 连接成功，请输入信息！\n", connectStr)
	this.CallMsg("COMMON_RegisterRequest")
	return true
}

func (this *ClientWebSocket) OnDisconnect() {
}

func (this *ClientWebSocket) OnNetFail(int) {
	this.Stop()
	this.CallMsg("DISCONNECT", this.m_ClientId)
}

func (this *ClientWebSocket) Run() bool {
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
