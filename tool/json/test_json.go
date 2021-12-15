package main

import (
	io "io/ioutil"

	json "encoding/json"

	"fmt"
)

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {

	return &JsonStruct{}

}

func (self *JsonStruct) Load(filename string, v interface{}) {

	data, err := io.ReadFile(filename)

	if err != nil {

		return

	}

	datajson := []byte(data)

	err = json.Unmarshal(datajson, v)

	if err != nil {

		return

	}

}

type ValueTestAtmp struct {
	StringValue string

	NumericalValue int

	BoolValue bool
}

type testdata struct {
	ValueTestA ValueTestAtmp
}

func main() {

	JsonParse := NewJsonStruct()

	v := testdata{}

	JsonParse.Load("D:\\workspace-go\\gonet\\tool\\json\\testdata.json", &v)

	fmt.Println(v)

	fmt.Println(v.ValueTestA.StringValue)

}
