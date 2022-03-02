package network

import (
	"flag"
	"fmt"
	"gonet/base/logger"
	"gonet/server/rpc"

	"github.com/gorilla/websocket"
	"gonet/grpc"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域访问
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type IWebSocketG interface {
	ISocket

	AssignClientId() uint32
	GetClientById(uint32) *WebSocketClientG
	LoadClient() *WebSocketClientG
	AddClinet(*websocket.Conn, string, int) *WebSocketClientG
	DelClinet(*WebSocketClientG) bool
	StopClient(uint32)
}

type WebSocketG struct {
	Socket
	m_nClientCount int
	m_nMaxClients  int
	m_nMinClients  int
	m_nIdSeed      uint32
	m_ClientList   map[uint32]*WebSocketClientG
	m_ClientLocker *sync.RWMutex
	m_Lock         sync.Mutex
}

func (this *WebSocketG) Init(ip string, port int, params ...OpOption) bool {
	this.Socket.Init(ip, port, params...)
	this.m_ClientList = make(map[uint32]*WebSocketClientG)
	this.m_ClientLocker = &sync.RWMutex{}
	this.m_sIP = ip
	this.m_nPort = port
	return true
}

func (this *WebSocketG) Start() bool {
	if this.m_sIP == "" {
		this.m_sIP = "127.0.0.1"
	}

	go func() {
		var strRemote = fmt.Sprintf("%s:%d", this.m_sIP, this.m_nPort)
		addr := flag.String("addr", strRemote, "http service address")
		//var strRemote1 = fmt.Sprintf("%d", this.m_sIP, this.m_nPort)
		//http.Handle("/ws", websocket.Handler(this.wserverRoutine))
		//http.Handle("/socket.io/", websocket.Handler(this.wserverRoutine))
		http.HandleFunc("/ws", this.wserverRoutine2)
		//http.HandleFunc("/socket.io/", helloHandler)

		err := http.ListenAndServe(*addr, nil)
		if err != nil {
			fmt.Errorf("WebSocketG ListenAndServe:%v", err)
		}
	}()

	fmt.Printf("WebSocket 启动监听，等待链接！\n")
	return true
}

func (this *WebSocketG) AssignClientId() uint32 {
	return atomic.AddUint32(&this.m_nIdSeed, 1)
}

func (this *WebSocketG) GetClientById(id uint32) *WebSocketClientG {
	this.m_ClientLocker.RLock()
	client, exist := this.m_ClientList[id]
	this.m_ClientLocker.RUnlock()
	if exist == true {
		return client
	}

	return nil
}

func (this *WebSocketG) AddClinet(tcpConn *websocket.Conn, addr string, connectType int) *WebSocketClientG {
	pClient := this.LoadClient()
	if pClient != nil {
		pClient.Init("", 0)
		pClient.m_pServer = this
		pClient.m_ReceiveBufferSize = this.m_ReceiveBufferSize
		pClient.SetMaxPacketLen(this.GetMaxPacketLen())
		pClient.m_ClientId = this.AssignClientId()
		pClient.m_sIP = addr
		pClient.SetConn(tcpConn.UnderlyingConn())
		pClient.SetConnectType(connectType)
		this.m_ClientLocker.Lock()
		this.m_ClientList[pClient.m_ClientId] = pClient
		this.m_ClientLocker.Unlock()

		this.m_nClientCount++
		return pClient
	} else {
		log.Printf("%s", "无法创建客户端连接对象")
	}
	return nil
}

func (this *WebSocketG) DelClinet(pClient *WebSocketClientG) bool {
	this.m_ClientLocker.Lock()
	delete(this.m_ClientList, pClient.m_ClientId)
	this.m_ClientLocker.Unlock()
	return true
}

func (this *WebSocketG) StopClient(id uint32) {
	pClient := this.GetClientById(id)
	if pClient != nil {
		pClient.Stop()
	}
}

func (this *WebSocketG) LoadClient() *WebSocketClientG {
	s := &WebSocketClientG{}
	return s
}

func (this *WebSocketG) Send(head rpc.RpcHead, buff []byte) int {
	pClient := this.GetClientById(head.SocketId)
	if pClient != nil {
		pClient.Send(head, buff)
	}
	return 0
}

func (this *WebSocketG) SendMsg(head rpc.RpcHead, funcName string, params ...interface{}) {
	pClient := this.GetClientById(head.SocketId)
	if pClient != nil {
		pClient.Send(head, grpc.Marshal(head, funcName, params...))
	}
}

func (this *WebSocketG) Restart() bool {
	return true
}
func (this *WebSocketG) Connect() bool {
	return true
}
func (this *WebSocketG) Disconnect(bool) bool {
	return true
}

func (this *WebSocketG) OnNetFail(int) {
}

func (this *WebSocketG) Close() {
	this.Clear()
}

func (this *WebSocketG) wserverRoutine(conn *websocket.Conn) {
	fmt.Printf("客户端：%s已连接！\n", conn.RemoteAddr().String())
	this.handleConn(conn, conn.RemoteAddr().String())
}

func (this *WebSocketG) handleConn(tcpConn *websocket.Conn, addr string) bool {
	if tcpConn == nil {
		return false
	}

	//tcpConn.PayloadType = websocket.BinaryFrame
	pClient := this.AddClinet(tcpConn, addr, this.m_nConnectType)
	if pClient == nil {
		return false
	}

	pClient.Start()
	return true
}

func (this *WebSocketG) wserverRoutine2(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello"))
	//收到http请求(upgrade),完成websocket协议转换
	//在应答的header中放上upgrade:websoket
	var (
		conn *websocket.Conn
		err  error
		//msgType int
		//data []byte
	)
	if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
		//报错了，直接返回底层的websocket链接就会终断掉
		logger.Debug("error[%v]", err)
		return
	}
	//type SS struct {
	//	user string
	//}
	//var s SS
	//websocket.ReadJSON(conn,&s )
	//得到了websocket.Conn长连接的对象，实现数据的收发
	//for {
	//	//Text(json), Binary
	//	//if _, data, err = conn.ReadMessage(); err != nil {
	//	if _, data, err = conn.ReadMessage(); err != nil {
	//		//报错关闭websocket
	//		goto ERR
	//	}
	//	//发送数据，判断返回值是否报错
	//	if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
	//		//报错了
	//		goto ERR
	//	}
	//}

	fmt.Printf("客户端：%s已连接！\n", conn.RemoteAddr().String())
	this.handleConn(conn, conn.RemoteAddr().String())

	//error的标签
	//ERR:
	//	conn.Close()

}

func (this *WebSocketG) helloHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello World!"))
	fmt.Printf("客户端：%s已连接！\n", r.RemoteAddr)

}
