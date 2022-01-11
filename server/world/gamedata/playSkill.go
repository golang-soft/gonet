package gamedata

import (
	"gonet/common/timer"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/table"
	"gonet/server/world/datafnc"
	"gonet/server/world/helper"
	"math"
	"time"
	"unsafe"
)

//玩家技能
func Skill(attack data.AttackData) bool {
	skillCfg := (*helper.SKILL_BASIC_INFO)[attack.SkillId]
	atkTimes := skillCfg.Last_Time
	if atkTimes != 0 && atkTimes >= 1 {
		for index := 0; index < atkTimes; index++ {
			deBattleDelay(attack, int64(index+1))
		}
	} else {
		deBattleImmed(attack)
	}
	return true
}

const skill_interval = 1000

func deBattleDelay(attack data.AttackData, time int64) {
	switch attack.SkillId {
	case int32(common.Skill.Skill_1101):
		setTimeout(skill_1101, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1102):
		setTimeout(skill_1102, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1103):
		setTimeout(skill_1103, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1104):
		setTimeout(skill_1104, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1201):
		setTimeout(skill_1201, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1202):
		setTimeout(skill_1202, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1203):
		setTimeout(skill_1203, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1204):
		setTimeout(skill_1204, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1301):
		setTimeout(skill_1301, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1302):
		setTimeout(skill_1302, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1303):
		setTimeout(skill_1303, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1304):
		setTimeout(skill_1304, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1401):
		setTimeout(skill_1401, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1402):
		setTimeout(skill_1402, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1403):
		setTimeout(skill_1403, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1404):
		setTimeout(skill_1404, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1501):
		setTimeout(skill_1501, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1502):
		setTimeout(skill_1502, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1503):
		setTimeout(skill_1503, time*skill_interval, attack)
		break
	case int32(common.Skill.Skill_1504):
		setTimeout(skill_1504, time*skill_interval, attack)
		break
	default:
		//console.log("not fund skill");
	}
}

func setTimeout(skillfunc func(attack data.AttackData), interval int64, attack data.AttackData) {
	//TODO setTimeout
	RegisterTimer(time.Duration(interval), skillfunc, attack)
}

func setTimeout2(skillfunc func(attack *data.AttackData, damage float64, critical bool, isSkip bool), interval int64, attack data.AttackData, damage float64, critical bool, isSkip bool) {
	//TODO setTimeout
	RegisterTimer2(time.Duration(interval), skillfunc, attack, damage, critical, isSkip)
}

func RegisterTimer(duration time.Duration, fun func(attack data.AttackData), attack data.AttackData, opts ...timer.OpOption) {
	timerId := time.Now().Unix()
	timer.RegisterTimer(&timerId, duration, func() {
		func1 := (*func(attack data.AttackData))(unsafe.Pointer(&fun))
		(*func1)(attack)
	}, opts...)
}

func RegisterTimer2(duration time.Duration, fun func(attack *data.AttackData, damage float64, critical bool, isSkip bool), attack data.AttackData, damage float64, critical bool, isSkip bool, opts ...timer.OpOption) {
	timerId := time.Now().Unix()
	timer.RegisterTimer(&timerId, duration, func() {
		func1 := (*func(attack *data.AttackData, damage float64, critical bool, isSkip bool))(unsafe.Pointer(&fun))
		(*func1)(&attack, damage, critical, isSkip)
	}, opts...)
}

