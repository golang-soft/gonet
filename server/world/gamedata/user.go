package gamedata

import (
	"fmt"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/world/cache"
	"gonet/server/world/helper"
	"reflect"
	"time"
)

func getDistance(round int, from string, to string) float64 {
	posFrom := PositionCtrl.GetPos(round, from)
	posTo := PositionCtrl.GetPos(round, to)
	posNowFrom := PositionCtrl.calPos(posFrom)
	posNowTo := PositionCtrl.calPos(posTo)

	return PositionCtrl.distance(*posNowFrom, *posNowTo)
}

// export async function makeDizzy(from: string, to: string, skill: number, skillId: number, ts: number) {
//     broadcastAll(USER_EVENT.USER.ATTACK_DIZZY, { from, to, skill, skillId, duration: ts, ts: Date.now() })
// }

type (
	SUserCtrl struct {
	}
	ISUserCtrl interface {
	}
)

var UserCtrl *SUserCtrl = &SUserCtrl{}

//func (this* SUserCtrl)canAttack(from string, round int, to string, skill int32, skillId int32, x float64, y float64 , z float64, msg string) bool {
func (this *SUserCtrl) canAttack(message data.AttackData) bool {
	from, round, to, skill, skillId, x, y, z, msg := message.From, message.Round, message.To, message.Skill, message.SkillId, message.X, message.Y, message.Z, message.Msg
	skillCfg := (*helper.SKILL_BASIC_INFO)[skillId]

	if skillCfg == nil {
		return false
	}
	//TODO: 判断游戏
	formUserAttr := GGame.GetUserById(round, from)
	// hp
	if (formUserAttr.Hp) <= 0 {
		return false
	}
	//判断cd
	cd_key := fmt.Sprintf("skill_%d_cd", skillId)
	field := reflect.ValueOf(formUserAttr).Elem().FieldByName(cd_key)

	if time.Now().Unix()-field.Int() < int64(skillCfg.CD)*1000 {
		return false
	}

	// 攻击目标
	if skillCfg.Need_Target == 1 {
		//  需要目标
		toUser := GGame.GetUserById(round, to)
		//同阵营不做攻击
		if toUser.Part == formUserAttr.Part {
			return false
		}
		// hp
		if toUser.Hp <= 0 {
			return false
		}
		// 距离
		distance := getDistance(round, from, to)
		// 有攻击目标,只需要判断攻击的距离是否满足即可
		if distance > skillCfg.Range {
			// await broadcastToSelf(USER_EVENT.USER.OUT_RANGE, from, {})
			// gvgBattleBroadcastAll(USER_EVENT.USER.OUT_RANGE, round, {
			//     from: from,
			//     to: to,
			//     skill: skill,
			//     skillId: skillId,
			//     x: x,
			//     y: y,
			//     z: z,
			//     msg: msg
			// })
			return false
		}
	}
	// 更新技能冷却时间
	GGame.updateSkillCD(round, from, skillId, time.Now().Unix())

	GvgBattleBroadcastAll("ATTACK_SUCCESS", round,
		&cmessage.AttackSuccessResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_AttackSuccessResp), 0),
			From:       from,
			To:         to,
			Skill:      skill,
			SkillId:    skillId,
			X:          x,
			Y:          y,
			Z:          z,
			Msg:        msg,
		},
	)
	return true
}

func (this *SUserCtrl) attack(attack data.AttackData) {
	// const keyFrom = REDIS_KEYS.user_round_basic + getRoundKey(attack.from, attack.round)
	fromUser := GGame.GetUserById(attack.Round, attack.From)
	itype := fromUser.Itype
	if itype == common.Role.Hunter {
		//猎人
		v := reflect.ValueOf(Hunter)
		method := fmt.Sprintf("skill_%d", attack.SkillId)
		m := v.MethodByName(method)

		m.Call([]reflect.Value{reflect.ValueOf(attack)})
	} else if itype == common.Role.Wizard {
		//法师
		v := reflect.ValueOf(Wizard)
		method := fmt.Sprintf("skill_%d", attack.SkillId)
		m := v.MethodByName(method)

		m.Call([]reflect.Value{reflect.ValueOf(attack)})

	} else if itype == common.Role.Warrior {
		//骑士
		v := reflect.ValueOf(Warrior)
		method := fmt.Sprintf("skill_%d", attack.SkillId)
		m := v.MethodByName(method)

		m.Call([]reflect.Value{reflect.ValueOf(attack)})
	} else if itype == common.Role.Adventurer {
		//炮手
		v := reflect.ValueOf(Adventurer)
		method := fmt.Sprintf("skill_%d", attack.SkillId)
		m := v.MethodByName(method)

		m.Call([]reflect.Value{reflect.ValueOf(attack)})
	} else if itype == common.Role.Rogue {
		//战士
		v := reflect.ValueOf(Rogue)
		method := fmt.Sprintf("skill_%d", attack.SkillId)
		m := v.MethodByName(method)

		m.Call([]reflect.Value{reflect.ValueOf(attack)})
	} else {
		return
	}
}

//TODO:复活
func (this *SUserCtrl) relivePlayer(round int, from string, part int) {
	userInfo := GGame.GetUserById(round, from)
	if userInfo.Hp > 0 {
		return
	}
	GameCtrl.relivePlayer(round, from, userInfo.Part)
}

//占旗
func (this *SUserCtrl) flag(body *data.FlagReqData) {
	FlagCtrl.Flag(
		body.Round,
		body.From,
		body.Part,
	)
}

//道具
func (this *SUserCtrl) item(body *data.ItemData) {
	ItemCtrl.useitem(body)
}

func (this *SUserCtrl) createPlayer(user string) *cache.PlayerData {
	return cache.GameCache.CeratePlayer(user)
}

func (this *SUserCtrl) getUser(user string) {
	userData := cache.GameCache.GetPlayer(user)
	if userData == nil {
		//TODO:用户不合法 临时创建
		userData = this.createPlayer(user)
	}
}

//夺旗成功
func (this *SUserCtrl) playerLogin(user string) {
	//userInfo := this.getUser(user)

}
