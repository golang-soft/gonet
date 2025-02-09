package login

import (
	"github.com/golang/protobuf/proto"
	"gonet/rpc"
	"gonet/server/message"
	"net/http"
)

//发送account
func SendToAccount(funcName string, params ...interface{}) {
	head := rpc.RpcHead{DestServerType: rpc.SERVICE_ACCOUNTSERVER, SendType: rpc.SEND_BALANCE, SrcClusterId: SERVER.GetCluster().Id()}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

//发送给客户端
func SendToClient(clusterId uint32, packet proto.Message) {
	pakcetHead := packet.(message.Packet).GetPacketHead()
	if pakcetHead != nil {
		SERVER.GetCluster().SendMsg(rpc.RpcHead{DestServerType: rpc.SERVICE_GATESERVER, ClusterId: clusterId, Id: pakcetHead.Id, SendType: rpc.SEND_BOARD_CAST}, "", proto.MessageName(packet), packet)
	}
}

//--------------发送给地图----------------------//
func SendToZone(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_ZONESERVER, SrcClusterId: SERVER.GetCluster().Id()}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

//--------------发送给中央服----------------------//
func SendToCenter(Id int64, ClusterId uint32, funcName string, params ...interface{}) {
	head := rpc.RpcHead{Id: Id, ClusterId: ClusterId, DestServerType: rpc.SERVICE_CENTERSERVER, SrcClusterId: SERVER.GetCluster().Id(), SendType: rpc.SEND_BOARD_CAST}
	SERVER.GetCluster().SendMsg(head, funcName, params...)
}

//--------------发送给中央服----------------------//
func SendToCenter2(clusterId uint32, funcName string, packet proto.Message) {
	//pakcetHead := packet.(message.Packet).GetPacketHead()
	//if pakcetHead != nil {
	SERVER.GetCluster().SendMsg(rpc.RpcHead{DestServerType: rpc.SERVICE_CENTERSERVER, ClusterId: clusterId, SendType: rpc.SEND_BOARD_CAST}, funcName, packet)
	//}
}

//--------------发送给游戏服----------------------//
func SendToWorld(clusterId uint32, funcName string, packet proto.Message) {
	//pakcetHead := packet.(message.Packet).GetPacketHead()
	//if pakcetHead != nil {
	SERVER.GetCluster().SendMsg(rpc.RpcHead{DestServerType: rpc.SERVICE_WORLDSERVER, ClusterId: clusterId, SendType: rpc.SEND_BOARD_CAST}, funcName, packet)
	//}
}

//--------------发送给grpc服----------------------//
func SendToGrpcServer(clusterId uint32, funcName string, packet proto.Message) {
	//pakcetHead := packet.(message.Packet).GetPacketHead()
	//if pakcetHead != nil {
	SERVER.GetCluster().SendMsg(rpc.RpcHead{DestServerType: rpc.SERVICE_GRPCSERVER, ClusterId: clusterId, SendType: rpc.SEND_BOARD_CAST}, funcName, packet)
	//}
}

func Test(w http.ResponseWriter, r *http.Request) {
	//SendToClient(1, &message.W_C_Test{PacketHead: message.BuildPacketHead(1, 0),})
	//SendToCenter(1, 0, "LoginCenter","")
	SendToCenter2(1, "PlayerData", &message.PlayerData{PlayerID: 1111, PlayerName: "顶顶顶顶"})
}

func TestWorld(w http.ResponseWriter, r *http.Request) {
	SendToWorld(1, "W_C_Test", &message.W_C_Test{PacketHead: message.BuildPacketHead(1, 0), PlayerId: 111222})
}

func TestGrpc(w http.ResponseWriter, r *http.Request) {
	SendToGrpcServer(1, "W_C_Test", &message.W_C_Test{PacketHead: message.BuildPacketHead(1, 0), PlayerId: 111222})
}
