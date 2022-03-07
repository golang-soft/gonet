package utils

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	data "gonet/server/common/data"
	"reflect"
)

func ConvertRoundData(data map[string]interface{}, game data.RoundGameData) {
	vt := reflect.TypeOf(game)
	vv := reflect.ValueOf(game)

	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		data[f.Name] = vv.FieldByName(f.Name).String()
	}
}

func ConvertGameAttrData(data map[string]interface{}, player *data.UserGameAttrData) {
	vt := reflect.TypeOf(*player)
	vv := reflect.ValueOf(*player)

	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		data[f.Name] = vv.FieldByName(f.Name).String()
	}
}

//func MapToStructDemo() {
//mapInstance := make(map[string]interface{})
//mapInstance["Name"] = "jqw"
//mapInstance["Age"] = 18
//
//var people People
//err := mapstructure.Decode(mapInstance, &people)
//if err != nil {
//	fmt.Println(err)
//}
//fmt.Println(people)
//}

func ConvertMapToStruct(mapInstance map[string]interface{}, obj *interface{}) {
	err := mapstructure.Decode(mapInstance, obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
}

func ConvertStructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func TestStructToMap() {
	//student := Student{10, "jqw", 18}
	//data := StructToMapDemo(student)
	//fmt.Println(data)
}
