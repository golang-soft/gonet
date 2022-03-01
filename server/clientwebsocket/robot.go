package main

import (
	"fmt"
	"gonet/base/utils"
	"gonet/network"
	"log"
	"time"
)

const (
	ROBOT_LOGIN   = 1
	ROBOT_GATEWAY = 2
	ROBOT_PLAYING = 3
)

type Robot struct {
	network.ClientWebSocket2

	loginip   string
	loginport int
	account   string
	password  string
	session   string
	//1:login,2:gateway,3:playing
	status uint32

	chClosed     chan bool
	inittm       int64
	eventProcess *EventProcess
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
	robot.eventProcess = new(EventProcess)
	robot.eventProcess.Init()
	robot.eventProcess.Robot = robot
	robot.BindPacketFunc(robot.eventProcess.PacketFunc)

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
		//robot.PACKET.LoginGate()
		e := eventmanager.GetEvent("LoginGateEvent")
		if e != nil {
			(*e).DoEvent(robot.eventProcess)
		}
	} else if robot.status == ROBOT_GATEWAY {
		log.Println("登录完毕")
	} else if robot.status == ROBOT_PLAYING {
		log.Println("玩家正在游戏中....")
	}

}

func (robot *Robot) Do() {
	if robot.status == ROBOT_PLAYING {
		rander := &utils.Rander{}
		rander.Init()

		randId := rander.RandInt(0, eventmanager.Count())
		//log.Printf("随机数 : %d", randId)

		time.Sleep(time.Duration(100 * time.Millisecond))
		//var e *IBaseEvent = eventmanager.GetEvent("AttackEvent")
		var e *IBaseEvent = eventmanager.GetEventById(randId)
		if e != nil {
			(*e).SendEvent(e, robot.eventProcess)
		}
	}
}

func (robot *Robot) GetInitSec() uint32 {
	return uint32(time.Now().Unix() - robot.inittm)
}
