package main

import (
	"fmt"
	"gonet/base"
	"gonet/common"
	"gonet/network"
	"gonet/server/message"
	"os"
	"os/signal"
)

type (
	Config struct {
		common.Server `yaml:"netgate"`
	}
)

var (
	m_Log base.CLog

	CONF   Config
	CLIENT *network.ClientWebSocket2
	//CLIENT *network.WebSocketClient
)

func main() {
	message.InitClient()

	m_Log.Init("client")

	base.ReadConf("D:\\workspace-go\\gonet\\server\\client\\gonet.yaml", &CONF)

	ShowMessage := func() {
		m_Log.Println("**********************************************************")
		m_Log.Printf("\tClient Version:\t%s", base.BUILD_NO)
		m_Log.Printf("\tClient(LAN):\t%s:%d", CONF.Server.Ip, CONF.Server.Port)
		m_Log.Println("**********************************************************")
	}
	ShowMessage()

	CLIENT = new(network.ClientWebSocket2)
	CLIENT.Init(CONF.Server.Ip, CONF.Server.Port)
	PACKET = new(EventProcess)
	PACKET.Init()
	CLIENT.BindPacketFunc(PACKET.PacketFunc)
	PACKET.Client = CLIENT
	host := fmt.Sprintf("%s:%d", CONF.Server.Ip, CONF.Server.Port)
	if !CLIENT.Start(host) {
		m_Log.Debugf("链接失败")
		return
	}
	m_Log.Debugf("链接成功 %s", host)

	PACKET.LoginGate()
	//PACKET.LoginGame()
	//PACKET.SendTest()

	InitCmd()

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
