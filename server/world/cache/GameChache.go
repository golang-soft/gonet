package cache

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gonet/server/common"
	"gonet/server/common/mredis"
	"gonet/server/world/datafnc"
	"gonet/server/world/redisInstnace"
	"strconv"
	"time"
)

type (
	SGameCache struct {
	}

	IGameCache interface {
	}
)

var (
	GameCache *SGameCache = &SGameCache{}
)

func (this *SGameCache) getAccountkey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_user, uuid)
}

func (this *SGameCache) getItemkey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_item, uuid)
}
func (this *SGameCache) getHerokey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_hero, uuid)
}
func (this *SGameCache) getEquipkey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_equip, uuid)
}

func (this *SGameCache) getEquityitemKey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_equityitem, uuid)
}

func (this *SGameCache) getBattleInfoKey(uuid string) string {
	return fmt.Sprintf("%s:%s", common.Tables_battle, uuid)
}

type PlayerInfo struct {
	Id       string
	NickName string
	Email    string
	Dvt      int
	CreateTs int
	UpdateTs int
}

//type PlayerBasic struct {
//	id       string
//	nickName string
//	email    string
//	dvt      int
//	createTs int
//	updateTs int
//}

type ItemInfo struct {
	id    int
	count int
}

type EquipData struct {
	ItemId int
}

type HeroInfo struct {
	uuid     string
	HeroType int32
	Equips   map[int]EquipData
}

type EquipInfo struct {
	uuid   int
	itemId int
	heroId string
}

type PlayerData struct {
	Player *PlayerInfo
	item   []ItemInfo
	Hero   map[string]*HeroInfo
	equip  []EquipInfo
	equity []EquityItem
	battle []BattleInfo
}

type EquityItem struct {
	uuid   int
	itemId int
	heroId string
}

type BattleInfo struct {
	uuid   int
	itemId int
	heroId string
}

func (this *SGameCache) CeratePlayer(uuid string) *PlayerData {
	playerKeyStr := this.getAccountkey(uuid)
	itemKeyStr := this.getItemkey(uuid)
	heroKeyStr := this.getHerokey(uuid)
	equipKeyStr := this.getEquipkey(uuid)

	accountData := this.initPlayer()
	accountData.Id = uuid
	itemData := this.initItem()
	heroData := this.initHero()
	equipData := this.initEquip(uuid)

	accountDataJson, _ := json.Marshal(accountData)
	itemDataJson, _ := json.Marshal(itemData)
	redisInstnace.M_pRedisClient.Set(playerKeyStr, accountDataJson, -1)
	redisInstnace.M_pRedisClient.Set(itemKeyStr, itemDataJson, -1)
	redisInstnace.M_pRedisClient.HMSet(heroKeyStr, heroData)
	redisInstnace.M_pRedisClient.HMSet(equipKeyStr, equipData)

	var heros = make(map[string]*HeroInfo)
	for _, data := range heroData {
		var vheroData *HeroInfo
		json.Unmarshal([]byte(data.(string)), vheroData)
		heros[vheroData.uuid] = vheroData
	}

	var equips = []EquipInfo{}
	for _, data := range equipData {
		var vequipData EquipInfo
		json.Unmarshal([]byte(data.(string)), vequipData)
		equips = append(equips, vequipData)
	}

	return &PlayerData{Player: accountData, item: itemData, Hero: heros, equip: equips}
}

func (this *SGameCache) initPlayer() *PlayerInfo {
	return &PlayerInfo{
		Id:       "address",
		NickName: "noName",
		Email:    "email",
		Dvt:      0,
		CreateTs: 0,
		UpdateTs: 0}
}

func (this *SGameCache) initItem() []ItemInfo {
	return []ItemInfo{
		{id: 1001, count: 999},
		{id: 1002, count: 999},
		{id: 1003, count: 999},
		{id: 1004, count: 999},
	}
}

func (this *SGameCache) initHero() map[string]interface{} {
	var newHero map[string]interface{}
	for i := 1; i < 6; i++ {
		heroId := int(uuid.New().ID()) + i
		heroInfo := &HeroInfo{uuid: string(heroId), HeroType: int32(i), Equips: make(map[int]EquipData)}
		jsonStr, _ := json.Marshal(heroInfo)
		newHero[strconv.Itoa(heroId)] = string(jsonStr)
	}
	return newHero
}

