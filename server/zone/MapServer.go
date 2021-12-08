package zone

import (
	"database/sql"
	"gonet/base"
	"gonet/base/config"
	"gonet/base/ini"
	"gonet/base/system"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/network"
	"gonet/rpc"
)

type (
	ServerMgr struct {
		m_pService     *network.ServerSocket
		m_pCluster     *cluster.Cluster
		m_pWorldClient *network.ClientSocket
		m_Inited       bool
		m_config       ini.Config
		m_Log          base.CLog
	}

	IServerMgr interface {
		Init() bool
		InitDB() bool
		GetDB() *sql.DB
		GetLog() *base.CLog
		GetServer() *network.ServerSocket
	}

	Config struct {
		common.Server `yaml:"zone"`
		common.Etcd   `yaml:"etcd"`
		common.Nats   `yaml:"nats"`
	}
)

var (
	CONF   Config
	SERVER ServerMgr
)

func (this *ServerMgr) Init() bool {
	if this.m_Inited {
		return true
	}

	//初始化log文件
	this.m_Log.Init("world")
	//初始配置文件
	this.InitConfig(&CONF)

	ShowMessage := func() {
		this.m_Log.Debugf("**********************************************************")
		this.m_Log.Printf("\tServer Version:\t%s", base.BUILD_NO)
		this.m_Log.Printf("\tMapServerIP(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
		this.m_Log.Printf("\tEnv:\t\t%s", system.Args.Env)
		this.m_Log.Println("**********************************************************")
	}
	ShowMessage()

	//初始化socket
	this.m_pService = new(network.ServerSocket)
	this.m_pService.Init(CONF.Server.Ip, CONF.Server.Port)
	this.m_pService.Start()

	//集群管理
	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_ZONESERVER, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)

	var packet EventProcess
	packet.Init()
	this.m_pCluster.BindPacketFunc(packet.PacketFunc)

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
