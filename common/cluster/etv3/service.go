package etv3

import (
	"encoding/json"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"gonet/common"
	"log"
	"time"
)

const (
	ETCD_DIR = "server/"
)

type (
	//注册服务器
	Service struct {
		*common.ClusterInfo
		m_Client  *clientv3.Client
		m_Lease   clientv3.Lease
		m_LeaseId clientv3.LeaseID
		isRun     bool
	}

	IService interface {
		Run()
		Grant()
		Put()
		KeepAlive()
		Lease()
		Revoke()
		Init(info *common.ClusterInfo, endpoints []string) bool
		Start() bool
		GetKey()
		GetValue(key string) []*mvccpb.KeyValue
	}
)

func (this *Service) Run() {
	if this.isRun {
		this.Grant()
		this.Put()
		//for {
		this.Lease()
		time.Sleep(time.Second * 1)
		//}
	} else {
		log.Printf("is not run !!!!!")
	}
}

//设置租约时间
func (this *Service) Grant() {
	leaseResp, _ := this.m_Lease.Grant(context.Background(), 2)
	this.m_LeaseId = leaseResp.ID
}

//通过租约put
func (this *Service) Put() {
	key := this.GetKey()
	if this.GetValue(key) == nil {
		data, _ := json.Marshal(this.ClusterInfo)
		this.m_Client.Put(context.Background(), key, string(data), clientv3.WithLease(this.m_LeaseId))
	}
}

func (this *Service) Lease() {
	if this.m_LeaseId > 0 {

		//监听租约
		go func() {
			for {
				ctx, _ := context.WithCancel(context.Background())
				leaseRespChan, err := this.m_Lease.KeepAlive(ctx, this.m_LeaseId)
				if err != nil {
					log.Fatal("Error: Service KeepAlive error :", err.Error())
				}

				select {
				case resp := <-leaseRespChan:
					if resp == nil {
						//log.Println("租约已经到期关闭")
						goto LEASE_OVER
					} else {
						//log.Println("续租成功")
						goto END
					}
				}
			LEASE_OVER:
				//log.Println("lease 监听结束")
				this.Grant()
				this.KeepAlive()
				this.Put()
				//break
			END:
				time.Sleep(1000 * time.Millisecond)
			}
		}()
	}
}

//续租
func (this *Service) KeepAlive() {
	if this.m_LeaseId > 0 {
		ctx, _ := context.WithCancel(context.Background())
		_, err := this.m_Lease.KeepAlive(ctx, this.m_LeaseId)
		if err != nil {
			log.Fatal("Error: Service KeepAlive error :", err.Error())
		}
	}
}

func (this *Service) Revoke() {
	//撤销租约
	_, err := this.m_Lease.Revoke(context.TODO(), this.m_LeaseId)
	if err != nil {
		log.Fatalf("撤销租约失败:%s\n", err.Error())
	}
	log.Println("撤销租约成功")
}

//检查是否被注册过
func (this *Service) CheckExist(info *common.ClusterInfo, endpoints []string) bool {
	cfg := clientv3.Config{
		Endpoints: endpoints,
	}

	etcdClient, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
		return false
	}
	this.m_Client = etcdClient
	this.ClusterInfo = info
	//lease := clientv3.NewLease(etcdClient)
	key := this.GetKey()
	if this.GetValue(key) != nil {
		log.Println("Error: Service.Start error， Exist key :", key)
		return false
	}
	return true
}

//注册服务器
func (this *Service) Init(info *common.ClusterInfo, endpoints []string) bool {
	cfg := clientv3.Config{
		Endpoints: endpoints,
	}

	etcdClient, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
		return false
	}
	lease := clientv3.NewLease(etcdClient)
	this.m_Client = etcdClient
	this.m_Lease = lease
	this.ClusterInfo = info
	this.isRun = false
	return this.Start()
}

func (this *Service) Start() bool {
	key := this.GetKey()
	if this.GetValue(key) != nil {
		log.Println("Error: Service.Start error， Exist key :", key)
		return false
	}
	this.isRun = true
	this.Run()
	return true
}

func (this *Service) GetKey() string {
	key := ETCD_DIR + this.String() + "/" + this.IpString()
	return key
}

func (this *Service) GetRootKey() string {
	key := ETCD_DIR
	return key
}

func (this *Service) GetValue(key string) []*mvccpb.KeyValue {
	getResp, err := this.m_Client.Get(context.TODO(), key)
	if err != nil {
		log.Fatalf("get 失败：%s", err.Error())
	}
	log.Printf("%v", getResp.Kvs)
	return getResp.Kvs
}
