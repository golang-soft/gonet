package main

import (
	"fmt"
	"gonet/base"
	"gonet/base/config"
	"gonet/server/table"
)

func main() {

	JsonParse := base.NewJsonTableData()
	vSkillTableData := table.SkillTableData{}
	config.InitEnv("local")
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("SkillTable.json"), &vSkillTableData)
	fmt.Println(vSkillTableData["1101"].Name_String)
	vBoxTableData := table.BoxTableData{}
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("BoxTable.json"), &vBoxTableData)
	fmt.Println(vBoxTableData["1001"].Item[0][1])

}
