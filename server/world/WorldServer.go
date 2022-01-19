package world

import (
	"database/sql"
	"gonet/base"
	_ "gonet/base/redis"
	"gonet/base/server"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/common/cluster/etv3"
	"gonet/db"
	"gonet/network"
	"gonet/rd"
	_ "gonet/server/common/mredis"
	"gonet/server/game"
	"gonet/server/glogger"
	"gonet/server/rpc"
	"gonet/server/smessage"
	"gonet/server/world/datafnc"
	"gonet/server/world/gamedata"
	"gonet/server/world/helper"
	"gonet/server/world/redisInstnace"
	"gonet/server/world/socket"
	"gonet/server/world/table"
	"gonet/server/world/wcluster"
	"gonet/server/world/wserver"
	"log"
	"strconv"
	"strings"
)

type (
	ServerMgr struct {
		server.BaseServer

		m_pService *network.ServerSocket
		//m_pCluster     *cluster.Cluster
		m_pActorDB *sql.DB
		m_Inited   bool
		//M_Log        base.CLog
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
		Type: uint32(rpc.SERVICE_WORLDSERVER),
		Ip:   thisip,
		Port: uint32(thisport),
	}
	wcluster.SendToCenter(1, 0, "ReqServerVerify", msg)
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
	glogger.M_Log.Init("world")
	//初始配置文件
	//base.ReadConf("D:\\workspace-go\\gonet\\server\\bin\\gonet.yaml", &CONF)
	this.InitConfig(&CONF)

	table.Init()
	datafnc.Init()
	glogger.M_Log.Printf("初始化配置表数据成功!")

	glogger.M_Log.Println("正在初始化数据库连接...")
	if this.InitDB() {
		glogger.M_Log.Printf("[%s]数据库连接是失败...", CONF.Db.Name)
		log.Fatalf("[%s]数据库连接是失败...", CONF.Db.Name)
		return false
	}
	glogger.M_Log.Printf("[%s]数据库初始化成功!", CONF.Db.Name)

	helper.InitConst()

	if CONF.Redis.OpenFlag {
		rd.OpenRedisPool(CONF.Redis.Ip, CONF.Redis.Password)
		redisInstnace.Init(CONF.Redis.Prefix,
			CONF.Redis.Ip,
			CONF.Redis.Port,
			CONF.Redis.Password,
			CONF.Redis.Db,
		)

	}
	this.InitCenterClient()

	//{
	//	value := this.M_pRedisClient.HGetAll("user:round:basic:10001:999")
	//	mapdata, _ := value.Result()
	//	this.M_Log.Debugf(">>>> %s", mapdata["user"])
	//
	//	value1 := this.M_pRedisClient.HGet("user:round:basic:10001:999", "user")
	//	this.M_Log.Debugf(">>>>《《《《 %s", value1.Val())
	//}

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
	m_pCluster := new(cluster.Cluster)
	m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_WORLDSERVER, Ip: thisip, Port: int32(thisport)}, CONF.Etcd.Endpoints, CONF.Nats.Endpoints)
	wcluster.SetCluster(m_pCluster)

	var eventProcess EventProcess
	eventProcess.Init()

	redisInstnace.InitRedis()

	wserver.NewGameServer()
	wserver.GameServer.Start()
	socket.Init()

	//var io = &wserver.Server{}
	//io.Start()

	gamedata.OnloadTimer.Init()

	var centerProcess CenterProcess
	centerProcess.Init()
	var gameprocess GameProcess
	gameprocess.Init()
	//this.m_pCluster.BindPacketFunc(eventProcess.PacketFunc)
	//this.m_pCluster.BindPacketFunc(centerProcess.PacketFunc)
	m_pCluster.BindPacketFunc(eventProcess.PacketFunc)
	m_pCluster.BindPacketFunc(centerProcess.PacketFunc)
	m_pCluster.BindPacketFunc(gameprocess.PacketFunc)

	ShowMessage := func() {
		glogger.M_Log.Debugf("**********************************************************")
		glogger.M_Log.Printf("\tWorldServer Version:\t%s", base.BUILD_NO)
		glogger.M_Log.Printf("\tWorldServerIP(LAN):\t%s:%d", thisip, thisport)
		glogger.M_Log.Printf("\tActorDBServer(LAN):\t%s", CONF.Db.Ip)
		glogger.M_Log.Printf("\tActorDBName:\t\t%s", CONF.Db.Name)
		glogger.M_Log.Println("**********************************************************")
	}
	ShowMessage()

	this.VerifyServer(thisip, thisport)

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
	return &glogger.M_Log
}

func (this *ServerMgr) GetServer() *network.ServerSocket {
	return this.m_pService
}

func (this *ServerMgr) GetPlayerRaft() *cluster.PlayerRaft {
	return this.m_PlayerRaft
}
