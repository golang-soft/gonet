package main

import (
	"fmt"
	"sync"
	"time"
)

type RobotManager struct {
	robotList map[string]*Robot

	ip   string
	port int
	num  uint32

	mutex sync.Mutex
}

func NewRobotManager() *RobotManager {
	mgr := &RobotManager{
		robotList: make(map[string]*Robot),
		num:       0,
	}

	return mgr
}

func (mgr *RobotManager) Do() {

	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	for _, robot := range mgr.robotList {
		robot.Do()
	}

	//for account, robot := range mgr.robotList {
	//if robot.GetInitSec() >= 100 {
	//	robot.Close()
	//	delete(mgr.robotList, account)
	//}
	//}

	for i := len(mgr.robotList); i < int(mgr.num); i++ {
		account := fmt.Sprint(time.Now().UnixNano() + int64(i))
		password := "123456"
		mgr.AddOne(account, password)
	}

	//log.Println("当前机器人数量：", len(mgr.robotList))

}

func (mgr *RobotManager) Add(num int, ip string, port int) {

	mgr.num += uint32(num)
	mgr.ip = ip
	mgr.port = port

	for i := 0; i != num; i++ {
		account := fmt.Sprint(time.Now().UnixNano() + int64(i))
		password := "123456"

		mgr.AddOne(account, password)
	}
}

func (mgr *RobotManager) AddOne(account string, password string) {

	robot := NewRobot(mgr.ip, mgr.port, account, password)
	go robot.Run()

	mgr.robotList[account] = robot
}
