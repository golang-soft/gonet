package main

import (
	"fmt"
)

type (
	BaseEvent struct {
		EventProcess
	}

	IBaseEvent interface {
		IEvent
		DoEvent(event *IBaseEvent)
		SendEvent(event *IBaseEvent, process *EventProcess)
		Name() string
		HandleEvent(event *IBaseEvent)
		EventID() int
	}
)

func (this *BaseEvent) DoEvent(event *BaseEvent) {
	fmt.Printf("baseEvent doEvent.......")
}

func (this *BaseEvent) Name() string {
	return "BaseEvent"
}

func (this *BaseEvent) HandleEvent(event *IBaseEvent) {
	fmt.Printf("baseEvent HandleEvent.......")

}

func (this *BaseEvent) SendEvent(event *IBaseEvent, process *EventProcess) {

}