//瞬发
func deBattleImmed(attack data.AttackData) {
	switch attack.SkillId {
	case int32((common.Skill.Skill_1101)):
		// setTimeout(skill_1101, time * skill_interval, attack, fromAttr)
		skill_1101(attack)
		break
	case int32(common.Skill.Skill_1102):
		// setTimeout(skill_1102, time * skill_interval, attack, fromAttr)
		skill_1102(attack)
		break
	case int32(common.Skill.Skill_1103):
		// setTimeout(skill_1103, time * skill_interval, attack, fromAttr)
		skill_1103(attack)
		break
	case int32(common.Skill.Skill_1104):
		// setTimeout(skill_1104, time * skill_interval, attack, fromAttr)
		skill_1104(attack)
		break
	case int32(common.Skill.Skill_1201):
		// setTimeout(skill_1201, time * skill_interval, attack, fromAttr)
		skill_1201(attack)
		break
	case int32(common.Skill.Skill_1202):
		// setTimeout(skill_1202, time * skill_interval, attack, fromAttr)
		skill_1202(attack)
		break
	case int32(common.Skill.Skill_1203):
		// setTimeout(skill_1203, time * skill_interval, attack, fromAttr)
		skill_1203(attack)
		break
	case int32(common.Skill.Skill_1204):
		// setTimeout(skill_1204, time * skill_interval, attack, fromAttr)
		skill_1204(attack)
		break
	case int32(common.Skill.Skill_1301):
		// setTimeout(skill_1301, time * skill_interval, attack, fromAttr)
		skill_1301(attack)
		break
	case int32(common.Skill.Skill_1302):
		// setTimeout(skill_1302, time * skill_interval, attack, fromAttr)
		skill_1302(attack)
		break
	case int32(common.Skill.Skill_1303):
		// setTimeout(skill_1303, time * skill_interval, attack, fromAttr)
		skill_1303(attack)
		break
	case int32(common.Skill.Skill_1304):
		// setTimeout(skill_1304, time * skill_interval, attack, fromAttr)
		skill_1304(attack)
		break
	case int32(common.Skill.Skill_1401):
		// setTimeout(skill_1401, time * skill_interval, attack, fromAttr)
		skill_1401(attack)
		break
	case int32(common.Skill.Skill_1402):
		// setTimeout(skill_1402, time * skill_interval, attack, fromAttr)
		skill_1402(attack)
		break
	case int32(common.Skill.Skill_1403):
		// setTimeout(skill_1403, time * skill_interval, attack, fromAttr)
		skill_1403(attack)
		break
	case int32(common.Skill.Skill_1404):
		// setTimeout(skill_1404, time * skill_interval, attack, fromAttr)
		skill_1404(attack)
		break
	case int32(common.Skill.Skill_1501):
		// setTimeout(skill_1501, time * skill_interval, attack, fromAttr)
		skill_1501(attack)
		break
	case int32(common.Skill.Skill_1502):
		// setTimeout( skill_1502, time * skill_interval, attack, fromAttr)
		skill_1502(attack)
		break
	case int32(common.Skill.Skill_1503):
		// setTimeout(skill_1503, time * skill_interval, attack, fromAttr)
		skill_1503(attack)
		break
	case int32(common.Skill.Skill_1504):
		// setTimeout(skill_1504, time * skill_interval, attack, fromAttr)
		skill_1504(attack)
		break
	default:
		//console.error("not fund skill");
	}
}

/*
   1-攻击
   2-防御
   3-血量
   4-暴击
   5-爆伤
   6-闪避
   9-全属性

*/

func getPlayAtk(fromAttr *helper.BasicInfo, skillCfg *table.SkillCfg, userInfo *data.UserGameAttrData, atkbuff int) float64 {
	var percent float64 = 0
	var buffAtk float64 = 0
	var equipAtk = 0
	if userInfo.AllAttr != 0 {
		//道具全属性增益
		percent = datafnc.All_Attr_Percent
	}

	//魔炮攻击buff
	buffAtk = datafnc.Skill_1402_Buff_AddAtk_Percent

	equipAtk = userInfo.Equip_1
	percent += userInfo.Equip_9 / 100

	atk := ((fromAttr.Atk + float64(equipAtk)) * float64(1+float64(percent)+float64(buffAtk))) *
		skillCfg.Rate / float64(common.Skill.Base_rate) * float64(common.Skill.AddAtk)
	return math.Floor(atk)
}

func getPlayDef(toAttr *helper.BasicInfo, userInfo *data.UserGameAttrData) float64 {
	var percent float64 = 0
	var equipDef float64 = 0
	if userInfo.AllAttr != 0 {
		//道具全属性增益
		percent = datafnc.All_Attr_Percent
	}

	equipDef = userInfo.Equip_2
	percent += userInfo.Equip_9 / 100

	return math.Floor((toAttr.Def + equipDef) * (1 + float64(percent)))
}

func getCritical(toAttr *helper.BasicInfo, userInfo *data.UserGameAttrData) float64 {
	var percent float64 = 0
	var equipCrit float64 = 0
	if userInfo.AllAttr != 0 {
		//道具全属性增益
		percent = datafnc.All_Attr_Percent
	}

	equipCrit = userInfo.Equip_4
	percent += userInfo.Equip_9 / 100

	return math.Floor((toAttr.Critical + equipCrit) * (1 + percent))
}

