package etv3

import (
	"encoding/json"
	"gonet/actor"
	"gonet/common"
	"gonet/server/rpc"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"

	"golang.org/x/net/context"
)

//监控服务器
type Master struct {
	m_ServiceMap map[uint32]*common.ClusterInfo
	m_Client     *clientv3.Client
	m_Actor      actor.IActor
	common.IClusterInfo
}

//监控服务器
func (this *Master) Init(info common.IClusterInfo, Endpoints []string, pActor actor.IActor) {
	cfg := clientv3.Config{
		Endpoints: Endpoints,
	}

	etcdClient, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	this.m_ServiceMap = make(map[uint32]*common.ClusterInfo)
	this.m_Client = etcdClient
	this.BindActor(pActor)

	this.Start()
	this.IClusterInfo = info
}

func (this *Master) Start() {
	go this.InitService()
	go this.Run()
}

func (this *Master) BindActor(pActor actor.IActor) {
	this.m_Actor = pActor
}

func (this *Master) AddService(info *common.ClusterInfo) {
	this.m_Actor.SendMsg(rpc.RpcHead{}, "Cluster_Add", info)
	this.m_ServiceMap[info.Id()] = info
}

func (this *Master) delService(info *common.ClusterInfo) {
	delete(this.m_ServiceMap, info.Id())
	this.m_Actor.SendMsg(rpc.RpcHead{}, "Cluster_Del", info)
}

func NodeToService(val []byte) *common.ClusterInfo {
	info := &common.ClusterInfo{}
	err := json.Unmarshal([]byte(val), info)
	if err != nil {
		log.Print(err)
	}
	return info
}

func (this *Master) ListServices() {
	for key, info := range this.m_ServiceMap {
		log.Printf("ListServices: key %v >>>> %v", key, info)
	}
}

func (this *Master) Run() {
	wch := this.m_Client.Watch(context.Background(), ETCD_DIR+this.String(), clientv3.WithPrefix(), clientv3.WithPrevKV())
	//wch := this.m_Client.Watch(context.Background(), ETCD_DIR, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for v := range wch {
		for _, v1 := range v.Events {
			log.Printf("Run------type:%v kv:%v  prevKey:%v \n ", v1.Type, string(v1.Kv.Key), v1.PrevKv)

			if v1.Type.String() == "PUT" {
				info := NodeToService(v1.Kv.Value)
				this.AddService(info)
			} else {
				log.Printf("Warn: delete key: %s", v1.PrevKv.Value)
				info := NodeToService(v1.PrevKv.Value)
				this.delService(info)
			}
		}
	}
}
func (this *Master) InitService() {
	kv := clientv3.NewKV(this.m_Client)
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	getResp, err := kv.Get(ctx, ETCD_DIR, clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("get 失败：%s", err.Error())
	}
	log.Printf("%v", getResp.Kvs)

	for _, data := range getResp.Kvs {
		//Key := string(data.Key)

		info := NodeToService(data.Value)
		//fmt.Printf("key %s", Key)
		//fmt.Printf("info %v", info)

		this.AddService(info)
	}

	return
}
