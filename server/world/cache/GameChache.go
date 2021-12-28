package cache

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gonet/server/common"
	"gonet/server/world"
	"gonet/server/world/datafnc"
	"gonet/server/worlddb/mredis"
	"strconv"
	"time"
)

type (
	GameCache struct {
	}

	IGameCache interface {
	}
)

func (this *GameCache) getAccountkey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_user, uuid)
}

func (this *GameCache) getItemkey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_item, uuid)
}
func (this *GameCache) getHerokey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_hero, uuid)
}
func (this *GameCache) getEquipkey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_equip, uuid)
}

type PlayerInfo struct {
	id       string
	nickName string
	email    string
	dvt      int
	createTs int
	updateTs int
}

type ItemInfo struct {
	id    int
	count int
}

type HeroInfo struct {
	uuid     int
	heroType int
	equips   []string
}

type EquipInfo struct {
	uuid   int
	itemId int
	heroId string
}

type PlayerData struct {
	player PlayerInfo
	item   []ItemInfo
	hero   []HeroInfo
	equip  []EquipInfo
}

func (this *GameCache) ceratePlayer(uuid string) *PlayerData {
	playerKeyStr := this.getAccountkey(uuid)
	itemKeyStr := this.getItemkey(uuid)
	heroKeyStr := this.getHerokey(uuid)
	equipKeyStr := this.getEquipkey(uuid)

	accountData := this.initPlayer()
	accountData.id = uuid
	itemData := this.initItem()
	heroData := this.initHero()
	equipData := this.initEquip(uuid)

	accountDataJson, _ := json.Marshal(accountData)
	itemDataJson, _ := json.Marshal(itemData)
	world.SERVER.M_pRedisClient.Set(playerKeyStr, accountDataJson, -1)
	world.SERVER.M_pRedisClient.Set(itemKeyStr, itemDataJson, -1)
	world.SERVER.M_pRedisClient.HMSet(heroKeyStr, heroData)
	world.SERVER.M_pRedisClient.HMSet(equipKeyStr, equipData)

	var heros = []HeroInfo{}
	for _, data := range heroData {
		var vheroData HeroInfo
		json.Unmarshal([]byte(data.(string)), vheroData)
		heros = append(heros, vheroData)
	}

	var equips = []EquipInfo{}
	for _, data := range equipData {
		var vequipData EquipInfo
		json.Unmarshal([]byte(data.(string)), vequipData)
		equips = append(equips, vequipData)
	}

	return &PlayerData{player: *accountData, item: itemData, hero: heros, equip: equips}
}

func (this *GameCache) initPlayer() *PlayerInfo {
	return &PlayerInfo{
		id:       "address",
		nickName: "noName",
		email:    "email",
		dvt:      0,
		createTs: 0,
		updateTs: 0}
}

func (this *GameCache) initItem() []ItemInfo {
	return []ItemInfo{
		{id: 1001, count: 999},
		{id: 1002, count: 999},
		{id: 1003, count: 999},
		{id: 1004, count: 999},
	}
}

func (this *GameCache) initHero() map[string]interface{} {
	var newHero map[string]interface{}
	for i := 1; i < 6; i++ {
		heroId := int(uuid.New().ID()) + i
		heroInfo := &HeroInfo{uuid: heroId, heroType: i, equips: []string{"0", "0", "0", "0", "0"}}
		jsonStr, _ := json.Marshal(heroInfo)
		newHero[strconv.Itoa(heroId)] = string(jsonStr)
	}
	return newHero
}

func (this *GameCache) initEquip(uuid1 string) map[string]interface{} {
	var equipData map[string]interface{}
	for i := 0; i < len(datafnc.Equip_Idx); i++ {
		eId := strconv.Itoa(int(uuid.New().ID()) + i)
		uuidV, _ := strconv.Atoi(eId)
		equipInfo := &EquipInfo{uuid: uuidV, itemId: datafnc.Equip_Idx[i], heroId: "0"}
		jsonStr, _ := json.Marshal(equipInfo)
		equipData[eId] = string(jsonStr)
	}

	return equipData
}

