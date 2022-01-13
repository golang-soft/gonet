package main

import (
	"fmt"
	"gonet/network"
	"log"
	"time"
)

const (
	ROBOT_LOGIN   = 1
	ROBOT_GATEWAY = 2
)

type Robot struct {
	network.ClientSocket

	loginip   string
	loginport int
	account   string
	password  string
	session   string
	//1:login,2:gateway
	status uint32

	chClosed chan bool
	inittm   int64
	PACKET   *EventProcess
}

func NewRobot(ip string, port int, account string, password string) *Robot {
	robot := &Robot{
		loginip:   ip,
		loginport: port,
		account:   account,
		password:  password,
		chClosed:  make(chan bool, 1),
		inittm:    time.Now().Unix(),
	}
	ret := robot.Init(robot.loginip, robot.loginport)

	if !ret {
		log.Println("机器人", robot.account, "连接登陆服务器失败")
		return nil
	}

	robot.init()

	return robot
}
func (robot *Robot) init() {

	robot.Derived = robot
	robot.PACKET = new(EventProcess)
	robot.PACKET.Init()
	robot.PACKET.Robot = robot
	robot.BindPacketFunc(robot.PACKET.PacketFunc)
	//robot.msgHandler.Reg(&command.RetUserVerify{}, robot.onRetUserVerify)
	//robot.msgHandler.Reg(&command.RetUserLogin{}, robot.onRetUserLogin)
	//robot.msgHandler.Reg(&command.RetGatewayLogin{}, robot.onRetGatewayLogin)
	//robot.msgHandler.Reg(&command.TestBroadcastAll{}, robot.onTestBroadcastAll)

}
func (robot *Robot) Run() bool {

	robot.status = ROBOT_LOGIN

	host := fmt.Sprintf("%s:%d", robot.loginip, robot.loginport)
	if !robot.Start() {
		m_Log.Debugf("链接失败")
		return false
	}
	m_Log.Debugf("链接成功 %s", host)

	robot.inittm = time.Now().Unix()

	return true
}

func (robot *Robot) OnConnected() {

	log.Println("connected", robot.status)

	if robot.status == ROBOT_LOGIN {

		log.Println("请求登录服验证")
		robot.PACKET.LoginGate()

	} else if robot.status == ROBOT_GATEWAY {

		log.Println("请求网关验证")
	}
}

func (robot *Robot) Do() {

	//robot.msgQueue.Do(robot.msgHandler.Process)

}

func (robot *Robot) GetInitSec() uint32 {
	return uint32(time.Now().Unix() - robot.inittm)
}
