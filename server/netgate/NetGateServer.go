package netgate

import (
	"gonet/base"
	"gonet/base/ini"
	"gonet/base/server"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/network"
	"gonet/rpc"
	"time"
)

type (
	ServerMgr struct {
		server.BaseServer
		//m_pService       *network.ServerSocket
		m_pService       *network.WebSocket
		m_Inited         bool
		m_config         ini.Config
		m_Log            base.CLog
		m_TimeTraceTimer *time.Ticker
		m_PlayerMgr      *PlayerManager
		m_pCluster       *cluster.Cluster
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
		common.Server `yaml:"netgate"`
		common.Etcd   `yaml:"etcd"`
		common.Nats   `yaml:"nats"`
	}
)

var (
	CONF   Config
	SERVER ServerMgr
)

func (this *ServerMgr) GetLog() *base.CLog {
	return &this.m_Log
}

//socket
//func (this *ServerMgr) GetServer() *network.ServerSocket{
// 	return this.m_pService
//}

//websocket
func (this *ServerMgr) GetServer() *network.WebSocket {
	return this.m_pService
}

func (this *ServerMgr) GetCluster() *cluster.Cluster {
	return this.m_pCluster
}

func (this *ServerMgr) GetPlayerMgr() *PlayerManager {
	return this.m_PlayerMgr
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

	ShowMessage := func() {
		this.m_Log.Println("**********************************************************")
		this.m_Log.Printf("\tNetGateServer Version:\t%s", base.BUILD_NO)
		this.m_Log.Printf("\tNetGateServerIP(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
		this.m_Log.Println("**********************************************************")
	}
	ShowMessage()

	//初始化socket
	//this.m_pService = new(network.ServerSocket)
	//this.m_pService.Init(CONF.Server.Ip, CONF.Server.Port)
	//this.m_pService.SetMaxPacketLen(base.MAX_CLIENT_PACKET)
	//this.m_pService.SetConnectType(network.CLIENT_CONNECT)
	////this.m_pService.Start()
	//packet := new(UserPrcoess)
	//packet.Init()
	//this.m_pService.BindPacketFunc(packet.PacketFunc)
	//this.m_pService.Start()

	//websocket
	this.m_pService = new(network.WebSocket)
	this.m_pService.Init(CONF.Server.Ip, CONF.Server.Port)
	this.m_pService.SetConnectType(network.CLIENT_CONNECT)
	//this.m_pService.Start()
	packet := new(UserPrcoess)
	packet.Init()
	this.m_pService.BindPacketFunc(packet.PacketFunc)
	this.m_pService.Start()
	//注册到集群服务器
	var packet1 EventProcess
	packet1.Init()
	this.m_pEventProcess = &packet1

	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_GATESERVER, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)
	this.m_pCluster.BindPacketFunc(packet1.PacketFunc)
	this.m_pCluster.BindPacketFunc(DispatchPacket)

	//初始玩家管理
	this.m_PlayerMgr = new(PlayerManager)
	this.m_PlayerMgr.Init()

	return false
}

func (this *ServerMgr) OnServerStart() {
	this.m_pService.Start()
}

func (this *ServerMgr) GetEventProcess() *EventProcess {
	return this.m_pEventProcess
}
