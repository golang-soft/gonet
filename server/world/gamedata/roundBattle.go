package gamedata

import (
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/world/helper"
	"math"
	"time"
)

type (
	SBattleCtrl struct {
	}
	ISBattleCtrl interface {
	}
)

var BattleCtrl *SBattleCtrl = &SBattleCtrl{}

//战场
//加血
func (this *SBattleCtrl) addHp(users []string) {

}

//减血
func (this *SBattleCtrl) desHp(round int, from string, to string, skill int32, skillId int32, damage float64, critical bool, isSkip bool) {

	userInfo := GGame.GetUserById(round, to)
	if userInfo == nil {
		return
	}

	//判断护盾， 护盾时间，护盾职
	play := helper.USER_BASIC_INFO[userInfo.Itype]
	var desHp float64 = 0
	var reduceShield float64 = 0
	if !isSkip && userInfo.Shield > 0 && userInfo.ShieldTs > time.Now().Unix() {
		if userInfo.Shield > damage {
			//护盾大于伤害
			reduceShield = damage
		} else {
			//护盾小于伤害减盾减血
			desHp = float64(damage - userInfo.Shield)
			reduceShield = userInfo.Shield
		}
	} else {
		desHp = damage
	}

	if reduceShield > 0 {
		GGame.desShield(round, from, to, reduceShield)
		// await gGame.addDps(round, from, reduceShield)
	}
	if desHp > 0 {
		if userInfo.ImmuneTs > time.Now().Unix() {
			//免疫锁血
			immuneHp := math.Floor(float64(play.Hp * (userInfo.Immune / 100)))
			if userInfo.Hp > immuneHp {
				reduceHp := userInfo.Hp - immuneHp
				if reduceHp < desHp {
					desHp = reduceHp
				}
			} else {
				desHp = 0
			}
		}

		realDamage := math.Floor(desHp)
		GGame.desHp(round, from, to, realDamage, 0)
		// await gGame.addDps(round, from, realDamage)
	}

	//buff 效果判断
	if userInfo.ImmuneTs < time.Now().Unix() {
		this.buffEffect(round, from, to, skill, skillId, *userInfo)
	}

	GvgBattleBroadcastAll(
		"ATTACK", round,
		&cmessage.AttackResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_AttackResp), 0),
			//Error:      int32(1),
			//PlayerId:   0,
			From:     from,
			To:       to,
			Skill:    skill,
			SkillId:  skillId,
			Shield:   userInfo.Shield,
			IsSkip:   desHp == 0,
			Hp:       userInfo.Hp,
			Damage:   desHp,
			Critical: critical,
			Ts:       time.Now().Unix(),
		},
	)

	//反伤
	if !isSkip {
		if userInfo.ThornsTs > time.Now().Unix() {
			this.thornsHp(round, to, from, common.Skill.Skill_1304, math.Floor(damage*(userInfo.Thorns/100)), false)
		}
	}
}

