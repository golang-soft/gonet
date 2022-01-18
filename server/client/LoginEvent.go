package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
)

type (
	LoginEvent struct {
		BaseEvent
	}
	ILoginEvent interface {
		IBaseEvent
	}
)

func NewLoginEvent() *LoginEvent {
	return &LoginEvent{}
}

func (this *LoginEvent) DoEvent(process *EventProcess) {
	fmt.Printf("LoginEvent doEvent.......")
	this.LoginGame(process)
}

func (this *LoginEvent) Name() string {
	return "LoginEvent"
}

func (this *LoginEvent) HandleEvent(event *IBaseEvent, message proto.Message) {

}

func (this *LoginEvent) EventID() int {
	return this.eventid
}

func (this *LoginEvent) SendEvent(event *IBaseEvent, process *EventProcess) {

}

func (this *LoginEvent) LoginGame(process *EventProcess) {
	packet1 := &cmessage.C_W_Game_LoginRequset{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_C_W_Game_LoginRequset, rpc.SERVICE_GATESERVER),
		PlayerId: process.PlayerId}
	this.SendPacket(process, packet1)
}
