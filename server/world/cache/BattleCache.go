package cache

import (
	"encoding/json"
	"fmt"
	"gonet/server/common"
	"gonet/server/world"
	"gonet/server/worlddb/mredis"
	"strconv"
	"time"
)

type (
	BattleCache struct {
	}

	IBattleCache interface {
	}
)

var (
	MBattleCache *BattleCache = NewBattleCache()
)

func NewBattleCache() *BattleCache {
	return &BattleCache{}
}

//加血
func (this *BattleCache) addHp(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "hp", num)
}

//加全属性
func (this *BattleCache) addAllAttr(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "hp", num)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "allAttr", 1)
}

//减血
func (this *BattleCache) desHp(round int, id string, num float64, dieTs float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	if dieTs > 0 {
		var hmsetData map[string]interface{}
		hmsetData["dieTs"] = dieTs
		world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
		world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "hp", -num)
	} else {
		world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "hp", -num)
	}
}

//dev易手
func (this *BattleCache) updataDvt(round int, from string, to string, dvt float64) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(fromkey, "getDvt", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(fromkey, "dvt", dvt)
	world.SERVER.M_pRedisClient.HIncrByFloat(tokey, "getDvt", -1)
	world.SERVER.M_pRedisClient.HIncrByFloat(tokey, "dvt", -dvt)
}

//更新击杀数和死亡数
func (this *BattleCache) updataKillAndDie(round int, from string, to string) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(fromkey, "kill", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(tokey, "die", 1)
}

func (this *BattleCache) getDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "getDvt", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "dvt", num)
}

func (this *BattleCache) desDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "desDvt", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "dvt", -num)
}

//================================================================================================
//眩晕
func (this *BattleCache) addDizzy(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 1
	hmsetData["dizzyTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除眩晕
func (this *BattleCache) removeDizzy(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 0
	hmsetData["dizzyTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//禁锢
func (this *BattleCache) addStopMove(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除禁锢
func (this *BattleCache) removeStopMove(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 0
	hmsetData["stopmoveTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//变身
func (this *BattleCache) addDeformation(round int, id string, addAtk float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除变身
func (this *BattleCache) removeDeformation(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 0
	hmsetData["stopmoveTs"] = 0
	hmsetData["addAtk"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//减速
func (this *BattleCache) desSpeed(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["reduceSpeedTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//恢复速度
func (this *BattleCache) recoverSpeed(round int, id string, speed float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["speed"] = speed
	hmsetData["reduceSpeedTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//加盾
func (this *BattleCache) addShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["shield"] = num
	hmsetData["shieldTs"] = timed
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减盾
func (this *BattleCache) desShield(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//移除盾
func (this *BattleCache) removeShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//================================================================================================
//无敌
func (this *BattleCache) addImmune(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = num
	hmsetData["immuneTs"] = timed
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减无敌
func (this *BattleCache) desImmune(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除无敌
func (this *BattleCache) removeImmune(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//反伤
func (this *BattleCache) addThorns(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["thorns"] = num
	hmsetData["thornsTs"] = timed
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减反伤
func (this *BattleCache) desThorns(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["thorns"] = 0
	hmsetData["thornsTs"] = 0
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除反伤
func (this *BattleCache) removeThorns(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["thorns"] = 0
	hmsetData["thornsTs"] = 0
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//===============================================
//输出伤害
func (this *BattleCache) addDps(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "dps", num)
}
func (this *BattleCache) updateSkillCd(round int, id string, skillId int, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	skill := fmt.Sprintf("skill_%d_cd", skillId)
	var hmsetData map[string]interface{}
	hmsetData[skill] = timed
	hmsetData["updateTs"] = timed
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//新建回合玩家数据
func (this *BattleCache) newRoundGameData(round int, id string, game map[string]interface{}) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	world.SERVER.M_pRedisClient.HMSet(keyTo, game)
}

//新建回合用户
func (this *BattleCache) newRoundUserIdx(round int, users map[string]interface{}) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_players] + strconv.Itoa(round)
	world.SERVER.M_pRedisClient.SADD(keyTo, users)
}

//新建回合玩家数据
func (this *BattleCache) newRoundPlayer(round int, id string, player map[string]interface{}) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HMSet(keyTo, player)
}

func (this *BattleCache) newRoundPlayerEquip(round int, id string, equip map[string]interface{}) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	jsonstr, _ := json.Marshal(equip)
	world.SERVER.M_pRedisClient.HSet(keyTo, "equip", string(jsonstr))
}

//道具 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func (this *BattleCache) getPlayerKey(table string, uuid string) string {
	return fmt.Sprintf("%s:%s", table, uuid)
}

func (this *BattleCache) getHeroData(user string, hid string) map[string]interface{} {
	redis_key := this.getPlayerKey(common.Tables_hero, user)
	heroDataStr := world.SERVER.M_pRedisClient.HGet(redis_key, hid)
	var hero map[string]interface{}
	err := json.Unmarshal(([]byte)(heroDataStr.Val()), hero)
	if err != nil {
		world.SERVER.M_Log.Debugf("转换出错误")
	}
	return hero
}

func (this *BattleCache) getItems(user string) map[int]interface{} {
	//redis_key := common.Redis_Prefix_Item + ":" + user
	item := make(map[int]interface{})

	return item
}

type ItemData struct {
	id    int
	count float64
}

func (this *BattleCache) updateItemData(round int, user string, item map[int]ItemData) {

	redis_key := common.Redis_Prefix_Item + ":" + user
	key := common.GetRoundKey(user, round)
	for _, data := range item {
		world.SERVER.M_pRedisClient.HIncrByFloat(redis_key, strconv.Itoa(data.id), data.count)
		world.SERVER.M_pRedisClient.HIncrByFloat(key, fmt.Sprintf("item_%d_cd", data.id), float64(time.Now().Unix()))
	}
}

//更新旗帜归属
func (this *BattleCache) updataFlagOwner(round int, user string, part int, partTs float64, flagTs float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	partKey := fmt.Sprintf("part%sTs", part)
	world.SERVER.M_pRedisClient.HSet(keyTo, "flagUpdateTs", flagTs)
	world.SERVER.M_pRedisClient.HSet(keyTo, "flagUser", user)
	world.SERVER.M_pRedisClient.HSet(keyTo, partKey, partTs)
}

//更新旗帜归属
func (this *BattleCache) updataFlagPart(round int, part int) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	partKey := fmt.Sprintf("part%sTs", part)
	world.SERVER.M_pRedisClient.HSet(keyTo, "flagUser", "")
	world.SERVER.M_pRedisClient.HSet(keyTo, partKey, 0)
}

//更新队伍积分
func (this *BattleCache) updataPartScore(round int, part float64, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	key := fmt.Sprintf("part%sscore", part)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, key, num)
}

//更新队伍积分
func (this *BattleCache) updataFlagTs(round int, part int, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	world.SERVER.M_pRedisClient.HSet(keyTo, "flagUpdateTs", timed)
	world.SERVER.M_pRedisClient.HSet(keyTo, "flagOwner", part)
}

func (this *BattleCache) battleEnd(round int, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	world.SERVER.M_pRedisClient.HSet(keyTo, "endTs", timed)
}
