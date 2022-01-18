package main

import (
	"github.com/golang/protobuf/proto"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/rpc"
	"log"
)

type (
	TestEvent struct {
		BaseEvent
	}
	ITestEvent interface {
		IBaseEvent
	}
)

func NewTestEvent() *TestEvent {
	return &TestEvent{}
}

func (this *TestEvent) SendTest(event *EventProcess) {
	aa := []int32{}
	for i := 0; i < 10; i++ {
		aa = append(aa, int32(1))
	}

	packet1 := &cmessage.W_C_Test{PacketHead: common.BuildPacketHead(0, rpc.SERVICE_GATESERVER),
		Recv: aa}
	this.SendPacket(event, packet1)
}

func (this *TestEvent) DoEvent(event *EventProcess) {
	log.Printf("TestEvent doEvent.......")
	this.SendTest(event)
}

func (this *TestEvent) Name() string {
	return "TestEvent"
}

func (this *TestEvent) HandleEvent(event *IBaseEvent, message proto.Message) {

}

func (this *TestEvent) EventID() int {
	return this.eventid
}

func (this *TestEvent) SendEvent(event *IBaseEvent, process *EventProcess) {
	log.Printf("TestEvent doEvent.......")
	this.SendTest(process)
}
