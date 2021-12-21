package et

import (
	"fmt"
	"gonet/rpc"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/client"
	"golang.org/x/net/context"
)

const (
	PLAYER_DIR   = "player/"
	OFFLINE_TIME = 15 * 60
)

//publish
type (
	PlayerRaft struct {
		m_KeysAPI      client.KeysAPI
		m_PlayerLocker *sync.RWMutex
		m_PlayerMap    map[int64]*rpc.PlayerClusterInfo
	}
)

//初始化pub
func (this *PlayerRaft) Init(endpoints []string) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}
	this.m_KeysAPI = client.NewKeysAPI(etcdClient)
	this.m_PlayerLocker = &sync.RWMutex{}
	this.m_PlayerMap = map[int64]*rpc.PlayerClusterInfo{}
	this.Start()
	this.InitPlayers()
}

func (this *PlayerRaft) Start() {
	go this.Run()
}

func (this *PlayerRaft) Publish(info *rpc.PlayerClusterInfo) bool {
	//info.LeaseId = int64(info.Id)
	//key := PLAYER_DIR + fmt.Sprintf("%d", info.Id)
	//data, _ := proto.Marshal(info)
	//_, err := this.m_KeysAPI.Set(context.Background(), key, string(data), &client.SetOptions{
	//	TTL: ttl_time, PrevExist: client.PrevNoExist, NoValueOnSuccess: true,
	//})
	//return err == nil
	return false
}

func (this *PlayerRaft) Lease(Id int64) error {
	key := PLAYER_DIR + fmt.Sprintf("%d", Id)
	_, err := this.m_KeysAPI.Set(context.Background(), key, "", &client.SetOptions{
		TTL: ttl_time, Refresh: true, NoValueOnSuccess: true,
	})
	return err
}

func (this *PlayerRaft) addPlayer(info *rpc.PlayerClusterInfo) {
	this.m_PlayerLocker.Lock()
	pPlayer, bOk := this.m_PlayerMap[info.Id]
	if !bOk {
		this.m_PlayerMap[info.Id] = info
	} else {
		*pPlayer = *info
	}
	this.m_PlayerLocker.Unlock()
}

func (this *PlayerRaft) delPlayer(info *rpc.PlayerClusterInfo) {
	this.m_PlayerLocker.Lock()
	delete(this.m_PlayerMap, int64(info.Id))
	this.m_PlayerLocker.Unlock()
}

func (this *PlayerRaft) GetPlayer(Id int64) *rpc.PlayerClusterInfo {
	this.m_PlayerLocker.RLock()
	pPlayer, bEx := this.m_PlayerMap[Id]
	this.m_PlayerLocker.RUnlock()
	if bEx {
		return pPlayer
	}
	return nil
}

// subscribe
func (this *PlayerRaft) Run() {
	watcher := this.m_KeysAPI.Watcher(PLAYER_DIR, &client.WatcherOptions{
		Recursive: true,
	})

	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			log.Println("Error watch service:", err)
			continue
		}
		if res.Action == "expire" {
			info := NodeToPlayer([]byte(res.PrevNode.Value))
			this.delPlayer(info)
		} else if res.Action == "set" {
			info := NodeToPlayer([]byte(res.Node.Value))
			this.addPlayer(info)
		} else if res.Action == "delete" {
			info := NodeToPlayer([]byte(res.Node.Value))
			this.delPlayer(info)
		}
	}
}

func (this *PlayerRaft) InitPlayers() {
	resp, err := this.m_KeysAPI.Get(context.Background(), PLAYER_DIR, &client.GetOptions{Recursive: true})
	if err == nil && (resp != nil && resp.Node != nil) {
		for _, v := range resp.Node.Nodes {
			info := NodeToPlayer([]byte(v.Value))
			this.addPlayer(info)
		}
	}
}

func NodeToPlayer(val []byte) *rpc.PlayerClusterInfo {
	//info := &rpc.PlayerClusterInfo{}
	//err := proto.Unmarshal([]byte(val), info)
	//if err != nil {
	//	log.Print(err)
	//}
	//return info
	return nil
}
