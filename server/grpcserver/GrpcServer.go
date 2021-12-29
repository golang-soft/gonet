package grpcserver

import (
	"gonet/base"
	"gonet/base/config"
	"gonet/base/ini"
	"gonet/base/server"
	"gonet/base/system"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/network"
	"gonet/rpc"
	"gonet/server/game"
	"gonet/server/smessage"
)

type (
	ServerMgr struct {
		server.BaseServer

		m_pService     *network.ServerSocket
		m_pCluster     *cluster.Cluster
		m_pWorldClient *network.ClientSocket
		m_Inited       bool
		m_config       ini.Config
		m_Log          base.CLog
	}

	IServerMgr interface {
		server.IBaseServer

		Init() bool
		InitDB() bool
		GetLog() *base.CLog
		GetServer() *network.ServerSocket
	}

	Config struct {
		common.Server `yaml:"grpcserver"`
		common.Etcd   `yaml:"etcd"`
		common.Nats   `yaml:"nats"`
		common.Center `yaml:"center"`
	}
)

var (
	CONF   Config
	SERVER ServerMgr
)

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
		Type: uint32(rpc.SERVICE_GRPCSERVER),
		Ip:   thisip,
		Port: uint32(thisport),
	}
	this.SendToCenter(1, 0, "ReqServerVerify", msg)
}

//--------------发送给中央服----------------------//
func (this *ServerMgr) SendToCenter(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_CENTERSERVER, SrcClusterId: SERVER.GetCluster().Id(), SendType: rpc.SEND_BOARD_CAST}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

func (this *ServerMgr) Init() bool {
	if this.m_Inited {
		return true
	}

	//初始化log文件
	this.m_Log.Init("grpcserver")
	//初始配置文件
	this.InitConfig(&CONF)

	ShowMessage := func() {
		this.m_Log.Debug("**********************************************************")
		this.m_Log.Debugf("\tServer Version:\t%s", base.BUILD_NO)
		this.m_Log.Debugf("\tGrpcServerIP(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
		this.m_Log.Debugf("\tEnv:\t\t%s", system.Args.Env)
		this.m_Log.Debugf("**********************************************************")
	}
	ShowMessage()
	this.InitCenterClient()

	//初始化socket
	this.m_pService = new(network.ServerSocket)
	this.m_pService.Init(CONF.Server.Ip, CONF.Server.Port)
	this.m_pService.Start()

	//集群管理
	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_GRPCSERVER, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)

	var packet EventProcess
	packet.Init()
	this.m_pCluster.BindPacketFunc(packet.PacketFunc)

	this.VerifyServer(CONF.Server.Ip, CONF.Server.Port)
	return false
}

func (this *ServerMgr) InitConfig(data interface{}) bool {
	config.Init(system.Args.Env, data)
	return true
}

func (this *ServerMgr) GetServer() *network.ServerSocket {
	return this.m_pService
}

func (this *ServerMgr) GetLog() *base.CLog {
	return &this.m_Log
}
func (this *ServerMgr) GetCluster() *cluster.Cluster {
	return this.m_pCluster
}
