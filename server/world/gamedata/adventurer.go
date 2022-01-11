package gamedata

import (
	"gonet/server/common"
	"gonet/server/common/data"
)

type (
	SAdventurer struct {
	}
)

var Adventurer *SAdventurer = &SAdventurer{}

//魔炮
func (this *SAdventurer) skill_1401(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1401)
	//gamedata.Skill(attack)
}
func (this *SAdventurer) skill_1402(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1402)
	//gamedata.Skill(attack)
}
func (this *SAdventurer) skill_1403(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1403)
	//gamedata.Skill(attack)
}
func (this *SAdventurer) skill_1404(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1404)
	//gamedata.Skill(attack)
}
