package main

import (
	"context"
	"gonet/actor"
	"gonet/common"
	"gonet/rpc"
	"gonet/server/message"
	"strconv"
)

type (
	CmdProcess struct {
		actor.Actor
	}

	ICmdProcess interface {
		actor.IActor
	}
)

func (this *CmdProcess) Init() {
	this.Actor.Init()
	this.RegisterCall("msg", func(ctx context.Context, args string) {
		packet1 := &message.C_W_ChatMessage{PacketHead: message.BuildPacketHead(PACKET.AccountId, rpc.SERVICE_GATESERVER),
			Sender:      PACKET.PlayerId,
			Recver:      0,
			MessageType: int32(message.CHAT_MSG_TYPE_WORLD),
			Message:     (args),
		}
		SendPacket(packet1)
	})

	this.RegisterCall("move", func(ctx context.Context, yaw string) {
		ya, _ := strconv.ParseFloat(yaw, 32)
		PACKET.Move(float32(ya), 100.0)
	})

	this.Actor.Start()
}

var (
	g_Cmd *CmdProcess
)

func InitCmd() {
	g_Cmd = &CmdProcess{}
	g_Cmd.Init()
	common.StartConsole(g_Cmd)
}
