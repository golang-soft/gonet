package gamedata

import (
	"gonet/server/common"
	"gonet/server/common/data"
)

type (
	SRogue struct {
	}
)

var Rogue *SRogue = &SRogue{}

//战士
func (this *SRogue) skill_1501(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1501)
	Skill(attack)
}
func (this *SRogue) skill_1502(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1502)
	Skill(attack)
}
func (this *SRogue) skill_1503(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1503)
	Skill(attack)
}
func (this *SRogue) skill_1504(attack data.AttackData) {
	attack.SkillId = int32(common.Skill.Skill_1504)
	Skill(attack)
}
