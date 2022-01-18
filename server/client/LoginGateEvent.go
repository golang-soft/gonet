package main

import (
	"github.com/golang/protobuf/proto"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
	"log"
)

type (
	LoginGateEvent struct {
		BaseEvent
	}
	ILoginGateEvent interface {
		IBaseEvent
	}
)

func NewLoginGateEvent() *LoginGateEvent {
	return &LoginGateEvent{}
}

func (this *LoginGateEvent) DoEvent(event *EventProcess) {
	log.Printf("LoginGateEvent doEvent.......")
	this.LoginGate(event)
}

func (this *LoginGateEvent) Name() string {
	return "LoginGateEvent"
}

func (this *LoginGateEvent) HandleEvent(event *IBaseEvent, message proto.Message) {

}

func (this *LoginGateEvent) EventID() int {
	return this.eventid
}

func (this *LoginGateEvent) SendEvent(event *IBaseEvent, process *EventProcess) {
	this.LoginGate(process)
}

func (this *LoginGateEvent) LoginGate(process *EventProcess) {
	packet := &cmessage.C_G_LoginResquest{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_C_G_LoginResquest, rpc.SERVICE_GATESERVER),
		Key: process.m_Dh.PubKey()}
	this.SendPacket(process, packet)
}
