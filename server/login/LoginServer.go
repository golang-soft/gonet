package login

import (
	"fmt"
	"gonet/base"
	"gonet/base/ini"
	"gonet/base/server"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/rpc"
	"net/http"
)

type (
	ServerMgr struct {
		server.BaseServer

		m_Inited      bool
		m_config      ini.Config
		m_Log         base.CLog
		m_FileMonitor common.IFileMonitor
		m_pCluster    *cluster.Cluster
	}

	IServerMgr interface {
		server.IBaseServer
		Init() bool
		GetLog() *base.CLog
		GetFileMonitor() common.IFileMonitor
	}

	Config struct {
		common.Server `yaml:"login"`
		common.Etcd   `yaml:"etcd"`
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
	this.m_Log.Init("login")
	//初始配置文件
	this.InitConfig(&CONF)

	//动态监控文件改变
	this.m_FileMonitor = &common.FileMonitor{}
	this.m_FileMonitor.Init()

	NETGATECONF.Init()
	http.HandleFunc("/listgates", GetNetGateS)
	addr := fmt.Sprintf("%s:%d", CONF.Server.Ip, CONF.Server.Port)
	http.ListenAndServe(addr, nil)

	//注册到集群服务器
	//var packet1 EventProcess
	//packet1.Init()
	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_LOGINSERVER, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)}, CONF.Etcd.Endpoints, "")
	//this.m_pCluster.BindPacketFunc(packet1.PacketFunc)
	//this.m_pCluster.BindPacketFunc(DispatchPacket)

	return false
}

func (this *ServerMgr) GetLog() *base.CLog {
	return &this.m_Log
}

func (this *ServerMgr) GetFileMonitor() common.IFileMonitor {
	return this.m_FileMonitor
}
