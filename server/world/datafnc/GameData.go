package datafnc

import (
	"fmt"
	"gonet/server/common"
	"gonet/server/world"
	"gonet/server/world/table"
	"strconv"
)

var (
	Equip_Idx                      = getEquipIdx()
	Hero_Role_Type                 = getHeroRoles()
	All_Attr_Percent               = getAllAttrPercent()         //全属性百分比
	Skill_1402_Buff_AddAtk_Percent = skill1402AddAtk()           //炮手攻击加成百分比
	Max_Kill                       = getMaxKill()                //最大击杀奖励次数
	Max_Death                      = getMaxDeath()               //最大死亡扣除次数
	Kill_Reward_Devt               = getKillRewardDevt()         //单次击杀奖励dvt
	Second_Score                   = getSecScore()               //每秒积分
	Battle_Time                    = getBattleEndTime()          //战场时间
	Max_Score                      = getMaxScore()               //最高分
	Score_Time_Interval            = getFlagCaptureTimes()       //夺旗验证时间
	Extra_Time_Reward              = getExtraTimeRewardScore()   //额外奖励
	Battle_Victory_Pieces_Reward   = getVictoryPiecesReward()    //胜利碎片奖励
	Battle_Defeated_Pieces_Reward  = getDefeatedPiecesReward()   //失败碎片奖励
	Battle_Victory_Dvt_Reward      = getVictoryDvtReward()       //战场胜利抽成dvt
	Resurrection_Time              = getResurrectionTime()       //复活时间
	Resurrection_Manual_Time       = getResurrectionManualTime() //死亡后操作时间（复活道具使用限时）
	Entry_Fee                      = getEntryFee()               //入场费用（门票+消费dvt）
	Camp_Player_Amount             = CampPlayerAmount()          //单阵营人数
	Room_Player_Max_Size           = RoomPlayerMaxSize()         //gvg战场房间最大人数
)

//夺旗间隔时间
func getAllAttrPercent() int {
	return (*table.ITEM_CONFIG)[strconv.Itoa(common.Item_allAttr)].Attribute[1] / 100
}

//夺旗间隔时间
func getFlagCaptureTimes() int {
	return getObjectValue("OccupyTime") * common.TIME.Second
}

//单场战斗最大分， 结算分数
func getMaxScore() int {
	return getObjectValue("Integral")
}

//每秒获得积分
func getSecScore() int {
	return getObjectValue("OccupyIntegralSec")
}

//单场战斗持续时间
func getBattleEndTime() int {
	return getObjectValue("BattlefieldTime") * common.TIME.Second
}

//单场战斗最大获取dvt次数
func getKillRewardDevt() int {
	return getObjectValue("KillAward")
}

//单场战斗最小获取dvt次数
func getMaxKill() int {
	return getObjectValue("MaxKillReward")
}

//单场战斗最大丢失dvt数量
func getMaxDeath() int {
	return getObjectValue("MaxDeathReward") * -1
}

//战斗阶段分数奖励
func getExtraTimeRewardScore() map[int]int {
	var rewardList map[int]int
	rewardList[common.Battle.Time_3] = getObjectValue("OccupyIntegralThree")
	rewardList[common.Battle.Time_5] = getObjectValue("OccupyIntegralFive")
	rewardList[common.Battle.Time_10] = getObjectValue("OccupyIntegralTen")
	rewardList[common.Battle.Time_15] = getObjectValue("OccupyIntegralFifteen")

	return rewardList
}

func getObjectValue(key string) int {
	var val int = 0
	var hasKey bool = false
	for _, config := range *table.WORLD_CONFIG {
		if config.Key == key {
			val, _ = strconv.Atoi(config.Value)
			hasKey = true
			break
		}
	}

	if !hasKey {
		world.SERVER.M_Log.Debugf("found key from WORLD_CONFIG faild")
	}
	return val
}

type RewardData struct {
	id    int
	count int
}

//胜方碎片奖励
func getVictoryPiecesReward() []RewardData {
	var rewardList []RewardData
	data := RewardData{id: common.Item_box_piece, count: getObjectValue("WinnerChipReward")}
	rewardList = append(rewardList, data)
	return rewardList
}

//败方碎片奖励
func getDefeatedPiecesReward() []RewardData {
	var rewardList []RewardData
	data := RewardData{id: common.Item_box_piece, count: getObjectValue("LoserChipReward")}
	rewardList = append(rewardList, data)
	return rewardList
}

//胜方抽成
func getVictoryDvtReward() int {
	var total int = 0
	total = getObjectValue("PlayerAmount") * getObjectValue("AdmissionTicket") * getObjectValue("WinnerRewardRatio")

	return total
}

//战场道具cd
func getBattleItemsById(battle_id int) map[string]int {
	var itemAttr map[string]int
	for _, config := range *table.ITEM_CONFIG {
		if config.Battle_id == battle_id {
			key := fmt.Sprintf("item_%d_cd", config.ID)
			itemAttr[key] = 0
		}
	}
	return itemAttr
}

//战场道具id枚举 【111，3333，4444】
func getBattleItemIdById(battle_id int) []int {
	var itemIdx []int
	for _, config := range *table.ITEM_CONFIG {
		if config.Battle_id == battle_id {
			itemIdx = append(itemIdx, config.ID)
		}
	}
	return itemIdx
}

//复活时间
func getResurrectionTime() int {
	return getObjectValue("ResurrectionTime")
}

//复活时间
func getResurrectionManualTime() int {
	return getObjectValue("ResurrectionManualTime")
}

//入场标准（包含门票，消费dvt）
func getEntryFee() int {
	return getObjectValue("EntryFee")
}

func CampPlayerAmount() int {
	return getObjectValue("CampPlayerAmount")
}

func RoomPlayerMaxSize() int {
	return getObjectValue("PlayerAmount")
}

func skill1402AddAtk() int {
	var data = (*table.SKILL_BASIC_INFO)[strconv.Itoa(common.Skill.Skill_1402)]
	if &data == nil {
		world.SERVER.M_Log.Debugf("error skill config")
	}

	return (*table.SKILL_BASIC_INFO)[strconv.Itoa(common.Skill.Skill_1402)].Buff[1] / 100
}

func getEquipIdx() []int {
	var idx []int
	for _, config := range *table.ITEM_CONFIG {
		if config.Type == common.Item_equip {
			idx = append(idx, config.ID)
		}
	}
	return idx
}

func getHeroRoles() []int {
	var role []int
	for _, config := range *table.RAW_HERO_ATTR {
		if config.Class == common.Item_equip {
			role = append(role, config.Class)
		}
	}
	return role
}
