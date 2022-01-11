package datafnc

import (
	"fmt"
	"gonet/server/common"
	"gonet/server/common/data"
	table2 "gonet/server/table"
	"gonet/server/world/logger"
	"gonet/server/world/table"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var (
	Equip_Idx                      []int
	Hero_Role_Type                 []int
	All_Attr_Percent               float64       //全属性百分比
	Skill_1402_Buff_AddAtk_Percent float64       //炮手攻击加成百分比
	Max_Kill                       int           //最大击杀奖励次数
	Max_Death                      int           //最大死亡扣除次数
	Kill_Reward_Devt               int           //单次击杀奖励dvt
	Second_Score                   float64       //每秒积分
	Battle_Time                    int           //战场时间
	Max_Score                      float64       //最高分
	Score_Time_Interval            int64         //夺旗验证时间
	Extra_Time_Reward              map[int]int   //额外奖励
	Battle_Victory_Pieces_Reward   []*RewardData //胜利碎片奖励
	Battle_Defeated_Pieces_Reward  []*RewardData //失败碎片奖励
	Battle_Victory_Dvt_Reward      int           //战场胜利抽成dvt
	Resurrection_Time              int           //复活时间
	Resurrection_Manual_Time       int           //死亡后操作时间（复活道具使用限时）
	Entry_Fee                      int           //入场费用（门票+消费dvt）
	Camp_Player_Amount             int           //单阵营人数
	Room_Player_Max_Size           int           //gvg战场房间最大人数
	Rank_Update_Cycle              float64       //排行榜周期
	Leaderboard_Prefix             float64       //排行榜前这
	Battle_Mod                     float64       //"攻防公式算法，1为减法公式，2为除法公式"
	Battle_Defence                 float64       // "除法公式下的防御参数，D值"
)

func Init() {
	Equip_Idx = getEquipIdx()
	Hero_Role_Type = getHeroRoles()
	All_Attr_Percent = getAllAttrPercent()                    //全属性百分比
	Skill_1402_Buff_AddAtk_Percent = skill1402AddAtk()        //炮手攻击加成百分比
	Max_Kill = getMaxKill()                                   //最大击杀奖励次数
	Max_Death = getMaxDeath()                                 //最大死亡扣除次数
	Kill_Reward_Devt = getKillRewardDevt()                    //单次击杀奖励dvt
	Second_Score = getSecScore()                              //每秒积分
	Battle_Time = getBattleEndTime()                          //战场时间
	Max_Score = getMaxScore()                                 //最高分
	Score_Time_Interval = getFlagCaptureTimes()               //夺旗验证时间
	Extra_Time_Reward = getExtraTimeRewardScore()             //额外奖励
	Battle_Victory_Pieces_Reward = getVictoryPiecesReward()   //胜利碎片奖励
	Battle_Defeated_Pieces_Reward = getDefeatedPiecesReward() //失败碎片奖励
	Battle_Victory_Dvt_Reward = getVictoryDvtReward()         //战场胜利抽成dvt
	Resurrection_Time = getResurrectionTime()                 //复活时间
	Resurrection_Manual_Time = getResurrectionManualTime()    //死亡后操作时间（复活道具使用限时）
	Entry_Fee = getEntryFee()                                 //入场费用（门票+消费dvt）
	Camp_Player_Amount = CampPlayerAmount()                   //单阵营人数
	Room_Player_Max_Size = RoomPlayerMaxSize()                //gvg战场房间最大人数
	Rank_Update_Cycle = RankUpdateCycle()                     //排行榜周期
	Leaderboard_Prefix = GetLeaderboardPrefix()               //排行榜前这
	Battle_Mod = BattleMod()                                  //"攻防公式算法，1为减法公式，2为除法公式"
	Battle_Defence = BattleDefence()                          // "除法公式下的防御参数，D值"
}

//夺旗间隔时间
func getAllAttrPercent() float64 {
	return float64((*table.ITEM_CONFIG)[int32(common.Item_allAttr)].Attribute[1] / 100)
}

//夺旗间隔时间
func getFlagCaptureTimes() int64 {
	return int64(getObjectValue("OccupyTime") * common.TIME.Second)
}

//单场战斗最大分， 结算分数
func getMaxScore() float64 {
	return float64(getObjectValue("Integral"))
}

//每秒获得积分
func getSecScore() float64 {
	return float64(getObjectValue("OccupyIntegralSec"))
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
	var rewardList map[int]int = make(map[int]int)
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
		logger.M_Log.Debugf("found key from WORLD_CONFIG faild")
	}
	return val
}

