package main

import (
	"fmt"
	"gonet/base"
	"gonet/base/config"
	"gonet/base/system"
	"gonet/common"
	"gonet/common/cluster/etv3"
	"gonet/network"
	common2 "gonet/server/common"
	"gonet/server/rpc"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

type (
	Config struct {
		common.MServer    `yaml:"mnetgate"`
		common.MWebsocket `yaml:"mnetgateweb"`
		common.Etcd       `yaml:"etcd"`
	}
)

var (
	m_Log base.CLog

	CONF   Config
	CLIENT *network.ClientWebSocketG
	//CLIENT *network.ClientSocket

	eventmanager *EventManager
)

func main() {
	m_Log.Init("client")
	common2.InitClient()

	//base.ReadConf("D:\\workspace-go\\gonet\\server\\client\\gonet.yaml", &CONF)
	config.Init(system.Args.Env, &CONF)

	//CLIENT = new(network.ClientWebSocket2)

	service := &etv3.Service{}
	thisip := "127.0.0.1"
	thisport := 3100

	for i := 0; i < len(CONF.MWebsocket.Websocket); i++ {
		sport := strings.Split(CONF.MWebsocket.Websocket[i], ":")[1]
		port, _ := strconv.Atoi(sport)
		ip := strings.Split(CONF.MWebsocket.Websocket[i], ":")[0]
		thisip = ip
		thisport, _ = strconv.Atoi(strings.Split(CONF.MWebsocket.Websocket[i], ":")[2])
		//index := this.GetIndex(this.m_pCluster.GetService().IpString())
		res := service.CheckExist(&common.ClusterInfo{Type: rpc.SERVICE_GATESERVER, Ip: ip, Port: int32(port)}, CONF.Etcd.Endpoints)
		if !res {
			break
		} else {
			continue
		}
	}
	//thisip = "192.168.1.206"
	//CLIENT.Init(thisip, thisport)
	//PACKET := new(EventProcess)
	//PACKET.Init()
	//CLIENT.BindPacketFunc(PACKET.PacketFunc)

	ShowMessage := func() {
		m_Log.Println("**********************************************************")
		m_Log.Printf("\tClient Version:\t%s", base.BUILD_NO)
		m_Log.Printf("\tClient(LAN):\t%s:%d", thisip, thisport)
		m_Log.Println("**********************************************************")
	}
	ShowMessage()

	host := fmt.Sprintf("%s:%d", thisip, thisport)
	//if !CLIENT.Start() {
	//	m_Log.Debugf("链接失败")
	//	return
	//}
	m_Log.Debugf("链接成功 %s", host)

	eventmanager = NewEventManager()
	eventmanager.Init()

	num := 1

	robotManager := NewRobotManager()
	robotManager.Add(num, thisip, thisport)

	for {
		robotManager.Do()
		time.Sleep(10 * time.Millisecond)
	}

	//PACKET.LoginGate()
	//PACKET.SendAttack()
	//PACKET.LoginGame()
	//PACKET.SendTest()

	//InitCmd()

	/*	for i := 0; i < 2; i++ {
		client := new(network.ClientWebSocket2)
		client.Init(CONF.Server.Ip, CONF.Server.Port)
		packet := new(EventProcess)
		packet.Init()
		client.BindPacketFunc(packet.PacketFunc)
		packet.Client = client
		if client.Start() {
			packet.LoginGate()
		}
	}*/
	//PACKET.LoginGame()
	//InitCmd()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Printf("client exit ------- signal:[%v]", s)
}