func (this *SBattleCtrl) desHp2(attack *data.AttackData, damage float64, critical bool, isSkip bool) {
	round := attack.Round
	from := attack.From
	to := attack.To
	skill := attack.Skill
	skillId := attack.SkillId

	userInfo := GGame.GetUserById(round, to)
	if userInfo == nil {
		return
	}

	//判断护盾， 护盾时间，护盾职
	play := helper.USER_BASIC_INFO[userInfo.Itype]
	var desHp float64 = 0
	var reduceShield float64 = 0
	if !isSkip && userInfo.Shield > 0 && userInfo.ShieldTs > time.Now().Unix() {
		if userInfo.Shield > damage {
			//护盾大于伤害
			reduceShield = damage
		} else {
			//护盾小于伤害减盾减血
			desHp = float64(damage - userInfo.Shield)
			reduceShield = userInfo.Shield
		}
	} else {
		desHp = damage
	}

	if reduceShield > 0 {
		GGame.desShield(round, from, to, reduceShield)
		// await gGame.addDps(round, from, reduceShield)
	}
	if desHp > 0 {
		if userInfo.ImmuneTs > time.Now().Unix() {
			//免疫锁血
			immuneHp := math.Floor(float64(play.Hp * (userInfo.Immune / 100)))
			if userInfo.Hp > immuneHp {
				reduceHp := userInfo.Hp - immuneHp
				if reduceHp < desHp {
					desHp = reduceHp
				}
			} else {
				desHp = 0
			}
		}

		realDamage := math.Floor(desHp)
		GGame.desHp(round, from, to, realDamage, 0)
		// await gGame.addDps(round, from, realDamage)
	}

	//buff 效果判断
	if userInfo.ImmuneTs < time.Now().Unix() {
		this.buffEffect(round, from, to, skill, skillId, *userInfo)
	}
	GvgBattleBroadcastAll(
		"ATTACK", round,
		&cmessage.AttackResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_AttackResp), 0),
			//Error:      int32(1),
			//PlayerId:   0,
			From:     from,
			To:       to,
			Skill:    skill,
			SkillId:  skillId,
			Shield:   userInfo.Shield,
			IsSkip:   desHp == 0,
			Hp:       userInfo.Hp,
			Damage:   desHp,
			Critical: critical,
			Ts:       time.Now().Unix(),
		},
	)

	//反伤
	if !isSkip {
		if userInfo.ThornsTs > time.Now().Unix() {
			this.thornsHp(round, to, from, common.Skill.Skill_1304, math.Floor(damage*(userInfo.Thorns/100)), false)
		}
	}
}

func (this *SBattleCtrl) buffEffect(round int, from string, to string, skill int32, skillId int32, userInfo data.UserGameAttrData) {
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]

	buffId := skillCfg.Buff[0]
	// buffId = 9
	switch buffId {
	case common.Skill.Buff_1:
		//1-眩晕目标  填眩晕的几率百分比 时间
		if helper.MaybeSuccessPercent(skillCfg.Buff[1]) {
			GGame.addDizzy(round, to, int64(time.Now().Unix()+int64(skillCfg.Buff[2])*1000-300))
		}
		break
	case common.Skill.Buff_2:
		//2-反弹伤害 A为2时，填反弹伤害的比例
		GGame.addThorns(round, to, skillCfg.Buff[1], time.Now().Unix()+int64(skillCfg.Buff[2])*1000)
		break
	case common.Skill.Buff_3:
		//3-免控+锁血A为3时，填锁血的比例
		GGame.addImmune(round, to, skillCfg.Buff[1], time.Now().Unix()+int64(skillCfg.Buff[2])*1000)
		break
	case common.Skill.Buff_4:
		//4-施加护盾A为4时，填护盾的吸收值
		GGame.addShield(round, to, skillCfg.Buff[1], time.Now().Unix()+int64(skillCfg.Buff[2])*1000)
		break
	case common.Skill.Buff_5:
		//5-减少防御   A为5时，填减少的防御百分比
		break
	case common.Skill.Buff_6:
		//6-限制目标移动+dot  A为6时，填dot的伤害间隔
		GGame.addStopMove(round, to, time.Now().Unix()+int64(skillCfg.Buff[2])*1000-300)
		break
	case common.Skill.Buff_7:
		//7-变身（架设炮台），提升攻击力，并且移速降为0    A为7时，填提升的攻击力数值
		GGame.addDeformation(round, to, skillCfg.Buff[1], int64(float64(time.Now().Unix())+skillCfg.Buff[2]*1000))
		break
	case common.Skill.Buff_8:
		//8-持续向前移动，不会被打断，并取消碰撞。A为8时，无buff效果值，填0
		break
	case common.Skill.Buff_9:
		//9-减速  A为9时，填减速百分比
		// await gGame.desSpeed(round, to, skillCfg.Buff[1], Date.now() + skillCfg.Buff[2] * 1000)
		break
	case common.Skill.Buff_10:
		//10-根据已损失的血量百分比，造成伤害 A为10时，填每百分之一掉血额外增伤，伤害基数为释放buff的单位攻击力
		break
	default:
		//0 无效果
	}
}

