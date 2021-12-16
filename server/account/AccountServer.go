package account

import (
	"database/sql"
	"gonet/base"
	"gonet/base/ini"
	"gonet/base/server"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/common/cluster/etv3"
	"gonet/db"
	"gonet/network"
	"gonet/rpc"
	"gonet/server/message"
	"log"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

type (
	ServerMgr struct {
		server.BaseServer
		m_pService   *network.ServerSocket
		m_pCluster   *cluster.Cluster
		m_pActorDB   *sql.DB
		m_Inited     bool
		m_config     ini.Config
		m_Log        base.CLog
		m_AccountMgr *AccountMgr
		m_SnowFlake  *cluster.Snowflake
		m_PlayerRaft *cluster.PlayerRaft
	}

	IServerMgr interface {
		server.IBaseServer
		Init() bool
		InitDB() bool
		GetDB() *sql.DB
		GetLog() *base.CLog
		GetServer() *network.ServerSocket
		GetCluster() *cluster.Cluster
		GetAccountMgr() *AccountMgr
		GetPlayerRaft() *cluster.PlayerRaft
	}

	Config struct {
		//common.Server    `yaml:"account"`
		common.MServer   `yaml:"maccount"`
		common.Db        `yaml:"accountDB"`
		common.Etcd      `yaml:"etcd"`
		common.SnowFlake `yaml:"snowflake"`
		common.Raft      `yaml:"raft"`
		common.Nats      `yaml:"nats"`
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
	this.m_Log.Init("account")
	//初始配置文件
	//base.ReadConf("D:\\workspace-go\\gonet\\server\\bin\\gonet.yaml", &CONF)
	this.InitConfig(&CONF)

	//etcd 的处理
	service := &etv3.Service{}
	thisip := "127.0.0.1"
	thisport := 31300

	for i := 0; i < len(CONF.MServer.Endpoints); i++ {
		sport := strings.Split(CONF.MServer.Endpoints[i], ":")[1]
		port, _ := strconv.Atoi(sport)
		ip := strings.Split(CONF.MServer.Endpoints[i], ":")[0]
		thisip = ip
		thisport = port
		//index := this.GetIndex(this.m_pCluster.GetService().IpString())
		res := service.CheckExist(&common.ClusterInfo{Type: rpc.SERVICE_ACCOUNTSERVER, Ip: ip, Port: int32(port)}, CONF.Etcd.Endpoints)
		if !res {
			continue
		} else {
			break
		}
	}

	this.m_Log.Println("正在初始化数据库连接...")
	if this.InitDB() {
		this.m_Log.Printf("[%s]数据库连接是失败...", CONF.Db.Name)
		log.Fatalf("[%s]数据库连接是失败...", CONF.Db.Name)
		return false
	}
	this.m_Log.Printf("[%s]数据库初始化成功!", CONF.Db.Name)

	//初始化socket
	this.m_pService = new(network.ServerSocket)
	this.m_pService.Init(thisip, thisport)
	this.m_pService.Start()

	//账号管理类
	this.m_AccountMgr = new(AccountMgr)
	this.m_AccountMgr.Init()

	//本身账号集群管理
	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_ACCOUNTSERVER, Ip: thisip, Port: int32(thisport)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)

	var packet EventProcess
	packet.Init()
	this.m_pCluster.BindPacketFunc(packet.PacketFunc)
	this.m_pCluster.BindPacketFunc(this.m_AccountMgr.PacketFunc)

	//snowflake
	//this.m_SnowFlake = cluster.NewSnowflake(CONF.SnowFlake.Endpoints)

	//playerraft
	this.m_PlayerRaft = cluster.NewPlayerRaft(CONF.Raft.Endpoints)

	ShowMessage := func() {
		this.m_Log.Println("**********************************************************")
		this.m_Log.Printf("\tAccountServer Version:\t%s", base.BUILD_NO)
		this.m_Log.Printf("\tAccountServerIP(LAN):\t%s:%d", thisip, thisport)
		this.m_Log.Printf("\tActorDBServer(LAN):\t%s", CONF.Db.Ip)
		this.m_Log.Printf("\tActorDBName:\t\t%s", CONF.Db.Name)
		this.m_Log.Println("**********************************************************")
	}
	ShowMessage()
	return false
}

func (this *ServerMgr) InitDB() bool {
	this.m_pActorDB = db.OpenDB(CONF.Db)
	err := this.m_pActorDB.Ping()
	return err != nil
}

func (this *ServerMgr) GetDB() *sql.DB {
	return this.m_pActorDB
}

func (this *ServerMgr) GetLog() *base.CLog {
	return &this.m_Log
}

func (this *ServerMgr) GetServer() *network.ServerSocket {
	return this.m_pService
}

func (this *ServerMgr) GetCluster() *cluster.Cluster {
	return this.m_pCluster
}

func (this *ServerMgr) GetAccountMgr() *AccountMgr {
	return this.m_AccountMgr
}

func (this *ServerMgr) GetPlayerRaft() *cluster.PlayerRaft {
	return this.m_PlayerRaft
}

func KickWorldPlayer(accountId int64) {
	BoardCastToWorld("G_ClientLost", accountId)
}

//发送world
func SendToWorld(ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{ClusterId: ClusterId, DestServerType: rpc.SERVICE_WORLDSERVER, SrcClusterId: SERVER.GetCluster().Id()}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

//广播world
func BoardCastToWorld(funcName string, params ...interface{}) {
	head := rpc.RpcHead{DestServerType: rpc.SERVICE_WORLDSERVER, SendType: rpc.SEND_BOARD_CAST, SrcClusterId: SERVER.GetCluster().Id()}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

//发送到客户端
func SendToClient(head rpc.RpcHead, packet proto.Message) {
	pakcetHead := packet.(message.Packet).GetPacketHead()
	if pakcetHead != nil {
		head.DestServerType = rpc.SERVICE_GATESERVER
		head.Id = pakcetHead.Id
	}
	SERVER.GetCluster().SendMsg(head, "", proto.MessageName(packet), packet)
}
