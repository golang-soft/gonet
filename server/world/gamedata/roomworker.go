package gamedata

import (
	"fmt"
	"gonet/server/common/data"
	"reflect"
)

func MakeService(rep interface{}) *service {
	ser := service{}
	ser.typ = reflect.TypeOf(rep)
	ser.rcvr = reflect.ValueOf(rep)
	// name返回其包中的类型名称，举个例子，这里会返回Person，tool
	name := reflect.Indirect(ser.rcvr).Type().Name()
	fmt.Println(name)
	ser.servers = map[string]reflect.Method{}
	fmt.Println(ser.typ.NumMethod(), ser.typ.Name())
	for i := 0; i < ser.typ.NumMethod(); i++ {
		method := ser.typ.Method(i)
		//mtype := method.Type
		//mtype := method.Type	// reflect.method
		mname := method.Name // string
		fmt.Println("mname : ", mname)
		ser.servers[mname] = method
	}
	return &ser
}

func HandleRoomTask(message data.RoomData) {
	var methodName = message.FuncName
	// 得到这个对象的全部方法，string对应reflect.method
	methods := MakeService(&SRoomCtrl{})
	// 利用得到的methods来调用其值
	if method, ok := methods.servers[methodName]; ok {
		// 得到第一个此method第1参数的Type，第零个当然就是结构体本身了
		size := method.Type.Size()
		fmt.Sprintf("%d", size)
		function := method.Func
		function.Call([]reflect.Value{methods.rcvr, reflect.ValueOf(message)})
	}
}
