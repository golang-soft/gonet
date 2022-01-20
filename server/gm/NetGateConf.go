package login

import (
	"gonet/base"
	"gonet/base/config"
	"gonet/base/ini"
	"net/http"
	"sync"
)

type (
	NetGateConf struct {
		m_config ini.Config
		m_Locker *sync.RWMutex
	}
)

var (
	NETGATECONF NetGateConf
)

func (this *NetGateConf) Init() bool {
	this.m_Locker = &sync.RWMutex{}
	this.Read()
	path := config.GetConfigPath("NETGATES.CFG")
	SERVER.GetFileMonitor().AddFile(path, this.Read)
	return true
}

func (this *NetGateConf) Read() {
	this.m_Locker.Lock()
	path := config.GetConfigPath("NETGATES.CFG")
	this.m_config.Read(path)
	this.m_Locker.Unlock()
}

func (this *NetGateConf) GetNetGates(Arena string) []string {
	this.m_Locker.RLock()
	arenas := this.m_config.Get6(Arena, "NetGates", ",")
	this.m_Locker.RUnlock()
	return arenas
}

func GetNetGateS(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	size := "1"
	arenas := NETGATECONF.GetNetGates(size)
	nLen := len(arenas)
	if nLen > 0 {
		nIndex := base.RAND.RandI(0, nLen-1)
		w.Write([]byte(arenas[nIndex]))
		return
	}

	w.Write([]byte(""))
}
