package center

import "sync"

var once sync.Once

type UserManager struct {
	onlineCount int64
}

var instance *UserManager

func GetUserManager() *UserManager {
	once.Do(func() {
		instance = new(UserManager)
	})
	return instance
}

func (this *UserManager) getOnlineUserCount() int64 {
	return this.onlineCount
}

func (this *UserManager) Add() {
	this.onlineCount++
}

func (this *UserManager) Sub() {
	this.onlineCount--
}