func getPlayDod(toAttr *helper.BasicInfo, userInfo *data.UserGameAttrData) float64 {
	var percent float64 = 0
	var equipDod float64 = 0
	if userInfo.AllAttr != 0 {
		//道具全属性增益
		percent = datafnc.All_Attr_Percent
	}
	equipDod = userInfo.Equip_5
	percent += userInfo.Equip_9 / 100

	return math.Floor((toAttr.Dodge + equipDod) * (1 + percent))
}

func getCriticalDamage(fromDamage float64, fromAttr *helper.BasicInfo, userInfo *data.UserGameAttrData) float64 {
	//暴击伤害 atk = （1+ （criticalDamage/2000））
	var percent float64 = 0
	var equipCritDmg float64 = 0
	if userInfo.AllAttr != 0 {
		//道具全属性增益
		percent = datafnc.All_Attr_Percent
	}

	equipCritDmg = userInfo.Equip_6

	percent += userInfo.Equip_9 / 100

	return math.Floor(fromDamage * (1 + (fromAttr.CriticalDamage+equipCritDmg)*(1+percent)/common.Attr.Crit_damage_base))
}

func getDamage(fromDamage float64, toDefend float64) float64 {
	if datafnc.Battle_Mod == common.Battle_mode.Defence {
		//除
		return math.Floor(fromDamage * (1 - toDefend/(toDefend+datafnc.Battle_Defence)))
	}

	return fromDamage - toDefend
}

//一次迅捷地射击,造成90点伤害。
func skill_1101(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	//skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg

	var isSkip = false
	var isCritical = false
	var fromDamage float64 = 0
	var toDefend float64 = 0

	var damage float64 = 0
	fromUser := GGame.GetUserById(round, from)
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]

	toPlayDod := getPlayDod(toAttr, toUser)
	// 闪避
	if helper.MaybeSuccess(toPlayDod) {
		fromDamage = 0
		isSkip = true
	} else {
		fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
		toDefend = getPlayDef(toAttr, toUser)
		//暴击
		fromCritical := getCritical(fromAttr, fromUser)
		if helper.MaybeSuccess(fromCritical) {
			fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
			isCritical = true
		}
		damage = getDamage(fromDamage, toDefend)
	}

	targetPoint := PositionCtrl.updateNewPosition(to, round)
	atkTime := PositionCtrl.attackSpendTime(data.PointData{X: fromUser.X, Y: fromUser.Y}, data.PointData{X: targetPoint.X, Y: targetPoint.Y}) / skillCfg.Bullet_speed

	setTimeout2(
		BattleCtrl.desHp2,
		int64(math.Floor(atkTime*1000)-20),
		attack, damage, isCritical, isSkip,
	)
}

//强劲的射击，对敌人造成207点伤害。
func skill_1102(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	//skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg

	var isSkip = false
	var isCritical = false
	var fromDamage float64 = 0
	var toDefend float64 = 0

	var damage float64 = 0
	fromUser := GGame.GetUserById(round, from)
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}

	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]

	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	toDefend = getPlayDef(toAttr, toUser)
	//暴击
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	damage = getDamage(fromDamage, toDefend)
	targetPoint := PositionCtrl.updateNewPosition(to, round)

	atkTime := PositionCtrl.attackSpendTime(data.PointData{X: fromUser.X, Y: fromUser.Y}, data.PointData{X: targetPoint.X, Y: targetPoint.Y}) / skillCfg.Bullet_speed

	setTimeout2(
		BattleCtrl.desHp2,
		int64(math.Floor(atkTime*1000)-20),
		attack, damage, isCritical, isSkip,
	)
}

