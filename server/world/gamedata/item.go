package gamedata

import (
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/world/datafnc"
	"gonet/server/world/helper"
	"gonet/server/world/table"
	"math"
)

//道具
type SItemCtrl struct {
}

var ItemCtrl *SItemCtrl = &SItemCtrl{}

func (this *SItemCtrl) useitem(body *data.ItemData) {
	var itemId int32
	itemCfg := (*table.ITEM_CONFIG)[itemId]

	fromUser := GGame.GetUserById(body.Round, body.From)
	if fromUser == nil {
		return
	}

	//itemCd := fmt.Sprintf("item_%d_cd", body.ItemId)

	play := helper.USER_BASIC_INFO[fromUser.Itype]
	switch itemCfg.Attribute[0] {
	//道具生效类型
	case common.ItemAttr.Type_1:
		//  A=1，增加自身全属性 （1，增加的属性百分比）
		if fromUser.Hp <= 0 || fromUser.AllAttr == 1 {
			return
		}
		GameCtrl.addAllAttr(body.Round, body.From, datafnc.All_Attr_Percent, play, fromUser)
		break
	case common.ItemAttr.Type_2:
		/*A=2，增加自身移速   （2，增加的移速百分比）*/
		break
	case common.ItemAttr.Type_3:
		/*A=3，隐藏自己模型   （3，无值填0）*/
		break
	case common.ItemAttr.Type_4:
		/*A=4，使自己无敌，免疫伤害和控制，可以攻击和使用技能、道具，但无法移动*/
		break
	case common.ItemAttr.Type_5:
		/*A=5，使用后自己的位置不会出现在对方的小地图雷达上*/
		break
	case common.ItemAttr.Type_6:
		/*A=6，立即减少当前复活时间*/
		if fromUser.Hp > 0 {
			return
		}
		//复活
		GameCtrl.relivePlayer(body.Round, body.From, fromUser.Part)
		break
	case common.ItemAttr.Type_7:
		/*A=7，恢复血量       （7，恢复的血量百分比）*/
		//满血不使用
		if play.Hp == fromUser.Hp {
			return
		}
		//判断血量死亡
		if fromUser.Hp <= 0 || fromUser.DieTs > 0 {
			return
		}
		//道具加血 武器属性增益
		maxHp := play.Hp
		var percent float64 = 0

		maxHp += float64(fromUser.Equip_3)
		if fromUser.AllAttr == 1 {
			percent += datafnc.All_Attr_Percent
		}
		percent += fromUser.Equip_9 / 100

		userHp := math.Floor(float64(maxHp)*1 + float64(percent))
		addHp := math.Floor(float64(userHp) * float64(itemCfg.Attribute[1]/100))
		realHp := addHp
		if float64(userHp)-float64(fromUser.Hp) < addHp {
			realHp = userHp - float64(fromUser.Hp)
		}
		GameCtrl.addHp(body.Round, body.From, realHp)
		//TODO:道具cd
		break
	case common.ItemAttr.Type_8:
		/*A=8，增加自己成为英雄的概*/
		break
	case common.ItemAttr.Type_9:
		/*A=9，立即按固定轨迹进行飞行（9，轨迹编号）*/
		break
	case common.ItemAttr.Type_10:
		break
	default:
		//console.error("item error");

	}
	//道具数量扣除
	// await GameCtrl.useItem(round, from, itemId, 1)
	GvgBattleBroadcastAll("USE_ITEM", body.Round,
		&cmessage.UseItemResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_UseItemResp), 0),
			From:       body.From,
			ItemId:     itemId,
			Count:      0,
			Hp:         fromUser.Hp,
			Msg:        "",
		},
	)
}
