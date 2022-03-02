package main

//
//import (
//	"fmt"
//	"reflect"
//)
//
//func test(event *IBaseEvent) {
//	(*event).Name()
//}
//
//func main() {
//	var event *UserEvent = new(UserEvent)
//	event.DoEvent(nil)
//	var eventmanager = NewEventManager()
//	eventmanager.Init()
//
//	var e *IBaseEvent = eventmanager.GetEvent("AttackEvent")
//	fmt.Printf("%v %s", e, reflect.TypeOf(e))
//	test(e)
//	(*e).DoEvent(e)
//
//	e = eventmanager.GetEvent("UserEvent")
//	fmt.Printf("%v %s", e, reflect.TypeOf(e))
//	test(e)
//	(*e).DoEvent(e)
//
//}