//发射数支箭矢，对前方敌人造成144点范围伤害。 扇形
func skill_1103(attack data.AttackData) {
	from := attack.From
	//to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg

	var fromDamage float64 = 0
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]

	User := GGame.getUsersByRound(round)
	isCritical := false
	var damage float64 = 0
	isSkip := false
	fromUser := GGame.GetUserById(round, from)
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	fromCritical := getCritical(fromAttr, fromUser)
	//暴击
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	rLen := skillCfg.Radius[0]
	skillAngle := skillCfg.Radius[1]
	startPoint := data.PointData{X: fromUser.X, Y: fromUser.Y}
	direction := fromUser.Direction
	for key, _ := range User {
		toUser := GGame.GetUserById(round, key)
		if toUser == nil || key == from || toUser.Part == fromUser.Part || toUser.Hp <= 0 {
			continue
		}

		var toDefend float64 = 0
		toAttr := helper.USER_BASIC_INFO[toUser.Itype]

		numData := data.UserPositionData{
			Speed:     toUser.Speed,
			Direction: toUser.Direction,
			Barrier:   toUser.Barrier,
		}
		pos := PositionCtrl.updateSkillPosition(toUser.User, round, numData)

		GvgBattleBroadcastAll("MOVING", round,
			&cmessage.MovingResp{
				PacketHead:    common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_MovingResp), 0),
				User:          toUser.User,
				Speed:         pos.Speed,
				ReduceSpeedTs: pos.ReduceSpeedTs,
				Direction:     pos.Direction,
				Barrier:       pos.Barrier,
				Dizzy:         pos.Dizzy,
				DizzyTs:       pos.DizzyTs,
				X:             pos.X,
				Y:             pos.Y,
				StopMove:      pos.StopMove,
				StopMoveTs:    pos.StopMoveTs,
				PosUpdateTs:   pos.PosUpdateTs,
			},
		)
		toUserPoint := data.PointData{X: pos.X, Y: pos.Y}
		if PositionCtrl.IsPointInFan(startPoint, direction, float64(skillAngle), float64(rLen), toUserPoint) {
			toDefend = getPlayDef(toAttr, toUser)
			damage = getDamage(fromDamage, toDefend)
			BattleCtrl.desHp(round, from, key, skill, skillId, math.Floor(damage), isCritical, isSkip)
		}
	}
}

//对目标敌人造成270点伤害，然后将其固定在原地3秒。
func skill_1104(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	//skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg

	var fromDamage float64 = 0
	var toDefend float64 = 0
	var isSkip = false
	var isCritical = false
	var fromUser = GGame.GetUserById(round, from)

	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, to)

	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	toDefend = getPlayDef(toAttr, toUser)
	//暴击
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}

	damage := getDamage(fromDamage, toDefend)
	targetPoint := PositionCtrl.updateNewPosition(to, round)
	atkTime := PositionCtrl.attackSpendTime(data.PointData{X: fromUser.X, Y: fromUser.Y},
		data.PointData{X: targetPoint.X, Y: targetPoint.Y}) / skillCfg.Bullet_speed

	setTimeout2(
		BattleCtrl.desHp2,
		int64(math.Floor(atkTime*1000)-20),
		attack, damage, isCritical, isSkip,
	)

}

//一次快速的法术攻击,造成95点伤害
func skill_1201(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	//skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var fromDamage float64 = 0
	var toDefend float64 = 0
	var isSkip = false
	var isCritical = false
	var damage float64 = 0

	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	// 闪避
	toPlayDod := getPlayDod(toAttr, toUser)
	if helper.MaybeSuccess(toPlayDod) {
		fromDamage = 0
		isSkip = true
	} else {
		fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
		toDefend = getPlayDef(toAttr, toUser)
		//暴击
		fromCritical := getCritical(fromAttr, fromUser)
		if helper.MaybeSuccess(fromCritical) {
			fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
			isCritical = true
		}
		damage = getDamage(fromDamage, toDefend)
	}
	targetPoint := PositionCtrl.updateNewPosition(to, round)
	atkTime := PositionCtrl.attackSpendTime(data.PointData{X: fromUser.X, Y: fromUser.Y}, data.PointData{X: targetPoint.X, Y: targetPoint.Y}) / skillCfg.Bullet_speed

	setTimeout2(
		BattleCtrl.desHp2,
		int64(math.Floor(atkTime*1000)-20),
		attack, damage, isCritical, isSkip,
	)
}

//召唤风雪的力量，5秒内对指定区域内的敌人施加每秒114点的持续伤害，并造成50%减速效果，持续1秒。圆
func skill_1202(attack data.AttackData) {
	from := attack.From
	//to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	x := attack.X
	y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var fromDamage float64 = 0
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	rLen := skillCfg.Radius[0]
	User := GGame.getUsersByRound(round)
	var isSkip = false
	var isCritical = false
	var damage float64 = 0
	fromUser := GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	for key, _ := range User {
		toUser := GGame.GetUserById(round, key)
		if toUser == nil || key == from || toUser.Part == fromUser.Part || toUser.Hp <= 0 {
			continue
		}
		var toDefend float64 = 0
		toAttr := helper.USER_BASIC_INFO[toUser.Itype]
		numData := data.UserPositionData{
			Speed:     toUser.Speed,
			Direction: toUser.Direction,
			Barrier:   toUser.Barrier,
		}
		pos := PositionCtrl.updateSkillPosition(toUser.User, round, numData)
		//gvgBattleBroadcastAll(USER_EVENT.USER.MOVING, round, { ...pos, user: toUser.user })

		GvgBattleBroadcastAll("MOVING", round,
			&cmessage.MovingResp{
				PacketHead:    common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_MovingResp), 0),
				User:          toUser.User,
				Speed:         pos.Speed,
				ReduceSpeedTs: pos.ReduceSpeedTs,
				Direction:     pos.Direction,
				Barrier:       pos.Barrier,
				Dizzy:         pos.Dizzy,
				DizzyTs:       pos.DizzyTs,
				X:             pos.X,
				Y:             pos.Y,
				StopMove:      pos.StopMove,
				StopMoveTs:    pos.StopMoveTs,
				PosUpdateTs:   pos.PosUpdateTs,
			},
		)

		if (PositionCtrl.IsPointInCircle(*pos, data.Pos{X: x, Y: y}, rLen)) {
			toDefend = getPlayDef(toAttr, toUser)
			damage = getDamage(fromDamage, toDefend)
			BattleCtrl.desHp(round, from, key, skill, skillId, math.Floor(damage), isCritical, isSkip)
		}
	}
}

