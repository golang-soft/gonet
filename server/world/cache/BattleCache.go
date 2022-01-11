package cache

import (
	"encoding/json"
	"fmt"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/common/mredis"
	"gonet/server/world/redisInstnace"
	"strconv"
	"time"
)

type (
	SBattleCache struct {
	}

	ISBattleCache interface {
	}
)

var (
	BattleCache *SBattleCache = NewBattleCache()
)

func NewBattleCache() *SBattleCache {
	return &SBattleCache{}
}

//加血
func (this *SBattleCache) AddHp(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "hp", float64(num))
}

//加全属性
func (this *SBattleCache) AddAllAttr(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "hp", num)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "allAttr", 1)
}

//减血
func (this *SBattleCache) DesHp(round int, id string, num float64, dieTs float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	if dieTs > 0 {
		var hmsetData map[string]interface{}
		hmsetData["dieTs"] = dieTs
		redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
		redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "hp", -num)
	} else {
		redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "hp", -num)
	}
}

//dev易手
func (this *SBattleCache) UpdataDvt(round int, from string, to string, dvt float64) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(fromkey, "getDvt", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(fromkey, "dvt", dvt)
	redisInstnace.M_pRedisClient.HIncrByFloat(tokey, "getDvt", -1)
	redisInstnace.M_pRedisClient.HIncrByFloat(tokey, "dvt", -dvt)
}

//更新击杀数和死亡数
func (this *SBattleCache) UpdataKillAndDie(round int, from string, to string) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(fromkey, "kill", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(tokey, "die", 1)
}

func (this *SBattleCache) getDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "getDvt", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "dvt", num)
}

func (this *SBattleCache) desDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "desDvt", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "dvt", -num)
}

//================================================================================================
//眩晕
func (this *SBattleCache) AddDizzy(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 1
	hmsetData["dizzyTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除眩晕
func (this *SBattleCache) RemoveDizzy(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 0
	hmsetData["dizzyTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//禁锢
func (this *SBattleCache) AddStopMove(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除禁锢
func (this *SBattleCache) RemoveStopMove(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 0
	hmsetData["stopmoveTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//变身
func (this *SBattleCache) AddDeformation(round int, id string, addAtk float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除变身
func (this *SBattleCache) RemoveDeformation(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 0
	hmsetData["stopmoveTs"] = 0
	hmsetData["addAtk"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//减速
func (this *SBattleCache) DesSpeed(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["reduceSpeedTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//恢复速度
func (this *SBattleCache) RecoverSpeed(round int, id string, speed float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["speed"] = speed
	hmsetData["reduceSpeedTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//加盾
func (this *SBattleCache) AddShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["shield"] = num
	hmsetData["shieldTs"] = timed
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减盾
func (this *SBattleCache) DesShield(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//移除盾
func (this *SBattleCache) removeShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//================================================================================================
//无敌
func (this *SBattleCache) AddImmune(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = num
	hmsetData["immuneTs"] = timed
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减无敌
func (this *SBattleCache) desImmune(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除无敌
func (this *SBattleCache) RemoveImmune(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//反伤
func (this *SBattleCache) AddThorns(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["thorns"] = num
	hmsetData["thornsTs"] = timed
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减反伤
func (this *SBattleCache) desThorns(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["thorns"] = 0
	hmsetData["thornsTs"] = 0
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除反伤
func (this *SBattleCache) RemoveThorns(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["thorns"] = 0
	hmsetData["thornsTs"] = 0
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//===============================================
//输出伤害
func (this *SBattleCache) AddDps(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "dps", num)
}

func (this *SBattleCache) UpdateSkillCd(round int, id string, skillId int32, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	skill := fmt.Sprintf("skill_%d_cd", skillId)
	var hmsetData map[string]interface{}
	hmsetData[skill] = timed
	hmsetData["updateTs"] = timed
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//新建回合玩家数据
func (this *SBattleCache) NewRoundGameData(round int, id string, game map[string]interface{}) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	redisInstnace.M_pRedisClient.HMSet(keyTo, game)
}

//新建回合用户
func (this *SBattleCache) NewRoundUserIdx(round int, users []string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_players] + strconv.Itoa(round)
	redisInstnace.M_pRedisClient.SADD(keyTo, users)
}

//新建回合玩家数据
func (this *SBattleCache) NewRoundPlayer(round int, id string, player map[string]interface{}) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HMSet(keyTo, player)
}

func (this *SBattleCache) NewRoundPlayerEquip(round int, id string, equip map[int]int) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	jsonstr, _ := json.Marshal(equip)
	redisInstnace.M_pRedisClient.HSet(keyTo, "equip", string(jsonstr))
}

//道具 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func (this *SBattleCache) getPlayerKey(table string, uuid string) string {
	return fmt.Sprintf("%s:%s", table, uuid)
}

func (this *SBattleCache) getHeroData(user string, hid string) map[string]interface{} {
	redis_key := this.getPlayerKey(common.Tables_hero, user)
	heroDataStr := redisInstnace.M_pRedisClient.HGet(redis_key, hid)
	var hero map[string]interface{}
	err := json.Unmarshal(([]byte)(heroDataStr.Val()), hero)
	if err != nil {
		//world.SERVER.M_Log.Debugf("转换出错误")
	}
	return hero
}

func (this *SBattleCache) GetItems(user string) map[int]data.ItemData {
	//redis_key := common.Redis_Prefix_Item + ":" + user
	item := make(map[int]data.ItemData)

	return item
}

//type ItemData struct {
//	id    int
//	count float64
//}

func (this *SBattleCache) UpdateItemData(round int, user string, item map[int]data.ItemData) {

	redis_key := common.Redis_Prefix_Item + ":" + user
	key := common.GetRoundKey(user, round)
	for _, data := range item {
		redisInstnace.M_pRedisClient.HIncrByFloat(redis_key, strconv.Itoa(data.ItemId), float64(data.Count))
		redisInstnace.M_pRedisClient.HIncrByFloat(key, fmt.Sprintf("item_%d_cd", data.ItemId), float64(time.Now().Unix()))
	}
}

//更新旗帜归属
func (this *SBattleCache) UpdataFlagOwner(round int, user string, part int32, partTs float64, flagTs float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	partKey := fmt.Sprintf("part%sTs", part)
	redisInstnace.M_pRedisClient.HSet(keyTo, "flagUpdateTs", flagTs)
	redisInstnace.M_pRedisClient.HSet(keyTo, "flagUser", user)
	redisInstnace.M_pRedisClient.HSet(keyTo, partKey, partTs)
}

//更新旗帜归属
func (this *SBattleCache) UpdataFlagPart(round int, part int32) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	partKey := fmt.Sprintf("part%sTs", part)
	redisInstnace.M_pRedisClient.HSet(keyTo, "flagUser", "")
	redisInstnace.M_pRedisClient.HSet(keyTo, partKey, 0)
}

//更新队伍积分
func (this *SBattleCache) UpdataPartScore(round int, part int32, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	key := fmt.Sprintf("part%sscore", part)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, key, num)
}

//更新队伍积分
func (this *SBattleCache) UpdataFlagTs(round int, part int32, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	redisInstnace.M_pRedisClient.HSet(keyTo, "flagUpdateTs", timed)
	redisInstnace.M_pRedisClient.HSet(keyTo, "flagOwner", part)
}

func (this *SBattleCache) BattleEnd(round int, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	redisInstnace.M_pRedisClient.HSet(keyTo, "endTs", timed)
}
