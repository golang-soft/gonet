package netgate

import (
	"gonet/base"
	"gonet/base/ini"
	"gonet/base/server"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/common/cluster/etv3"
	"gonet/network"
	common2 "gonet/server/common"
	"gonet/server/game"
	"gonet/server/rpc"
	"gonet/server/smessage"
	"strconv"
	"strings"
	"time"
)

type (
	ServerMgr struct {
		server.BaseServer
		m_pTcpService    *network.ServerSocket
		m_pService       *network.WebSocket
		m_pServiceG      *network.WebSocketG
		m_Inited         bool
		m_config         ini.Config
		m_Log            base.CLog
		m_TimeTraceTimer *time.Ticker
		m_PlayerMgr      *PlayerManager
		m_pCluster       *cluster.Cluster
		m_pClusterWs     *cluster.Cluster
		m_pEventProcess  *EventProcess
	}

	IServerMgr interface {
		server.IBaseServer

		Init() bool
		GetLog() *base.CLog
		GetServer() *network.ServerSocket
		GetCluster() *cluster.Service
		GetPlayerMgr() *PlayerManager
		OnServerStart()
	}

	Config struct {
		common.MServer    `yaml:"mnetgate"`
		common.MWebsocket `yaml:"mnetgateweb"`
		common.Etcd       `yaml:"etcd"`
		common.Nats       `yaml:"nats"`
		common.Center     `yaml:"center"`
	}
)

var (
	CONF          Config
	SERVER        ServerMgr
	IsWebsocket   bool   = true
	WebsocketMode string = "gorilla" //gorilla 或者 golang
)

func (this *ServerMgr) GetLog() *base.CLog {
	return &this.m_Log
}

//socket
func (this *ServerMgr) GetServer() *network.ServerSocket {
	return this.m_pTcpService
}

//websocket
func (this *ServerMgr) GetWebSocketServer() *network.WebSocket {
	return this.m_pService
}
func (this *ServerMgr) GetWebSocketServerG() *network.WebSocketG {
	return this.m_pServiceG
}

func (this *ServerMgr) GetCluster() *cluster.Cluster {
	if IsWebsocket {
		return this.m_pClusterWs
	}
	return this.m_pCluster
}

func (this *ServerMgr) GetPlayerMgr() *PlayerManager {
	return this.m_PlayerMgr
}

func (this *ServerMgr) InitCenterClient() bool {
	//初始化grpc连接
	this.M_pGrpcClient = game.NewGrpcClient()
	this.M_pGrpcClient.ConnectToServer(CONF.Center.GrpcPort)
	this.SetId(this.M_pGrpcClient.ReqServerId())
	return true
}

func (this *ServerMgr) VerifyServer(thisip string, thisport int) {
	msg := &smessage.ReqServerVerify{}
	msg.Info = &smessage.ServerInfo{
		Id:   uint32(this.GetId()),
		Type: uint32(rpc.SERVICE_GATESERVER),
		Ip:   thisip,
		Port: uint32(thisport),
	}
	this.SendToCenter(1, 0, "ReqServerVerify", msg)
}

//--------------发送给中央服----------------------//
func (this *ServerMgr) SendToCenter(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, DestClusterId: ClusterId, DestServerType: rpc.SERVICE_CENTERSERVER, SrcClusterId: SERVER.GetCluster().Id(), SendType: rpc.SEND_BOARD_CAST}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

func (this *ServerMgr) InitSocket(thisip string, thisport int) bool {
	this.m_pTcpService = new(network.ServerSocket)
	this.m_pTcpService.Init(thisip, thisport)
	this.m_pTcpService.SetMaxPacketLen(base.MAX_CLIENT_PACKET)
	this.m_pTcpService.SetConnectType(network.CLIENT_CONNECT)
	//this.m_pService.Start()
	packet := new(UserPrcoess)
	packet.Init()
	this.m_pTcpService.BindPacketFunc(packet.PacketFunc)
	this.m_pTcpService.Start()
	common2.Init()
	return true
}

func (this *ServerMgr) InitWebsocket(thisip string, thisport int) bool {
	this.m_pService = new(network.WebSocket)
	this.m_pService.Init(thisip, thisport)
	this.m_pService.SetMaxPacketLen(base.MAX_CLIENT_PACKET)
	this.m_pService.SetConnectType(network.CLIENT_CONNECT)
	//this.m_pService.Start()
	packet := new(UserPrcoess)
	packet.Init()
	this.m_pService.BindPacketFunc(packet.PacketFunc)
	this.m_pService.Start()
	common2.Init()
	return true
}
func (this *ServerMgr) InitWebsocketG(thisip string, thisport int) bool {
	this.m_pServiceG = new(network.WebSocketG)
	this.m_pServiceG.Init(thisip, thisport)
	this.m_pServiceG.SetMaxPacketLen(base.MAX_CLIENT_PACKET)
	this.m_pServiceG.SetConnectType(network.CLIENT_CONNECT)
	//this.m_pService.Start()
	packet := new(UserPrcoess)
	packet.Init()
	this.m_pServiceG.BindPacketFunc(packet.PacketFunc)
	this.m_pServiceG.Start()
	common2.Init()
	return true
}