//扔出一发寒冰箭，对目标造成209点伤害，并减速50%。减速持续3秒。
func skill_1203(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	//skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var isSkip = false
	var isCritical = false
	var damage float64 = 0
	var fromDamage float64 = 0
	var toDefend float64 = 0
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	toDefend = getPlayDef(toAttr, toUser)
	//暴击
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	damage = getDamage(fromDamage, toDefend)
	targetPoint := PositionCtrl.updateNewPosition(to, round)
	atkTime := PositionCtrl.attackSpendTime(data.PointData{X: fromUser.X, Y: fromUser.Y}, data.PointData{X: targetPoint.X, Y: targetPoint.Y}) / skillCfg.Bullet_speed

	setTimeout2(
		BattleCtrl.desHp2,
		int64(math.Floor(atkTime*1000)-20),
		attack, damage, isCritical, isSkip,
	)
}

//施展法术为自己施加护盾。吸收380点伤害。
func skill_1204(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	// 护盾
	BattleCtrl.buffEffect(round, from, from, skill, skillId, *toUser)

	GvgBattleBroadcastAll("SKILL_SHIELD", round,
		&cmessage.SkillShieldResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_SkillShieldResp), 0),
			User:       from,
			Shield:     toUser.Shield,
			ShieldTs:   toUser.ShieldTs,
		},
	)

}

//一次灵敏地挥剑,造成100点伤害。
func skill_1301(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var isSkip = false
	var isCritical = false
	var fromDamage float64 = 0
	var toDefend float64 = 0
	var damage float64 = 0
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	// 闪避
	toPlayDod := getPlayDod(toAttr, toUser)
	if helper.MaybeSuccess(toPlayDod) {
		fromDamage = 0
		isSkip = true
	} else {
		fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
		toDefend = getPlayDef(toAttr, toUser)
		//暴击
		fromCritical := getCritical(fromAttr, fromUser)
		if helper.MaybeSuccess(fromCritical) {
			fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
			isCritical = true
		}
		damage = getDamage(fromDamage, toDefend)
	}
	BattleCtrl.desHp(round, from, to, skill, skillId, damage, isCritical, isSkip)
}

//对目标敌人造成120点伤害，并根据其损失的血量百分比造成额外伤害，最高额外造成576点伤害。
func skill_1302(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var isSkip = false
	var isCritical = false
	var fromDamage float64 = 0
	var toDefend float64 = 0
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	//血量百分比
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	baseAtk := skillCfg.Buff[1]
	percent := math.Floor((toAttr.Hp - toUser.Hp) / toAttr.Hp * 100)
	//额外增伤
	fromDamage += baseAtk * percent
	toDefend = getPlayDef(toAttr, toUser)
	// }
	//暴击
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	//额外伤害判定
	damage := getDamage(fromDamage, toDefend)
	BattleCtrl.desHp(round, from, to, skill, skillId, damage, isCritical, isSkip)
}