type RewardData struct {
	Id    int
	Count int
}

//胜方碎片奖励
func getVictoryPiecesReward() []*RewardData {
	var rewardList []*RewardData
	data := RewardData{Id: common.Item_box_piece, Count: getObjectValue("WinnerChipReward")}
	rewardList = append(rewardList, &data)
	return rewardList
}

//败方碎片奖励
func getDefeatedPiecesReward() []*RewardData {
	var rewardList []*RewardData
	data := RewardData{Id: common.Item_box_piece, Count: getObjectValue("LoserChipReward")}
	rewardList = append(rewardList, &data)
	return rewardList
}

//胜方抽成
func getVictoryDvtReward() int {
	var total int = 0
	total = getObjectValue("PlayerAmount") * getObjectValue("AdmissionTicket") * getObjectValue("WinnerRewardRatio")

	return total
}

//战场道具cd
func GetBattleItemsById(battle_id int) map[string]int {
	var itemAttr map[string]int
	for _, config := range *table.ITEM_CONFIG {
		for _, id := range config.Battle_id {
			if id == battle_id {
				key := fmt.Sprintf("item_%d_cd", config.ID)
				itemAttr[key] = 0
			}
		}
	}
	return itemAttr
}

//战场道具id枚举 【111，3333，4444】
func getBattleItemIdById(battle_id int) []int {
	var itemIdx []int
	for _, config := range *table.ITEM_CONFIG {
		for _, id := range config.Battle_id {
			if id == battle_id {
				itemIdx = append(itemIdx, config.ID)
			}
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

func RankUpdateCycle() float64 {
	return float64(getObjectValue("RankUpdateCycle") / 86400)
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

func skill1402AddAtk() float64 {
	var data = (*table.SKILL_BASIC_INFO)[(common.Skill.Skill_1402)]
	if &data == nil {
		//world.SERVER.M_Log.Debugf("error skill config")
	}

	return (*table.SKILL_BASIC_INFO)[(common.Skill.Skill_1402)].Buff[1] / 100
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

func GetEquipCfg(id int32) *table2.ItemTableCfg {
	return (*table.ITEM_CONFIG)[id]
}

func getServerStartTime() int64 {
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-12-30 10:00:00", local)
	fmt.Println(t)
	return t.Unix()
}

func GetLeaderboardPrefix() float64 {
	server_start := getServerStartTime()
	timeStamp1 := math.Floor(float64(server_start / 1000))
	timeStamp2 := math.Floor(float64(time.Now().Unix() / 1000))
	day1 := math.Floor(timeStamp1 / 86400)
	day2 := math.Floor(timeStamp2 / 86400)

	return math.Ceil((day2 - day1 + 0.5) / Rank_Update_Cycle)
}
func BattleMod() float64 {
	return float64(getObjectValue("BattleMod"))
}

func BattleDefence() float64 {
	return float64(getObjectValue("BattleDefence"))
}

func GetBornPoint(part int32, mapId int) *data.Pos {
	var bornX float64 = 0
	var bornY float64 = 0
	var minX int = 0
	var maxX int = 0
	var minY int = 0
	var maxY int = 0
	MapCfg := table.MAP_CONFIG
	mapData := (*MapCfg)[mapId]

	partInfo := (*mapData).Red
	if part == 2 {
		partInfo = (*mapData).Blue
	}
	minX = partInfo[0] - partInfo[2]
	maxX = partInfo[0] + partInfo[2]
	minY = partInfo[1] - partInfo[2]
	maxY = partInfo[1] + partInfo[2]

	bornX = float64(rand.Intn(maxX-minX) + minX)
	bornY = float64(rand.Intn(maxY-minY) + minY)
	return &data.Pos{X: bornX, Y: bornY}
}
