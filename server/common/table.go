package common

const (
	Tables_user        = "account"
	tables_mail        = "mail"
	Tables_item        = "item"
	Tables_equip       = "equip"
	tables_task        = "task"
	Tables_hero        = "heroes"
	tables_nonce       = "nonce"
	tables_token       = "token"
	tables_leaderboard = "leaderboard"
	tables_contract    = "contract"

	Redis_Prefix_Item = "items"
)

const (
	Item_consume   = 1
	Item_equip     = 2
	Item_box       = 1001 //宝箱
	Item_allAttr   = 1004 //全属性
	Item_box_piece = 3001 //宝箱碎片
)

type DTIME struct {
	Second int
	Minute int
}

var TIME = &DTIME{
	Second: 1000,
	Minute: 60,
}

type DBattle struct {
	Time_3  int
	Time_5  int
	Time_10 int
	Time_15 int
}

var Battle = &DBattle{
	Time_3:  3,
	Time_5:  5,
	Time_10: 10,
	Time_15: 15,
}

type DSkill struct {
	addAtk          int
	base_rate       int
	base_step       int
	base_skill_Time int
	skill_1101      int
	skill_1102      int
	skill_1103      int
	skill_1104      int
	skill_1201      int
	skill_1202      int
	skill_1203      int
	skill_1204      int
	skill_1301      int
	skill_1302      int
	skill_1303      int
	skill_1304      int
	skill_1401      int
	Skill_1402      int
	skill_1403      int
	skill_1404      int
	skill_1501      int
	skill_1502      int
	skill_1503      int
	skill_1504      int

	buff_0  int
	buff_1  int
	buff_2  int
	buff_3  int
	buff_4  int
	buff_5  int
	buff_6  int
	buff_7  int
	buff_8  int
	buff_9  int
	buff_10 int
}

var Skill = &DSkill{
	addAtk:          1,
	base_rate:       100,
	base_step:       100,
	base_skill_Time: 100,
	skill_1101:      1101,
	skill_1102:      1102,
	skill_1103:      1103,
	skill_1104:      1104,
	skill_1201:      1201,
	skill_1202:      1202,
	skill_1203:      1203,
	skill_1204:      1204,
	skill_1301:      1301,
	skill_1302:      1302,
	skill_1303:      1303,
	skill_1304:      1304,
	skill_1401:      1401,
	Skill_1402:      1402,
	skill_1403:      1403,
	skill_1404:      1404,
	skill_1501:      1501,
	skill_1502:      1502,
	skill_1503:      1503,
	skill_1504:      1504,

	buff_0:  0,
	buff_1:  1,
	buff_2:  2,
	buff_3:  3,
	buff_4:  4,
	buff_5:  5,
	buff_6:  6,
	buff_7:  7,
	buff_8:  8,
	buff_9:  9,
	buff_10: 10,
}