func (this *GameCache) createAccount(key string) {

}

func (this *GameCache) createItem(key string) {

}
func (this *GameCache) createHero(key string) {

}
func (this *GameCache) createEquip(key string) {

}

func (this *GameCache) getPlayer(uuid string) *PlayerData {
	playerKeyStr := this.getAccountkey(uuid)
	itemKeyStr := this.getItemkey(uuid)
	heroKeyStr := this.getHerokey(uuid)
	equipKeyStr := this.getEquipkey(uuid)

	playerKey := world.SERVER.M_pRedisClient.Get(playerKeyStr)
	var player PlayerInfo
	json.Unmarshal([]byte(playerKey.Val()), player)

	itemKey := world.SERVER.M_pRedisClient.Get(itemKeyStr)
	var item []ItemInfo
	json.Unmarshal([]byte(itemKey.Val()), item)

	heroKey := world.SERVER.M_pRedisClient.HGetAll(heroKeyStr)
	var heros = []HeroInfo{}
	for _, data := range heroKey.Val() {
		var vhero HeroInfo
		json.Unmarshal([]byte(data), vhero)
		heros = append(heros, vhero)
	}

	equipKey := world.SERVER.M_pRedisClient.HGetAll(equipKeyStr)
	var equips = []EquipInfo{}
	for _, data := range equipKey.Val() {
		var vequipData EquipInfo
		json.Unmarshal([]byte(data), vequipData)
		equips = append(equips, vequipData)
	}

	return &PlayerData{player: player, item: item, hero: heros, equip: equips}
}

//加全属性
func (this *GameCache) addAllAttr(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "hp", num)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "allAttr", 1)
}

//减血
func (this *GameCache) desHp(round int, id string, num float64, dieTs float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "hp", num)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "allAttr", 1)

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
func (this *GameCache) updataDvt(round int, from string, to string, dvt float64) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(fromkey, "getDvt", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(fromkey, "dvt", dvt)
	world.SERVER.M_pRedisClient.HIncrByFloat(tokey, "getDvt", -1)
	world.SERVER.M_pRedisClient.HIncrByFloat(tokey, "dvt", -dvt)
}

//更新击杀数和死亡数
func (this *GameCache) updataKillAndDie(round int, from string, to string) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(fromkey, "kill", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(tokey, "die", 1)
}

func (this *GameCache) getDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "getDvt", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "dvt", num)
}
func (this *GameCache) desDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "desDvt", 1)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "dvt", -num)
}

//================================================================================================
//眩晕
func (this *GameCache) addDizzy(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 1
	hmsetData["dizzyTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除眩晕
func (this *GameCache) removeDizzy(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 0
	hmsetData["dizzyTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//禁锢
func (this *GameCache) addStopMove(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除禁锢
func (this *GameCache) removeStopMove(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 0
	hmsetData["stopmoveTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//变身
func (this *GameCache) addDeformation(round int, id string, addAtk float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除变身
func (this *GameCache) removeDeformation(round int, id string) {
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
func (this *GameCache) desSpeed(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["reduceSpeedTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//恢复速度
func (this *GameCache) recoverSpeed(round int, id string, speed float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["speed"] = speed
	hmsetData["reduceSpeedTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//加盾
func (this *GameCache) addShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["shield"] = num
	hmsetData["shieldTs"] = timed
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减盾
func (this *GameCache) desShield(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//移除盾
func (this *GameCache) removeShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	world.SERVER.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//================================================================================================
//无敌
func (this *GameCache) addImmune(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = num
	hmsetData["immuneTs"] = timed
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减无敌
func (this *GameCache) desImmune(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除无敌
func (this *GameCache) removeImmune(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	world.SERVER.M_pRedisClient.HMSet(keyTo, hmsetData)
}
