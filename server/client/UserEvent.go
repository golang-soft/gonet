package main

import (
	"fmt"
)

type (
	UserEvent struct {
		//BaseEvent
	}
	IUserEvent interface {
		IBaseEvent
	}
)

func NewUserEvent() *UserEvent {
	return &UserEvent{}
}

func (this *UserEvent) DoEvent(event *IBaseEvent) {
	fmt.Printf("UserEvent doEvent.......")
	//event.DoEvent(event)
}

func (this *UserEvent) Name() string {
	return "UserEvent"
}

func (this *UserEvent) HandleEvent(event *IBaseEvent) {

}

func (this *UserEvent) EventID() int {
	return 2
}

func (this *UserEvent) SendEvent(event *IBaseEvent, process *EventProcess) {

}
