package table

import (
	"gonet/base"
	"gonet/base/config"
	"gonet/server/table"
)

var (
	BOX_CONFIG            = &table.BoxTableData{}
	HELP_CONFIG           = &table.HelpTableData{}
	RAW_HERO_ATTR         = &table.HeroAttrTableData{}
	ITEM_CONFIG           = &table.ItemTableData{}
	MAP_CONFIG            = &table.MapData{}
	RAND_NAME_CONFIG      = &table.RandNameTableData{}
	RANKREWARD_TABLE_DATA = &table.RankRewardTableData{}
	SKILL_BASIC_INFO      = &table.SkillTableData{}
	STRENGTHEN_TABLE_DATA = &table.StrengthenTableData{}
	STR_TABLE_CONFIG      = &table.StrTableData{}
	WORLD_CONFIG          = &table.WorldParamCfgTableData{}
)

func Init() {
	config.InitEnv("local")
	JsonParse := base.NewJsonTableData()

	JsonParse.LoadJsonTableData(config.GetConfigTablePath("HeroAttrTable.json"), RAW_HERO_ATTR)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("SkillTable.json"), SKILL_BASIC_INFO)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("Map.json"), MAP_CONFIG)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("RandNameTable.json"), RAND_NAME_CONFIG)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("ItemTable.json"), ITEM_CONFIG)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("WorldParamTable.json"), WORLD_CONFIG)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("RankRewardTable.json"), RANKREWARD_TABLE_DATA)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("StrengthenTable.json"), STRENGTHEN_TABLE_DATA)

	JsonParse.LoadJsonTableData(config.GetConfigTablePath("SuiTable.json"), STRENGTHEN_TABLE_DATA)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("BattlePassRankRewardTable.json"), STRENGTHEN_TABLE_DATA)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("game_server.json"), STRENGTHEN_TABLE_DATA)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("room_server.json"), STRENGTHEN_TABLE_DATA)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("create_new_player.json"), STRENGTHEN_TABLE_DATA)

	JsonParse.LoadJsonTableData(config.GetConfigTablePath("BoxTable.json"), BOX_CONFIG)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("HelpTable.json"), HELP_CONFIG)
	JsonParse.LoadJsonTableData(config.GetConfigTablePath("StrTable.json"), STR_TABLE_CONFIG)

}
