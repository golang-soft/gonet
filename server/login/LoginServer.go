package login

import (
	"fmt"
	"gonet/base"
	"gonet/base/ini"
	"gonet/base/server"
	"gonet/common"
	"gonet/common/cluster"
	"gonet/rpc"
	"gonet/server/game"
	"gonet/server/message"
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
		common.PvpWeb `yaml:pvpweb`
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
	msg := &message.ReqServerVerify{}
	msg.Info = &message.ServerInfo{
		Id:   uint32(this.GetId()),
		Type: uint32(rpc.SERVICE_LOGINSERVER),
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
	this.m_Log.Init("login")
	//初始配置文件
	this.InitConfig(&CONF)

	//动态监控文件改变
	this.m_FileMonitor = &common.FileMonitor{}
	this.m_FileMonitor.Init()

	NETGATECONF.Init()
	this.InitCenterClient()

	//注册到集群服务器
	//var packet1 EventProcess
	//packet1.Init()
	this.m_pCluster = new(cluster.Cluster)
	this.m_pCluster.Init(&common.ClusterInfo{Type: rpc.SERVICE_LOGINSERVER, Ip: CONF.Server.Ip, Port: int32(CONF.Server.Port)}, CONF.Etcd.Endpoints, "")
	//this.m_pCluster.BindPacketFunc(packet1.PacketFunc)
	//this.m_pCluster.BindPacketFunc(DispatchPacket)

	http.HandleFunc("/listgates", GetNetGateS)
	http.HandleFunc("/test", Test)
	http.HandleFunc("/testworld", TestWorld)
	http.HandleFunc("/testgrpc", TestGrpc)

	//grpc接口
	http.HandleFunc("/grpcaddEquip", GrpcAddEquip)
	http.HandleFunc("/grpcaddHero", GrpcAddHero)
	http.HandleFunc("/grpcaddItem", GrpcAddItem)

	//grpc 测试接口
	http.HandleFunc("/grpcGetRooms", GrpcGetRooms)

	//http接口
	http.HandleFunc("/createPlayer", createPlayer)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/addHero", addHero)
	http.HandleFunc("/bindHeroEquip", bindHeroEquip)
	http.HandleFunc("/tdownHeroEquip", tdownHeroEquip)
	http.HandleFunc("/addItem", addItem)
	http.HandleFunc("/addEquip", addEquip)
	http.HandleFunc("/getGoodsByReduce", getGoodsByReduce)
	http.HandleFunc("/openBox", openBox)
	http.HandleFunc("/getToken", getToken)
	http.HandleFunc("/getNonce", getNonce)
	http.HandleFunc("/singVerify", singVerify)
	http.HandleFunc("/getLeaderboard", getLeaderboard)
	http.HandleFunc("/refreshUserLeaderboard", refreshUserLeaderboard)

	this.VerifyServer(CONF.Server.Ip, CONF.Server.Port)
	addr := fmt.Sprintf("%s:%d", CONF.Server.Ip, CONF.Server.Port)

	ShowMessage := func() {
		this.m_Log.Println("**********************************************************")
		this.m_Log.Printf("\tNetGateServer Version:\t%s", base.BUILD_NO)
		this.m_Log.Printf("\tNetGateServerIP(LAN):\t%s", addr)
		this.m_Log.Println("**********************************************************")
	}
	ShowMessage()

	http.ListenAndServe(addr, nil)

	return false
}

func (this *ServerMgr) GetLog() *base.CLog {
	return &this.m_Log
}

func (this *ServerMgr) GetFileMonitor() common.IFileMonitor {
	return this.m_FileMonitor
}

func (this *ServerMgr) GetCluster() *cluster.Cluster {
	return this.m_pCluster
}
