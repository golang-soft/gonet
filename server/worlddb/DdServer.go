package worlddb

import (
	"database/sql"
	"gonet/base"
	"gonet/base/ini"
	"gonet/base/server"
	"gonet/base/system"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/db"
	"gonet/network"
	"gonet/rpc"
	"log"
)

type (
	ServerMgr struct {
		server.BaseServer

		m_pService     *network.ServerSocket
		m_pCluster     *cluster.Cluster
		m_pWorldClient *network.ClientSocket
		m_pPlayerMgr   *PlayerMgr
		m_pActorDB     *sql.DB
		m_Inited       bool
		m_config       ini.Config
		m_Log          base.CLog
	}

	IServerMgr interface {
		server.IBaseServer

		Init() bool
		InitDB() bool
		GetDB() *sql.DB
		GetLog() *base.CLog
		GetServer() *network.ServerSocket
	}

	Config struct {
		common.Server `yaml:"worlddbserver"`
		common.Db     `yaml:"worldDB"`
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
	//base.ReadConf("D:\\workspace-go\\gonet\\server\\bin\\gonet.yaml", &CONF)
	this.InitConfig(&CONF)

	ShowMessage := func() {
		this.m_Log.Debugf("**********************************************************")
		this.m_Log.Printf("\tServer Version:\t%s", base.BUILD_NO)
		this.m_Log.Printf("\tDbServerIP(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
		this.m_Log.Printf("\tActorDBServer(LAN):\t%s", CONF.Db.Ip)
		this.m_Log.Printf("\tActorDBName:\t\t%s", CONF.Db.Name)
		this.m_Log.Printf("\tEnv:\t\t%s", system.Args.Env)
		this.m_Log.Println("**********************************************************")
	}
	ShowMessage()

	this.m_Log.Println("正在初始化数据库连接...")
	if this.InitDB() {
		this.m_Log.Printf("[%s]数据库连接是失败...", CONF.Db.Name)
		log.Fatalf("[%s]数据库连接是失败...", CONF.Db.Name)
		return false
	}
	this.m_Log.Printf("[%s]数据库初始化成功!", CONF.Db.Name)

	//初始化socket
	this.m_pService = new(network.ServerSocket)
	this.m_pService.Init(CONF.Server.Ip, CONF.Server.Port)
	this.m_pService.Start()

	//var packet EventProcess
	//packet.Init()
	//this.m_pService.BindPacketFunc(packet.PacketFunc)

	this.m_pPlayerMgr = new(PlayerMgr)
	this.m_pPlayerMgr.Init()

	//集群管理
	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_WORLDDBSERVER, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)

	var packet EventProcess
	packet.Init()
	this.m_pCluster.BindPacketFunc(packet.PacketFunc)
	this.m_pCluster.BindPacketFunc(this.m_pPlayerMgr.PacketFunc)

	return true
}

func (this *ServerMgr) InitDB() bool {
	this.m_pActorDB = db.OpenDB(CONF.Db)
	err := this.m_pActorDB.Ping()
	return err != nil
}

func (this *ServerMgr) GetDB() *sql.DB {
	return this.m_pActorDB
}

func (this *ServerMgr) GetServer() *network.ServerSocket {
	return this.m_pService
}

func (this *ServerMgr) GetLog() *base.CLog {
	return &this.m_Log
}
