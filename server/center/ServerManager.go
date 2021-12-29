package center

import (
	"gonet/server/smessage"
	"sync"
)

type ServerManager struct {
	mutex     sync.Mutex
	serverMap map[uint32]*smessage.ServerInfo
}

func NewServerManager() *ServerManager {
	mgr := &ServerManager{
		serverMap: make(map[uint32]*smessage.ServerInfo),
	}
	return mgr
}

func (this *ServerManager) GetById(id uint32) *smessage.ServerInfo {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if _, ok := this.serverMap[id]; !ok {
		return nil
	}

	return this.serverMap[id]
}

func (this *ServerManager) GetByType(tp uint32) []*smessage.ServerInfo {

	this.mutex.Lock()
	defer this.mutex.Unlock()

	serverlist := make([]*smessage.ServerInfo, 0, 0)
	for _, task := range this.serverMap {

		if task.Type != tp {
			continue
		}
		serverlist = append(serverlist, task)
	}

	return serverlist
}

func (this *ServerManager) UniqueAdd(task *smessage.ServerInfo) bool {

	if this.GetById(task.Id) != nil {
		SERVER.m_Log.Println("重复添加服务器 %s", task.Id)
		return false
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.serverMap[task.Id] = task

	return true
}

func (mgr *ServerManager) UniqueRemove(task *smessage.ServerInfo) {

	mgr.mutex.Lock()
	defer mgr.mutex.Unlock()

	delete(mgr.serverMap, task.Id)
}

func (mgr *ServerManager) DebugServerList() {
	for _, server := range mgr.serverMap {
		SERVER.m_Log.Debugf(">>>>>>> %d, %d, %s, %d", server.Id, server.Type, server.Ip, server.Port)
	}

}
