package main

type EventManager struct {
	events  map[string]*IBaseEvent
	events2 map[int]*IBaseEvent
}

func NewEventManager() *EventManager {
	return &EventManager{
		events:  make(map[string]*IBaseEvent),
		events2: make(map[int]*IBaseEvent),
	}
}

func (this *EventManager) Init() {

	this.Add("UserEvent", NewUserEvent())
	this.AddById(NewUserEvent())

	this.Add("AttackEvent", NewAttackEvent())
	this.AddById(NewAttackEvent())

}

func (this *EventManager) Add(name string, event IBaseEvent) {
	this.events[name] = &event
}

func (this *EventManager) AddById(event IBaseEvent) {
	this.events2[(event).EventID()] = &event
}

func (this *EventManager) GetEvent(name string) *IBaseEvent {
	return this.events[name]
}

func (this *EventManager) GetEventById(id int) *IBaseEvent {
	return this.events2[id]
}
