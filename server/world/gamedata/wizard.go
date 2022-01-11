package gamedata

import (
	"gonet/server/common"
	"gonet/server/common/data"
)

//魔法（男巫）
type (
	SWizard struct {
	}
)

var Wizard *SWizard = &SWizard{}

//战士
func (this *SWizard) skill_1201(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1201)
	Skill(attack)
}
func (this *SWizard) skill_1202(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1202)
	Skill(attack)
}
func (this *SWizard) skill_1203(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1203)
	Skill(attack)
}
func (this *SWizard) skill_1204(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1204)
	Skill(attack)
}
