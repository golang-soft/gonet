package center

import (
	"sync"
)

var sonce sync.Once

type ServerIdManager struct {
	curId int64
	mutex sync.Mutex
}

var serverIdManagerInstance *ServerIdManager

func GetServerIdManager() *ServerIdManager {
	sonce.Do(func() {
		serverIdManagerInstance = new(ServerIdManager)
	})
	return serverIdManagerInstance
}

func (this *ServerIdManager) getCurServerId() int64 {
	return this.curId
}

func (this *ServerIdManager) getNextServerId() int64 {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.curId++

	return this.curId
}
