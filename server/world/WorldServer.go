package world

import (
	"database/sql"
	"github.com/golang/protobuf/proto"
	"gonet/base"
	"gonet/base/server"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/common/cluster/etv3"
	"gonet/db"
	"gonet/network"
	"gonet/rd"
	"gonet/rpc"
	"gonet/server/message"
	"log"
	"strconv"
	"strings"
)

type (
	ServerMgr struct {
		server.BaseServer

		m_pService   *network.ServerSocket
		m_pCluster   *cluster.Cluster
		m_pActorDB   *sql.DB
		m_Inited     bool
		m_Log        base.CLog
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
		GetPlayerRaft() *cluster.PlayerRaft
		SendToCenter(Id int64, ClusterId uint32, funcName string, params ...interface{})
	}

	Config struct {
		common.MServer `yaml:"mworld"`
		//common.Server    `yaml:"world"`
		common.Db        `yaml:"worldDB"`
		common.Redis     `yaml:"redis"`
		common.Etcd      `yaml:"etcd"`
		common.SnowFlake `yaml:"snowflake"`
		common.Raft      `yaml:"raft"`
		common.Nats      `yaml:"nats"`
		common.Center    `yaml:"center"`
	}
)

var (
	CONF   Config
	SERVER ServerMgr

	RdID int
)

type A struct {
	k int
}

func (this *ServerMgr) Init() bool {
	if this.m_Inited {
		return true
	}
	//test reload file
	/*file := &common.FileMonitor{}
	file.Init()
	file.AddFile("GONET_SERVER.CFG", func() {base.ReadConf("gonet.yaml", &CONF)})
	file.AddFile(data.SKILL_DATA_NAME, func() {
		data.SKILLDATA.Read()
	})*/

	//初始化log文件
	this.m_Log.Init("world")
	//初始配置文件
	//base.ReadConf("D:\\workspace-go\\gonet\\server\\bin\\gonet.yaml", &CONF)
	this.InitConfig(&CONF)

	this.m_Log.Println("正在初始化数据库连接...")
	if this.InitDB() {
		this.m_Log.Printf("[%s]数据库连接是失败...", CONF.Db.Name)
		log.Fatalf("[%s]数据库连接是失败...", CONF.Db.Name)
		return false
	}
	this.m_Log.Printf("[%s]数据库初始化成功!", CONF.Db.Name)

	if CONF.Redis.OpenFlag {
		rd.OpenRedisPool(CONF.Redis.Ip, CONF.Redis.Password)
	}

	//snowflake
	this.m_SnowFlake = cluster.NewSnowflake(CONF.SnowFlake.Endpoints)

	//playerraft
	this.m_PlayerRaft = cluster.NewPlayerRaft(CONF.Raft.Endpoints)

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
		res := service.CheckExist(&common.ClusterInfo{Type: rpc.SERVICE_WORLDSERVER, Ip: ip, Port: int32(port)}, CONF.Etcd.Endpoints)
		if !res {
			continue
		} else {
			break
		}
	}

	this.m_pService = new(network.ServerSocket)
	this.m_pService.Init(thisip, thisport)
	this.m_pService.Start()

	//本身world集群管理
	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_WORLDSERVER, Ip: thisip, Port: int32(thisport)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)

	var packet EventProcess
	packet.Init()

	var centerProcess CenterProcess
	centerProcess.Init()

	this.m_pCluster.BindPacketFunc(packet.PacketFunc)
	this.m_pCluster.BindPacketFunc(centerProcess.PacketFunc)

	ShowMessage := func() {
		this.m_Log.Println("**********************************************************")
		this.m_Log.Printf("\tWorldServer Version:\t%s", base.BUILD_NO)
		this.m_Log.Printf("\tWorldServerIP(LAN):\t%s:%d", thisip, thisport)
		this.m_Log.Printf("\tActorDBServer(LAN):\t%s", CONF.Db.Ip)
		this.m_Log.Printf("\tActorDBName:\t\t%s", CONF.Db.Name)
		this.m_Log.Println("**********************************************************")
	}
	ShowMessage()

	//this.SendToCenter(1, 0, "LoginCenter")
	return false
}

func (this *ServerMgr) GetIndex(ipstring string) int {
	for index, value := range CONF.MServer.Endpoints {
		if value == ipstring {
			return index
		}
	}
	return 0
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

func (this *ServerMgr) GetPlayerRaft() *cluster.PlayerRaft {
	return this.m_PlayerRaft
}

//发送account
func SendToAccount(funcName string, params ...interface{}) {
	head := rpc.RpcHead{DestServerType: rpc.SERVICE_ACCOUNTSERVER, SendType: rpc.SEND_BALANCE, SrcClusterId: SERVER.GetCluster().Id()}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

//发送给客户端
func SendToClient(clusterId uint32, packet proto.Message) {
	pakcetHead := packet.(message.Packet).GetPacketHead()
	if pakcetHead != nil {
		SERVER.GetCluster().SendMsg(rpc.RpcHead{DestServerType: rpc.SERVICE_GATESERVER, ClusterId: clusterId, Id: pakcetHead.Id}, "", proto.MessageName(packet), packet)
	}
}

//--------------发送给地图----------------------//
func SendToZone(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_ZONESERVER, SrcClusterId: SERVER.GetCluster().Id()}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

//--------------发送给中央服----------------------//
func (this *ServerMgr) SendToCenter(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_CENTERSERVER, SrcClusterId: SERVER.GetCluster().Id(), SendType: rpc.SEND_BALANCE}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}
