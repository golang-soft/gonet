package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

type (
	UserEvent struct {
		BaseEvent
	}
	IUserEvent interface {
		IBaseEvent
	}
)

func NewUserEvent() *UserEvent {
	return &UserEvent{}
}

func (this *UserEvent) DoEvent(event *EventProcess) {
	fmt.Printf("UserEvent doEvent.......")
	//event.DoEvent(event)
}

func (this *UserEvent) Name() string {
	return "UserEvent"
}

func (this *UserEvent) HandleEvent(event *IBaseEvent, message proto.Message) {

}

func (this *UserEvent) EventID() int {
	return this.eventid
}

func (this *UserEvent) SendEvent(event *IBaseEvent, process *EventProcess) {

}