func (this *SBattleCtrl) removeBuff(round int, from string, to string, skill int32, skillId int32, userInfo data.UserGameAttrData) {
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	if skillCfg == nil {
		return
	}
	buffId := skillCfg.Buff[0]
	// buffId = 9
	switch buffId {
	case common.Skill.Buff_1:
		//1-眩晕目标  填眩晕的几率百分比 时间

		break
	case common.Skill.Buff_2:
		//2-反弹伤害 A为2时，填反弹伤害的比例
		GGame.removeThorns(round, to, skillCfg.Buff[1], skillCfg.Buff[2]+float64(time.Now().Unix()))
		break
	case common.Skill.Buff_3:
		//3-免控+锁血A为3时，填锁血的比例
		// await gGame.addImmune(round, to, skillCfg.Buff[1], skillCfg.Buff[2] + Date.now())
		break
	case common.Skill.Buff_4:
		//4-施加护盾A为4时，填护盾的吸收值
		// await gGame.addShield(round, to, skillCfg.Buff[1], Date.now() + skillCfg.Buff[2])
		break
	case common.Skill.Buff_5:
		//5-减少防御   A为5时，填减少的防御百分比
		break
	case common.Skill.Buff_6:
		//6-限制目标移动+dot  A为6时，填dot的伤害间隔
		GGame.removeStopMove(round, to, skillCfg.Buff[2], 0)
		break
	case common.Skill.Buff_7:
		//7-变身（架设炮台），提升攻击力，并且移速降为0    A为7时，填提升的攻击力数值
		GGame.removeDeformation(round, to, skillCfg.Buff[1], int(skillCfg.Buff[2]))
		break
	case common.Skill.Buff_8:
		//8-持续向前移动，不会被打断，并取消碰撞。A为8时，无buff效果值，填0
		break
	case common.Skill.Buff_9:
		//9-减速  A为9时，填减速百分比
		// await gGame.desSpeed(round, to, skillCfg.Buff[1], skillCfg.Buff[2])
		break
	case common.Skill.Buff_10:
		//10-根据已损失的血量百分比，造成伤害 A为10时，填每百分之一掉血额外增伤，伤害基数为释放buff的单位攻击力
		break
	default:
		//0 无效果
	}
}

//反伤
func (this *SBattleCtrl) thornsHp(round int, from string, to string, skill int32, desHp float64, critical bool) {
	GGame.desHp(round, from, to, math.Floor(float64(desHp)), 0)
	GvgBattleBroadcastAll(
		"ATTACK", round,
		&cmessage.AttackResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_AttackResp), 0),
			From:       from,
			To:         to,
			Skill:      skill,
			IsSkip:     false,
			Hp:         GGame.Game[round].user[to].Hp,
			Damage:     desHp,
			Critical:   critical,
			Ts:         time.Now().Unix(),
		},
	)

}

//加防
func (this *SBattleCtrl) addDef(users []string) {
}

//减防
func (this *SBattleCtrl) desDef(users []string) {
}

//加速
func (this *SBattleCtrl) addSpeed(users []string) {
}

//减速
func (this *SBattleCtrl) desSpeed(users []string) {
}

//眩晕
func (this *SBattleCtrl) addDizzy(users []string) {
}

//减晕 （移除）
func (this *SBattleCtrl) removeDizzy(users []string) {
}

//伤害反弹
func (this *SBattleCtrl) addThorns(users []string) {
}

//移除伤害反弹
func (this *SBattleCtrl) removeThorns(users []string) {
}

//加护盾
func (this *SBattleCtrl) addShield(users []string) {
}

//减护盾
func (this *SBattleCtrl) desShield(round int, to string, desNum int) {
	GGame.GetUserById(round, to)
}

//移除护盾
func (this *SBattleCtrl) removeShield(users []string) {

}

//添加免疫
func (this *SBattleCtrl) addImmune(users []string) {

}

//移除免疫
func (this *SBattleCtrl) removeImmune(users []string) {

}

//复活
func (this *SBattleCtrl) revive(users []string) {

}
