package main

import (
	"github.com/golang/protobuf/proto"
	"log"
)

type (
	PingEvent struct {
		BaseEvent
	}
	IPingEvent interface {
		IBaseEvent
	}
)

func NewPingEvent() *PingEvent {
	return &PingEvent{}
}

func (this *PingEvent) DoEvent(event *EventProcess) {
	log.Printf("PingEvent doEvent.......")
	this.Ping(event)
}

func (this *PingEvent) Name() string {
	return "PingEvent"
}

func (this *PingEvent) HandleEvent(event *IBaseEvent, message proto.Message) {

}

func (this *PingEvent) EventID() int {
	return this.eventid
}

func (this *PingEvent) SendEvent(event *IBaseEvent, process *EventProcess) {
	this.Ping(process)
}

func (this *PingEvent) Ping(process *EventProcess) {
	//packet := &cmessage.HeartPacket{PacketHead: common.BuildPacketHead(cmessage.MessageID_MSG_C_HeartPacket, rpc.SERVICE_GATESERVER)}
	//this.SendPacket(process, packet)
}