//将强大的能量灌入你脚下的土地，有30%几率对敌人造成眩晕，持续1秒。 圆
func skill_1303(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	rLen := skillCfg.Radius[0]
	User := GGame.getUsersByRound(round)
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}

	var isCritical = false
	var damage float64 = 0
	var isSkip = false
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	damage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	for key, _ := range User {
		toUser := GGame.GetUserById(round, key)
		if toUser == nil || key == from || toUser.Part == fromUser.Part || toUser.Hp <= 0 {
			continue
		}

		numData := data.UserPositionData{
			Speed:     toUser.Speed,
			Direction: toUser.Direction,
			Barrier:   toUser.Barrier,
		}
		pos := PositionCtrl.updateSkillPosition(toUser.User, round, numData)

		GvgBattleBroadcastAll("MOVING", round,
			&cmessage.MovingResp{
				PacketHead:    common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_MovingResp), 0),
				User:          toUser.User,
				Speed:         pos.Speed,
				ReduceSpeedTs: pos.ReduceSpeedTs,
				Direction:     pos.Direction,
				Barrier:       pos.Barrier,
				Dizzy:         pos.Dizzy,
				DizzyTs:       pos.DizzyTs,
				X:             pos.X,
				Y:             pos.Y,
				StopMove:      pos.StopMove,
				StopMoveTs:    pos.StopMoveTs,
				PosUpdateTs:   pos.PosUpdateTs,
			},
		)
		if (PositionCtrl.IsPointInCircle(*pos, data.Pos{X: fromUser.X, Y: fromUser.Y}, rLen)) {
			//眩晕不闪
			BattleCtrl.desHp(round, from, to, skill, skillId, damage, isCritical, isSkip)
		}
	}
}

//凝聚强大的力量护佑自身，使攻击你的敌人受到50%反弹伤害，持续8秒。
func skill_1304(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	// 反伤
	BattleCtrl.buffEffect(round, from, from, skill, skillId, *toUser)

	GvgBattleBroadcastAll("SKILL_THORNS", round,
		&cmessage.SkillThornsResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_SkillThornsResp), 0),
			User:       toUser.User,
			Thorns:     toUser.Thorns,
			ThornsTs:   toUser.ThornsTs,
		},
	)
}

//一次沉重地炮击，造成110点伤害。
func skill_1401(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	//skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var isSkip = false
	var isCritical = false
	var fromDamage float64 = 0
	var toDefend float64 = 0
	var damage float64 = 0
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	// 闪避
	toPlayDod := getPlayDod(toAttr, toUser)
	if helper.MaybeSuccess(toPlayDod) {
		fromDamage = 0
		isSkip = true
	} else {
		fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, int(fromUser.AddAtk))
		toDefend = getPlayDef(toAttr, toUser)
		//暴击
		fromCritical := getCritical(fromAttr, fromUser)
		if helper.MaybeSuccess(fromCritical) {
			fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
			isCritical = true
		}
		damage = getDamage(fromDamage, toDefend)
	}
	targetPoint := PositionCtrl.updateNewPosition(to, round)
	atkTime := PositionCtrl.attackSpendTime(data.PointData{X: fromUser.X, Y: fromUser.Y}, data.PointData{X: targetPoint.X, Y: targetPoint.Y}) / skillCfg.Bullet_speed

	setTimeout2(
		BattleCtrl.desHp2,
		int64(math.Floor(atkTime*1000)-20),
		attack, damage, isCritical, isSkip,
	)

}

//架设一座炮台，自身将不能移动，攻击力提升40%/拆卸一座炮台。
func skill_1402(attack data.AttackData) {
	from := attack.From
	//to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}

	// 变身 加伤害
	if fromUser.AddAtk > 0 {
		//移除变身
		BattleCtrl.removeBuff(round, from, from, skill, skillId, *fromUser)
	} else {
		//添加变身
		BattleCtrl.buffEffect(round, from, from, skill, skillId, *fromUser)
	}

	GvgBattleBroadcastAll("SKILL_DEFORMATION", round,
		&cmessage.SkillDeformationResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_SkillDeformationResp), 0),
			User:       from,
			AddAtk:     fromUser.AddAtk,
		},
	)
}

