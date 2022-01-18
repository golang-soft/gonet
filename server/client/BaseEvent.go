package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

type (
	BaseEvent struct {
		EventProcess
		eventid int
	}

	IBaseEvent interface {
		IEvent
		DoEvent(process *EventProcess)
		SendEvent(event *IBaseEvent, process *EventProcess)
		Name() string
		HandleEvent(event *IBaseEvent, message proto.Message)
		EventID() int
		SetEventId(eventid int)
	}
)

func (this *BaseEvent) DoEvent(process *EventProcess) {
	fmt.Printf("baseEvent doEvent.......")
}

func (this *BaseEvent) Name() string {
	return "BaseEvent"
}

func (this *BaseEvent) HandleEvent(event *IBaseEvent, message proto.Message) {
	fmt.Printf("baseEvent HandleEvent.......")

}

func (this *BaseEvent) SetEventId(eventid int) {
	this.eventid = eventid
}

func (this *BaseEvent) SendEvent(event *IBaseEvent, process *EventProcess) {

}
