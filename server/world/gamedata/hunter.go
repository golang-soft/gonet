package gamedata

import (
	"gonet/server/common"
	"gonet/server/common/data"
)

type (
	SHunter struct {
	}
)

var Hunter *SHunter = &SHunter{}

//猎人
func (this *SHunter) skill_1101(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1101)
	Skill(attack)
}
func (this *SHunter) skill_1102(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1102)
	Skill(attack)
}
func (this *SHunter) skill_1103(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1103)
	Skill(attack)
}
func (this *SHunter) skill_1104(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1104)
	Skill(attack)
}