//发射缓慢移动的弹药，在6秒内对指定区域内的所有敌人造成每秒220点的持续伤害。 圆
func skill_1403(attack data.AttackData) {
	from := attack.From
	//to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	x := attack.X
	y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var fromDamage float64 = 0
	var isSkip = false
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	var rLen = skillCfg.Radius[0]
	User := GGame.getUsersByRound(round)
	var isCritical = false

	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}

	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, int(fromUser.AddAtk))
	//暴击
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	for key, _ := range User {
		toUser := GGame.GetUserById(round, key)
		if toUser == nil || key == from || toUser.Part == fromUser.Part || toUser.Hp <= 0 {
			continue
		}

		var toDefend float64 = 0
		toAttr := helper.USER_BASIC_INFO[toUser.Itype]
		numData := data.UserPositionData{
			Speed:     toUser.Speed,
			Direction: toUser.Direction,
			Barrier:   toUser.Barrier,
		}
		pos := PositionCtrl.updateSkillPosition(toUser.User, round, numData)

		GvgBattleBroadcastAll("MOVING", round,
			&cmessage.MovingResp{
				PacketHead:    common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_MovingResp), 0),
				User:          toUser.User,
				Speed:         pos.Speed,
				ReduceSpeedTs: pos.ReduceSpeedTs,
				Direction:     pos.Direction,
				Barrier:       pos.Barrier,
				Dizzy:         pos.Dizzy,
				DizzyTs:       pos.DizzyTs,
				X:             pos.X,
				Y:             pos.Y,
				StopMove:      pos.StopMove,
				StopMoveTs:    pos.StopMoveTs,
				PosUpdateTs:   pos.PosUpdateTs,
			},
		)
		if (PositionCtrl.IsPointInCircle(*pos, data.Pos{X: x, Y: y}, rLen)) {
			toDefend = getPlayDef(toAttr, toUser)
			damage := getDamage(fromDamage, toDefend)
			BattleCtrl.desHp(round, from, key, skill, skillId, math.Floor(damage), isCritical, isSkip)

		}
	}
}

//向前方持续发射穿刺炮弹，对路径上的所有人造成每秒396点持续伤害。 矩形
func skill_1404(attack data.AttackData) {
	from := attack.From
	//to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var fromDamage float64 = 0
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	var direction float64 = 0
	User := GGame.getUsersByRound(round)
	var isCritical = false
	//let damage = 0
	var isSkip = false
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}

	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, int(fromUser.AddAtk))
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	direction = fromUser.Direction
	for key, _ := range User {
		toUser := GGame.GetUserById(round, key)
		if toUser == nil || key == from || toUser.Hp <= 0 {
			continue
		}
		toAttr := helper.USER_BASIC_INFO[toUser.Itype]
		numData := data.UserPositionData{
			Speed:     toUser.Speed,
			Direction: toUser.Direction,
			Barrier:   toUser.Barrier,
		}
		pos := PositionCtrl.updatePosition(toUser.User, round, numData)
		toUserPoint := data.PointData{X: pos.X, Y: pos.Y}
		GvgBattleBroadcastAll("MOVING", round,
			&cmessage.MovingResp{
				PacketHead:    common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_MovingResp), 0),
				User:          toUser.User,
				Speed:         pos.Speed,
				ReduceSpeedTs: pos.ReduceSpeedTs,
				Direction:     pos.Direction,
				Barrier:       pos.Barrier,
				Dizzy:         pos.Dizzy,
				DizzyTs:       pos.DizzyTs,
				X:             pos.X,
				Y:             pos.Y,
				StopMove:      pos.StopMove,
				StopMoveTs:    pos.StopMoveTs,
				PosUpdateTs:   pos.PosUpdateTs,
			},
		)
		isPointIn, _ := PositionCtrl.IsPointInMatrix(data.PointData{X: fromUser.X, Y: fromUser.Y}, direction, float64(skillCfg.Radius[0]), float64(skillCfg.Radius[1]), toUserPoint)
		if isPointIn {
			toDefend := getPlayDef(toAttr, toUser)
			damage := getDamage(fromDamage, toDefend)
			BattleCtrl.desHp(round, from, key, skill, skillId, math.Floor(damage), isCritical, isSkip)
		}
	}
}

//一次沉重地打击，造成85点伤害。
func skill_1501(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var isSkip = false
	var isCritical = false
	var fromDamage float64 = 0
	var toDefend float64 = 0
	var damage float64 = 0
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, to)
	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	// 闪避
	toPlayDod := getPlayDod(toAttr, toUser)
	if helper.MaybeSuccess(toPlayDod) {
		fromDamage = 0
		isSkip = true
	} else {
		fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
		toDefend = getPlayDef(toAttr, toUser)
		//暴击
		fromCritical := getCritical(fromAttr, fromUser)
		if helper.MaybeSuccess(fromCritical) && !isSkip {
			fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
			isCritical = true
		}
		damage = getDamage(fromDamage, toDefend)
	}

	BattleCtrl.desHp(round, from, to, skill, skillId, damage, isCritical, isSkip)
}

