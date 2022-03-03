package network

import (
	"crypto/tls"
	"fmt"
	"gonet/base"
	"gonet/common/timer"
	"gonet/server/rpc"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

type IWebSocketClientG interface {
	ISocket
}

type WebSocketClientG struct {
	Socket
	m_wConn    *websocket.Conn
	m_pServer  *WebSocketG
	m_SendChan chan []byte //对外缓冲队列
	m_TimerId  *int64
}

var (
	//是否支持心跳超时检测
	checkHeard bool = false
)

func (this *WebSocketClientG) Init(ip string, port int, params ...OpOption) bool {
	this.Socket.Init(ip, port, params...)
	return true
}

func (this *WebSocketClientG) Start() bool {
	if this.m_pServer == nil {
		return false
	}

	if this.m_nConnectType == CLIENT_CONNECT {
		this.m_SendChan = make(chan []byte, MAX_SEND_CHAN)
		this.m_TimerId = new(int64)
		*this.m_TimerId = int64(this.m_ClientId)
		timer.RegisterTimer(this.m_TimerId, (HEART_TIME_OUT/3)*time.Second, func() {
			this.Update()
		})
	}

	if this.m_PacketFuncList.Len() == 0 {
		this.m_PacketFuncList = this.m_pServer.m_PacketFuncList
	}

	go this.Run()
	if this.m_nConnectType == CLIENT_CONNECT {
		go this.SendLoop()
	}

	return true
}

func (this *WebSocketClientG) Send(head rpc.RpcHead, buff []byte) int {
	defer func() {
		if err := recover(); err != nil {
			base.TraceCode(err)
		}
	}()

	if this.m_nConnectType == CLIENT_CONNECT { //对外链接send不阻塞
		select {
		case this.m_SendChan <- buff:
		default: //网络太卡,tcp send缓存满了并且发送队列也满了
			this.OnNetFail(1)
		}
	} else {
		return this.DoSend(buff)
	}
	return len(buff)
}

func (this *WebSocketClientG) DoSend(buff []byte) int {
	if this.m_wConn == nil {
		return 0
	}

	err := this.m_wConn.WriteMessage(websocket.BinaryMessage, this.m_PacketParser.Write(buff))
	handleError(err)
	if len(buff) > 0 {
		return len(buff)
	}

	return 0
}

func (this *WebSocketClientG) OnNetFail(error int) {
	this.Stop()

	if this.m_nConnectType == CLIENT_CONNECT { //netgate对外格式统一
		stream := base.NewBitStream(make([]byte, 32), 32)
		stream.WriteInt(int(DISCONNECTINT), 32)
		stream.WriteInt(int(this.m_ClientId), 32)
		this.HandlePacket(stream.GetBuffer())
	} else {
		this.CallMsg("DISCONNECT", this.m_ClientId)
	}
	if this.m_pServer != nil {
		this.m_pServer.DelClinet(this)
	}
}

func (this *WebSocketClientG) OnNetClose() {
	err := this.m_wConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		m_Log.Errorf("write close error : %v", err)
		return
	}
}

func (this *WebSocketClientG) Close() {
	if this.m_nConnectType == CLIENT_CONNECT {
		//close(this.m_SendChan)
		timer.StopTimer(this.m_TimerId)
	}
	this.Socket.Close()
	if this.m_pServer != nil {
		this.m_pServer.DelClinet(this)
	}
}

func (this *WebSocketClientG) Run() bool {
	//var buff = make([]byte, this.m_ReceiveBufferSize)
	this.SetState(SSF_RUN)
	loop := func() bool {
		defer func() {
			if err := recover(); err != nil {
				base.TraceCode(err)
			}
		}()

		if this.m_wConn == nil {
			return false
		}

		_, buff, err := this.m_wConn.ReadMessage()
		if err == io.EOF {
			fmt.Printf("远程链接：%s已经关闭！\n", this.m_wConn.RemoteAddr().String())
			this.OnNetFail(0)
			return false
		}
		if err == io.ErrUnexpectedEOF {
			fmt.Printf("远程链接：%s已经关闭！\n", this.m_wConn.RemoteAddr().String())
			this.OnNetFail(0)
			return false
		}
		if err != nil {
			handleError(err)
			this.OnNetClose()
			this.OnNetFail(0)
			return false
		}
		if len(buff) > 0 {
			this.m_PacketParser.Read(buff[:])
		}
		this.SetLastHeardTime(int(time.Now().Unix()) + HEART_TIME_OUT)
		base.GLOG.Printf("调整心跳时间: %s", strconv.Itoa(this.m_HeartTime))
		return true
	}

	for {
		if !loop() {
			break
		}
	}

	this.Close()
	fmt.Printf("%s关闭连接", this.m_sIP)
	return true
}

// heart
func (this *WebSocketClientG) Update() bool {
	now := int(time.Now().Unix())
	if this.m_HeartTime > 0 && this.m_HeartTime < now && checkHeard {
		this.OnNetFail(2)
		return false
	}

	return true
}

func (this *WebSocketClientG) SendLoop() bool {
	for {
		defer func() {
			if err := recover(); err != nil {
				base.TraceCode(err)
			}
		}()

		select {
		case buff := <-this.m_SendChan:
			if buff == nil { //信道关闭
				return false
			} else {
				this.DoSend(buff)
			}
		}
	}
	return true
}

func (this *WebSocketClientG) Connect() bool {
	server := "127.0.0.1:31700"
	scheme := "ws"
	u := url.URL{Scheme: scheme, Host: server, Path: "/"}
	d := websocket.DefaultDialer
	d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c, _, err := d.Dial(u.String(), nil)
	if err != nil {
		fmt.Printf("连接失败：%s \n", err)
		return false
	}

	this.SetConn(c)

	fmt.Printf("连接成功：%s\n", this.m_wConn.RemoteAddr().String())
	return true
}

func (this *WebSocketClientG) SetConn(conn *websocket.Conn) {
	this.m_wConn = conn
}

func (this *WebSocketClientG) SetLastHeardTime(time int) {
	this.m_HeartTime = time
}

func (this *WebSocketClientG) GetLastHeardTime() int {
	return this.m_HeartTime
}
