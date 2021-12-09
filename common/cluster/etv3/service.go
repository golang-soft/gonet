package etv3

import (
	"encoding/json"
	"go.etcd.io/etcd/clientv3"
	"gonet/common"
	"log"
	"time"

	"golang.org/x/net/context"
)

const (
	ETCD_DIR = "server/"
)

//注册服务器
type Service struct {
	*common.ClusterInfo
	m_Client  *clientv3.Client
	m_Lease   clientv3.Lease
	m_LeaseId clientv3.LeaseID
}

func (this *Service) Run() {
	//this.Grant()
	this.Put()
	for {
		this.KeepAlive()
		time.Sleep(time.Second * 3)
	}
}

//设置租约时间
func (this *Service) Grant() {
	leaseResp, _ := this.m_Lease.Grant(context.Background(), 10)
	this.m_LeaseId = leaseResp.ID
}

//通过租约put
func (this *Service) Put() {
	key := ETCD_DIR + this.String() + "/" + this.IpString()
	data, _ := json.Marshal(this.ClusterInfo)
	this.m_Client.Put(context.Background(), key, string(data), clientv3.WithLease(this.m_LeaseId))
}

//续租
func (this *Service) KeepAlive() {
	if this.m_LeaseId > 0 {
		ctx, _ := context.WithCancel(context.Background())
		leaseRespChan, err := this.m_Lease.KeepAlive(ctx, this.m_LeaseId)
		if err != nil {
			log.Fatal("Error: Service KeepAlive error :", err.Error())
		}

		//监听租约
		go func() {
			for {
				select {
				case resp := <-leaseRespChan:
					if resp == nil {
						log.Println("租约已经到期关闭")
						goto LEASE_OVER
					} else {
						//log.Println("续租成功")
						goto END
					}
				}
			LEASE_OVER:
				log.Println("lease 监听结束")
				this.Grant()
				this.Put()
				break
			END:
				time.Sleep(500 * time.Millisecond)
			}
		}()

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

//注册服务器
func (this *Service) Init(info *common.ClusterInfo, endpoints []string) {
	cfg := clientv3.Config{
		Endpoints: endpoints,
	}

	etcdClient, err := clientv3.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}
	lease := clientv3.NewLease(etcdClient)
	this.m_Client = etcdClient
	this.m_Lease = lease
	this.ClusterInfo = info
	this.Start()
}

func (this *Service) Start() {
	this.Grant()
	go this.Run()
}