func (this *SGameCache) initEquip(uuid1 string) map[string]interface{} {
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

func (this *SGameCache) createAccount(key string) {

}

func (this *SGameCache) createItem(key string) {

}
func (this *SGameCache) createHero(key string) {

}
func (this *SGameCache) createEquip(key string) {

}

func (this *SGameCache) GetPlayer(uuid string) *PlayerData {
	playerKeyStr := this.getAccountkey(uuid)
	itemKeyStr := this.getItemkey(uuid)
	heroKeyStr := this.getHerokey(uuid)
	equipKeyStr := this.getEquipkey(uuid)
	equityitemStr := this.getEquityitemKey(uuid) //equity
	battleInfoStr := this.getBattleInfoKey(uuid) //battleInfo

	playerKey := redisInstnace.M_pRedisClient.Get(playerKeyStr)
	var player PlayerInfo
	json.Unmarshal([]byte(playerKey.Val()), player)

	itemKey := redisInstnace.M_pRedisClient.Get(itemKeyStr)
	var item []ItemInfo
	json.Unmarshal([]byte(itemKey.Val()), item)

	heroKey := redisInstnace.M_pRedisClient.HGetAll(heroKeyStr)
	var heros = make(map[string]*HeroInfo)
	for _, data := range heroKey.Val() {
		var vhero *HeroInfo
		json.Unmarshal([]byte(data), vhero)
		heros[vhero.uuid] = vhero
	}

	equipKey := redisInstnace.M_pRedisClient.HGetAll(equipKeyStr)
	var equips = []EquipInfo{}
	for _, data := range equipKey.Val() {
		var vequipData EquipInfo
		json.Unmarshal([]byte(data), vequipData)
		equips = append(equips, vequipData)
	}

	equityitemKey := redisInstnace.M_pRedisClient.HGetAll(equityitemStr)
	var equityitems = []EquityItem{}
	for _, data := range equityitemKey.Val() {
		var vequipData EquityItem
		json.Unmarshal([]byte(data), vequipData)
		equityitems = append(equityitems, vequipData)
	}

	battleInfoKey := redisInstnace.M_pRedisClient.HGetAll(battleInfoStr)
	var battles = []BattleInfo{}
	for _, data := range battleInfoKey.Val() {
		var vequipData BattleInfo
		json.Unmarshal([]byte(data), vequipData)
		battles = append(battles, vequipData)
	}

	return &PlayerData{Player: &player, item: item, Hero: heros, equip: equips, equity: equityitems, battle: battles}
}

//加全属性
func (this *SGameCache) addAllAttr(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "hp", num)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "allAttr", 1)
}

//减血
func (this *SGameCache) desHp(round int, id string, num float64, dieTs float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "hp", num)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "allAttr", 1)

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
func (this *SGameCache) updataDvt(round int, from string, to string, dvt float64) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(fromkey, "getDvt", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(fromkey, "dvt", dvt)
	redisInstnace.M_pRedisClient.HIncrByFloat(tokey, "getDvt", -1)
	redisInstnace.M_pRedisClient.HIncrByFloat(tokey, "dvt", -dvt)
}

//更新击杀数和死亡数
func (this *SGameCache) updataKillAndDie(round int, from string, to string) {
	fromkey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(from, round)
	tokey := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(to, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(fromkey, "kill", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(tokey, "die", 1)
}

func (this *SGameCache) getDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "getDvt", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "dvt", num)
}
func (this *SGameCache) desDvt(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "desDvt", 1)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "dvt", -num)
}

//================================================================================================
//眩晕
func (this *SGameCache) addDizzy(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 1
	hmsetData["dizzyTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除眩晕
func (this *SGameCache) removeDizzy(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["dizzy"] = 0
	hmsetData["dizzyTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//禁锢
func (this *SGameCache) addStopMove(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除禁锢
func (this *SGameCache) removeStopMove(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 0
	hmsetData["stopmoveTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//变身
func (this *SGameCache) addDeformation(round int, id string, addAtk float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["stopmove"] = 1
	hmsetData["stopmoveTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除变身
func (this *SGameCache) removeDeformation(round int, id string) {
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
func (this *SGameCache) desSpeed(round int, id string, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["reduceSpeedTs"] = timed
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//恢复速度
func (this *SGameCache) recoverSpeed(round int, id string, speed float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["speed"] = speed
	hmsetData["reduceSpeedTs"] = 0
	hmsetData["posUpdateTs"] = time.Now()
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//================================================================================================
//加盾
func (this *SGameCache) addShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["shield"] = num
	hmsetData["shieldTs"] = timed
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减盾
func (this *SGameCache) desShield(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//移除盾
func (this *SGameCache) removeShield(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	redisInstnace.M_pRedisClient.HIncrByFloat(keyTo, "shield", -num)
}

//================================================================================================
//无敌
func (this *SGameCache) addImmune(round int, id string, num float64, timed float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = num
	hmsetData["immuneTs"] = timed
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//减无敌
func (this *SGameCache) desImmune(round int, id string, num float64) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}

//移除无敌
func (this *SGameCache) removeImmune(round int, id string) {
	keyTo := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(id, round)
	var hmsetData map[string]interface{}
	hmsetData["immune"] = 0
	hmsetData["immuneTs"] = 0
	redisInstnace.M_pRedisClient.HMSet(keyTo, hmsetData)
}
