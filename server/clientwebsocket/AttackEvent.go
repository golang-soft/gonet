package main

import (
	"github.com/golang/protobuf/proto"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
	"log"
)

type (
	AttackEvent struct {
		BaseEvent
	}
	IAttackEvent interface {
		IBaseEvent
	}
)

func NewAttackEvent() *AttackEvent {
	return &AttackEvent{}
}

func (this *AttackEvent) SendAttack(process *EventProcess) {
	packet := &cmessage.AttackReq{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_AttackReq, rpc.SERVICE_GATESERVER), Round: 1}
	this.SendPacket(process, packet)
	m_Log.Debugf("玩家 %d 攻击", process.PlayerId)
}

func (this *AttackEvent) DoEvent(process *EventProcess) {
	log.Printf("AttackEvent doEvent.......")
	//(*event).DoEvent(event)
	this.SendAttack(process)
}

func (this *AttackEvent) SendEvent(event *IBaseEvent, process *EventProcess) {
	this.SendAttack(process)
}

func (this *AttackEvent) Name() string {
	return "AttackEvent"
}

func (this *AttackEvent) HandleEvent(event *IBaseEvent, message proto.Message) {

}

func (this *AttackEvent) EventID() int {
	return this.eventid
}