//旋转自身，向前突进，途中免疫控制，并对前进路线范围内的敌人造成204点伤害。  矩形
func skill_1502(attack data.AttackData) {
	from := attack.From
	//to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var fromDamage float64 = 0
	var isSkip = false
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]
	var direction float64 = 0
	User := GGame.getUsersByRound(round)
	var isCritical = false
	//let damage = 0
	var fromUser = GGame.GetUserById(round, from)
	if fromUser == nil {
		return
	}

	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	if fromUser.AddAtk > 0 {
		fromDamage += math.Floor(fromDamage * (fromUser.AddAtk / 100))
	}
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	direction = fromUser.Direction
	for key, _ := range User {
		toUser := GGame.GetUserById(round, key)
		if toUser == nil || key == from || toUser.Part == fromUser.Part {
			continue
		}

		var toDefend float64 = 0
		toAttr := helper.USER_BASIC_INFO[toUser.Itype]
		numData := data.UserPositionData{
			Speed:     toUser.Speed,
			Direction: toUser.Direction,
			Barrier:   toUser.Barrier,
		}
		pos := PositionCtrl.updatePosition(toUser.User, round, numData)
		toUserPoint := data.PointData{X: pos.X, Y: pos.Y}
		GvgBattleBroadcastAll("MOVING", round,
			&cmessage.MovingResp{
				PacketHead:    common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_MovingResp), 0),
				User:          toUser.User,
				Speed:         pos.Speed,
				ReduceSpeedTs: pos.ReduceSpeedTs,
				Direction:     pos.Direction,
				Barrier:       pos.Barrier,
				Dizzy:         pos.Dizzy,
				DizzyTs:       pos.DizzyTs,
				X:             pos.X,
				Y:             pos.Y,
				StopMove:      pos.StopMove,
				StopMoveTs:    pos.StopMoveTs,
				PosUpdateTs:   pos.PosUpdateTs,
			},
		)
		isPointIn, _ := PositionCtrl.IsPointInMatrix(data.PointData{X: fromUser.X, Y: fromUser.Y}, direction, float64(skillCfg.Radius[0]), float64(skillCfg.Radius[1]), toUserPoint)
		if isPointIn {
			toDefend = getPlayDef(toAttr, toUser)
			damage := getDamage(fromDamage, toDefend)
			BattleCtrl.desHp(round, from, key, skill, skillId, math.Floor(damage), isCritical, isSkip)
		}
	}
}

//汇聚愤怒的力量，对单体目标造成255点伤害。
func skill_1503(attack data.AttackData) {
	from := attack.From
	to := attack.To
	round := attack.Round
	skill := attack.Skill
	skillId := attack.SkillId
	//x := attack.X
	//y := attack.Y
	//z := attack.Z
	//msg := attack.Msg
	var isSkip = false
	var isCritical = false
	var fromDamage float64 = 0
	var toDefend float64 = 0
	var damage float64 = 0
	fromUser := GGame.GetUserById(round, attack.From)

	if fromUser == nil {
		return
	}
	toUser := GGame.GetUserById(round, attack.To)

	if toUser == nil {
		return
	}
	fromAttr := helper.USER_BASIC_INFO[fromUser.Itype]
	//toAttr := helper.USER_BASIC_INFO[toUser.Itype]
	skillCfg := (*helper.SKILL_BASIC_INFO)[attack.SkillId]

	fromDamage = getPlayAtk(fromAttr, skillCfg, fromUser, 0)
	//toDefend := getPlayDef(toAttr, toUser)
	//暴击
	fromCritical := getCritical(fromAttr, fromUser)
	if helper.MaybeSuccess(fromCritical) && !isSkip {
		fromDamage = getCriticalDamage(fromDamage, fromAttr, fromUser)
		isCritical = true
	}
	damage = getDamage(fromDamage, toDefend)
	BattleCtrl.desHp(round, from, to, skill, skillId, damage, isCritical, isSkip)
}

//进入极端愤怒状态，免疫一切控制，使自身血量无法下降到10%以下。持续5秒。
func skill_1504(attack data.AttackData) {
	fromUser := GGame.GetUserById(attack.Round, attack.From)

	if fromUser == nil {
		return
	}
	//无敌buff
	BattleCtrl.buffEffect(attack.Round, attack.From, attack.From, attack.Skill, attack.SkillId, *fromUser)

	GvgBattleBroadcastAll("SKILL_IMMUNE", attack.Round,
		&cmessage.SkillImmuneResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_SkillImmuneResp), 0),
			User:       fromUser.User,
			Immune:     fromUser.Immune,
			ImmuneTs:   fromUser.ImmuneTs,
		},
	)
}
