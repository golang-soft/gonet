package main

import (
	"github.com/golang/protobuf/proto"
	"gonet/actor"
	"gonet/grpc"
	"gonet/server/common"
	"gonet/server/rpc"
)

type SuperEventProcess struct {
	actor.Actor

	Robot *Robot
}

func (this *SuperEventProcess) SendPacket(packet proto.Message) {
	buff := common.Encode(packet)
	this.Robot.Send(rpc.RpcHead{}, buff)
}

func (this *SuperEventProcess) PacketFunc(packet1 rpc.Packet) bool {
	packetId, data := common.Decode(packet1.Buff)
	packet := common.GetPakcet(packetId)
	if packet == nil {
		return true
	}
	err := common.UnmarshalText(packet, data)
	if err == nil {
		this.Send(rpc.RpcHead{}, grpc.Marshal(rpc.RpcHead{}, common.GetMessageName(packet), packet))
		return true
	}

	return true
}