func (this *ServerMgr) Init() bool {
	if this.m_Inited {
		return true
	}

	//初始化log文件
	this.m_Log.Init("netgate")
	//初始配置文件
	//base.ReadConf("D:\\workspace-go\\gonet\\server\\bin\\gonet.yaml", &CONF)
	this.InitConfig(&CONF)

	//etcd 的处理
	service := &etv3.Service{}
	thisip := "127.0.0.1"
	thisport := 31300

	if !IsWebsocket {
		for i := 0; i < len(CONF.MServer.Endpoints); i++ {
			sport := strings.Split(CONF.MServer.Endpoints[i], ":")[1]
			port, _ := strconv.Atoi(sport)
			ip := strings.Split(CONF.MServer.Endpoints[i], ":")[0]
			thisip = ip
			thisport = port
			//index := this.GetIndex(this.m_pCluster.GetService().IpString())
			res := service.CheckExist(&common.ClusterInfo{Type: rpc.SERVICE_GATESERVER, Ip: ip, Port: int32(port)}, CONF.Etcd.Endpoints)
			if !res {
				continue
			} else {
				break
			}
		}

		this.InitCenterClient()
		//初始化socket
		this.InitSocket(thisip, thisport)

	}

	thiswip := "127.0.0.1"
	thiswport := 31300

	if IsWebsocket {

		for i := 0; i < len(CONF.MWebsocket.Websocket); i++ {
			sport := strings.Split(CONF.MWebsocket.Websocket[i], ":")[1]
			port, _ := strconv.Atoi(sport)
			ip := strings.Split(CONF.MWebsocket.Websocket[i], ":")[0]
			thiswip = ip
			thiswport = port
			//index := this.GetIndex(this.m_pCluster.GetService().IpString())
			res := service.CheckExist(&common.ClusterInfo{Type: rpc.SERVICE_GATESERVER, Ip: ip, Port: int32(port)}, CONF.Etcd.Endpoints)
			if !res {
				continue
			} else {
				break
			}
		}

		//websocket 暂时不用
		//this.InitWebsocket(thiswip, thiswport)
		this.InitWebsocketG(thiswip, thiswport)
	}

	//注册到集群服务器
	var packet1 EventProcess
	packet1.Init()
	this.m_pEventProcess = &packet1

	if !IsWebsocket {
		this.m_pCluster = new(cluster.Cluster)
		this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_GATESERVER, Ip: thisip, Port: int32(thisport)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)
		this.m_pCluster.BindPacketFunc(packet1.PacketFunc)
		this.m_pCluster.BindPacketFunc(DispatchPacket)
	}

	if IsWebsocket {
		this.m_pClusterWs = new(cluster.Cluster)
		this.m_pClusterWs.Init(&common.ClusterInfo{Type: rpc.SERVICE_GATESERVER, Ip: thiswip, Port: int32(thiswport)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)
		this.m_pClusterWs.BindPacketFunc(packet1.PacketFunc)
		this.m_pClusterWs.BindPacketFunc(DispatchPacket)
	}

	//初始玩家管理
	this.m_PlayerMgr = new(PlayerManager)
	this.m_PlayerMgr.Init()

	ShowMessage := func() {
		this.m_Log.Println("**********************************************************")
		this.m_Log.Printf("\tNetGateServer Version:\t%s", base.BUILD_NO)
		if IsWebsocket {
			this.m_Log.Printf("\tNetGateServerIP(LAN):\t%s:%d", thiswip, thiswport)
		} else {
			this.m_Log.Printf("\tNetGateServerIP(LAN):\t%s:%d", thisip, thisport)
		}

		this.m_Log.Println("**********************************************************")
	}
	ShowMessage()
	if IsWebsocket {
		this.VerifyServer(thiswip, thiswport)
	} else {
		this.VerifyServer(thisip, thisport)
	}

	return false
}

//func (this *ServerMgr) OnServerStart() {
//	this.m_pService.Start()
//}

func (this *ServerMgr) GetEventProcess() *EventProcess {
	return this.m_pEventProcess
}

func (this *ServerMgr) CheckIsWebsocket() bool {
	return IsWebsocket
}

func (this *ServerMgr) WebsocketModeIsGorilla() bool {
	if WebsocketMode == "gorilla" {
		return true
	}
	return false
}
