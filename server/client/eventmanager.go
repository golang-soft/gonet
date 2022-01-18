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
	this.Add("TestEvent", NewTestEvent())
	this.AddById(0, NewTestEvent())

	this.Add("UserEvent", NewUserEvent())
	this.AddById(1, NewUserEvent())

	this.Add("AttackEvent", NewAttackEvent())
	this.AddById(2, NewAttackEvent())

	this.Add("LoginEvent", NewLoginEvent())
	this.AddById(3, NewLoginEvent())
	this.Add("LoginAccountEvent", NewLoginAccountEvent())
	this.AddById(4, NewLoginAccountEvent())
	this.Add("LoginGateEvent", NewLoginGateEvent())
	this.AddById(5, NewLoginGateEvent())

}

func (this *EventManager) Add(name string, event IBaseEvent) {
	this.events[name] = &event
}

func (this *EventManager) AddById(eventid int, event IBaseEvent) {
	event.SetEventId(eventid)
	this.events2[(event).EventID()] = &event
}

func (this *EventManager) GetEvent(name string) *IBaseEvent {
	return this.events[name]
}

func (this *EventManager) GetEventById(id int) *IBaseEvent {
	return this.events2[id]
}
