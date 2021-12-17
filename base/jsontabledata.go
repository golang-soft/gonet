package base

import (
	"encoding/json"
	io "io/ioutil"
	"log"
)

type JsonTableData struct {
}

func NewJsonTableData() *JsonTableData {
	return &JsonTableData{}
}

func (self *JsonTableData) LoadJsonTableData(filename string, v interface{}) {
	data, err := io.ReadFile(filename)
	if err != nil {
		return
	}

	datajson := []byte(data)
	err = json.Unmarshal(datajson, v)
	if err != nil {
		log.Printf("解析文件错误, %v", err)
		return
	}
}
