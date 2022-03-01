package netgate

import (
	"bytes"
	"encoding/gob"
	"github.com/golang/protobuf/proto"
	"gonet/base"
	"gonet/grpc"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
	"reflect"
)

var (
	A_C_RegisterResponse = proto.MessageName(&cmessage.A_C_RegisterResponse{})
	A_C_LoginResponse    = proto.MessageName(&cmessage.A_C_LoginResponse{})
	W_C_Test             = proto.MessageName(&cmessage.W_C_Test{})
)

func SendToClient(socketId uint32, packet proto.Message) {
	if SERVER.CheckIsWebsocket() {
		SERVER.GetWebSocketServer().Send(rpc.RpcHead{SocketId: socketId}, common.Encode(packet))
	} else {
		SERVER.GetServer().Send(rpc.RpcHead{SocketId: socketId}, common.Encode(packet))
	}
}

func DispatchPacket(packet rpc.Packet) bool {
	defer func() {
		if err := recover(); err != nil {
			base.TraceCode(err)
		}
	}()

	rpcPacket, head := grpc.Unmarshal(packet.Buff)
	switch head.DestServerType {
	case rpc.SERVICE_GATESERVER:
		messageName := ""
		buf := bytes.NewBuffer(rpcPacket.RpcBody)
		dec := gob.NewDecoder(buf)
		dec.Decode(&messageName)
		packet := reflect.New(proto.MessageType(messageName).Elem()).Interface().(proto.Message)
		dec.Decode(packet)
		buff := common.Encode(packet)
		if messageName == A_C_RegisterResponse || messageName == A_C_LoginResponse {
			if SERVER.CheckIsWebsocket() {
				SERVER.GetWebSocketServer().Send(rpc.RpcHead{SocketId: head.SocketId}, buff)
			} else {
				SERVER.GetServer().Send(rpc.RpcHead{SocketId: head.SocketId}, buff)
			}
		} else if messageName == W_C_Test {
			SERVER.GetEventProcess().SendMsg(rpc.RpcHead{SocketId: head.SocketId}, "W_C_Test")
		} else {
			socketId := head.SocketId //SERVER.GetPlayerMgr().GetSocket(head.Id)
			if SERVER.CheckIsWebsocket() {
				SERVER.GetWebSocketServer().Send(rpc.RpcHead{SocketId: socketId}, buff)
			} else {
				SERVER.GetServer().Send(rpc.RpcHead{SocketId: socketId}, buff)
			}

		}
	default:
		SERVER.GetCluster().Send(head, packet.Buff)
	}

	return true
}
