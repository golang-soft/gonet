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
	AddAtk          int
	Base_rate       int
	Base_step       int
	Base_skill_Time int
	Skill_1101      int
	Skill_1102      int
	Skill_1103      int
	Skill_1104      int
	Skill_1201      int
	Skill_1202      int
	Skill_1203      int
	Skill_1204      int
	Skill_1301      int
	Skill_1302      int
	Skill_1303      int
	Skill_1304      int32
	Skill_1401      int
	Skill_1402      int32

	Skill_1403 int
	Skill_1404 int
	Skill_1501 int
	Skill_1502 int
	Skill_1503 int
	Skill_1504 int

	buff_0  float64
	Buff_1  float64
	Buff_2  float64
	Buff_3  float64
	Buff_4  float64
	Buff_5  float64
	Buff_6  float64
	Buff_7  float64
	Buff_8  float64
	Buff_9  float64
	Buff_10 float64
}

var Skill = &DSkill{
	AddAtk:          1,
	Base_rate:       100,
	Base_step:       100,
	Base_skill_Time: 100,
	Skill_1101:      1101,
	Skill_1102:      1102,
	Skill_1103:      1103,
	Skill_1104:      1104,
	Skill_1201:      1201,
	Skill_1202:      1202,
	Skill_1203:      1203,
	Skill_1204:      1204,
	Skill_1301:      1301,
	Skill_1302:      1302,
	Skill_1303:      1303,
	Skill_1304:      1304,
	Skill_1401:      1401,
	Skill_1402:      1402,
	Skill_1403:      1403,
	Skill_1404:      1404,
	Skill_1501:      1501,
	Skill_1502:      1502,
	Skill_1503:      1503,
	Skill_1504:      1504,

	buff_0:  0,
	Buff_1:  1,
	Buff_2:  2,
	Buff_3:  3,
	Buff_4:  4,
	Buff_5:  5,
	Buff_6:  6,
	Buff_7:  7,
	Buff_8:  8,
	Buff_9:  9,
	Buff_10: 10,
}
