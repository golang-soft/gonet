package gamedata

import (
	"gonet/server/common"
	"gonet/server/common/data"
)

//骑士
type (
	SWarrior struct {
	}
)

var Warrior *SWarrior = &SWarrior{}

//战士
func (this *SWarrior) skill_1301(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1301)
	Skill(attack)
}
func (this *SWarrior) skill_1302(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1302)
	Skill(attack)
}
func (this *SWarrior) skill_1303(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1303)
	Skill(attack)
}
func (this *SWarrior) skill_1304(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1304)
	Skill(attack)
}
