package helper

import (
	"fmt"
	table2 "gonet/server/table"
	"gonet/server/world/table"
	"strconv"
)

type USER_KIND int

const (
	HUNTER USER_KIND = 1
	WIZARD
	KNIGHT
	GUNNER
	SOLDIER
)

type BasicInfo struct {
	Clz            int
	Atk            float64
	Def            float64
	Hp             float64
	Dodge          float64
	Critical       float64
	CriticalDamage float64
	Speed          float32
	Skill          []int
	Skills         map[string]interface{}
}

var USER_BASIC_INFO map[int]*BasicInfo
var SKILL_BASIC_INFO *table2.SkillTableData = table.SKILL_BASIC_INFO

func InitConst() {
	loadById(string(HUNTER))
	loadById(string(WIZARD))
	loadById(string(KNIGHT))
	loadById(string(GUNNER))
	loadById(string(SOLDIER))
}

func loadById(id string) {
	userAttr := (*table.RAW_HERO_ATTR)[id]
	skillAttr := make(map[string]interface{})
	for _, skillId := range userAttr.Skill {
		key := fmt.Sprintf("skill_%d_cd", skillId)
		skillAttr[key] = 0
	}
	s, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	USER_BASIC_INFO[s] = &BasicInfo{
		Clz:            userAttr.Class,       // 职业
		Atk:            userAttr.Attack,      // 攻击
		Def:            userAttr.Defend,      // 防御
		Hp:             userAttr.HP,          // 血量
		Dodge:          userAttr.Dodge,       // 闪避
		Critical:       userAttr.Crit,        //暴击
		CriticalDamage: userAttr.Crit_Damage, //爆伤
		Speed:          userAttr.Speed,       // 移速
		Skill:          userAttr.Skill,
	}
	USER_BASIC_INFO[s].Skills = skillAttr
}
