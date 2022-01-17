package main

import (
	"fmt"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
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

func (this *AttackEvent) SendAttack() {
	packet := &cmessage.AttackReq{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_AttackReq, rpc.SERVICE_GATESERVER), Round: 1}
	this.SendPacket(packet)
	//m_Log.Debugf("玩家 %d 攻击", this.PlayerId)
}

func (this *AttackEvent) DoEvent(event *IBaseEvent) {
	fmt.Printf("AttackEvent doEvent.......")
	//(*event).DoEvent(event)

}

func (this *AttackEvent) SendEvent(event *IBaseEvent, process *EventProcess) {

}

func (this *AttackEvent) Name() string {
	return "AttackEvent"
}

func (this *AttackEvent) HandleEvent(event *IBaseEvent) {

}

func (this *AttackEvent) EventID() int {
	return 1
}
