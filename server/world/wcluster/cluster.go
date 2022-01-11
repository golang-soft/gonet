package wcluster

import (
	"github.com/golang/protobuf/proto"
	"gonet/common/cluster"
	"gonet/server/common"
	"gonet/server/rpc"
)

var m_pCluster *cluster.Cluster

func SetCluster(pCluster *cluster.Cluster) {
	m_pCluster = pCluster
}
func GetCluster() *cluster.Cluster {
	return m_pCluster
}

//发送account
func SendToAccount(funcName string, params ...interface{}) {
	head := rpc.RpcHead{DestServerType: rpc.SERVICE_ACCOUNTSERVER, SendType: rpc.SEND_BALANCE, SrcClusterId: GetCluster().Id()}
	GetCluster().SendMsg(head, funcName, params...)
}

//发送给客户端
func SendToClient(clusterId uint32, packet proto.Message) {
	pakcetHead := packet.(common.Packet).GetPacketHead()
	if pakcetHead != nil {
		GetCluster().SendMsg(rpc.RpcHead{DestServerType: rpc.SERVICE_GATESERVER, ClusterId: clusterId, Id: int64(pakcetHead.Id)}, "", proto.MessageName(packet), packet)
	}
}

//--------------发送给地图----------------------//
func SendToZone(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_ZONESERVER, SrcClusterId: GetCluster().Id()}
	GetCluster().SendMsg(head, funcName, params...)
}

//--------------发送给中央服----------------------//
func SendToCenter(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_CENTERSERVER, SrcClusterId: GetCluster().Id(), SendType: rpc.SEND_BOARD_CAST}
	GetCluster().SendMsg(head, funcName, params...)
}

func SendToDB(funcName string, params ...interface{}) {
	head := rpc.RpcHead{DestServerType: rpc.SERVICE_WORLDDBSERVER, SrcClusterId: GetCluster().Id(), SendType: rpc.SEND_BALANCE}
	GetCluster().SendMsg(head, funcName, params...)
}
