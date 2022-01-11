package gamedata

import (
	"container/list"
	"fmt"
	"github.com/goinggo/mapstructure"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/common/mredis"
	"gonet/server/world/cache"
	common2 "gonet/server/world/common"
	"gonet/server/world/datafnc"
	"gonet/server/world/redisInstnace"
	"gonet/server/world/socket"
	"gonet/server/world/table"
	"reflect"
	"sync"
	"time"
)

type GameData struct {
	user         map[string]*data.UserGameAttrData
	round        int
	startTs      int64
	endTs        int64
	flagOwner    int32
	flagUser     string
	flagUpdateTs int64
	part1score   float64
	part1Ts      int64
	part2score   float64
	part2Ts      int64
	userIdx      interface{}
}
type (
	ServerGame struct {
		User map[string]*socket.SocketData
		Game map[int]*GameData
	}

	IServerGame interface {
	}
)

var insServerGame *ServerGame
var onceServerGame sync.Once
var GGame = &ServerGame{}

//添加用户
func (this *ServerGame) addUser(uuid string, data *socket.SocketData) {
	if this.User[uuid] == nil {
		this.User[uuid] = data
	}
}

//检查用户
func (this *ServerGame) CheckUserEmpty(uuid string) bool {
	return this.User[uuid] == nil
}

//移除用户
func (this *ServerGame) RemoveUser(uuid string) {
	delete(this.User, uuid)
}

//全部用户
func (this *ServerGame) getAllUser(round int) interface{} {
	if this.Game[round] != nil && this.Game[round].user != nil {
		return this.Game[round].user
	}
	return nil
}

//游戏数数据
func (this *ServerGame) getGameData() map[int]*GameData {
	if this.Game != nil {
		return this.Game
	}
	return nil
}

//同阵营
func (this *ServerGame) getPartUser(round int, part int32) list.List {
	var userList list.List

	if this.Game[round] != nil && this.Game[round].user != nil {
		for _, user := range this.Game[round].user {
			if user.Part == part {
				userList.PushBack(user)
			}
		}
	}

	return userList
}

//根据id获取用户数据
func (this *ServerGame) GetUserById(round int, uuid string) *data.UserGameAttrData {
	if this.Game[round] != nil && this.Game[round].user[uuid] != nil {
		return this.Game[round].user[uuid]
	}

	keyFrom := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(uuid, round)
	playData := redisInstnace.M_pRedisClient.HGetAll(keyFrom)

	var userInfo data.UserGameAttrData
	err := mapstructure.Decode(playData.Val(), &userInfo)
	if err != nil {
		//world.SERVER.M_Log.Debugf("转换对象出错 %v", err)
		return nil
	}

	if this.Game[round].user != nil {
		this.Game[round].user[uuid] = &userInfo
		return &userInfo
	}

	return nil
}

//round=============================================================================================================
//注册战斗
func (this *ServerGame) registerGame(round int, userIdx []string, userObj map[string]*data.UserGameAttrData, game data.RoundGameData) *GameData {
	//处事很好
	this.Game[round] = &GameData{}

	//回合数
	this.Game[round].round = game.Round
	//开始时间
	this.Game[round].startTs = game.StartTs
	//结算时间
	this.Game[round].endTs = game.EndTs
	//持旗队伍
	this.Game[round].flagOwner = game.FlagOwner
	//持旗队伍
	this.Game[round].flagUser = game.FlagUser
	//持续持旗时间
	this.Game[round].flagUpdateTs = game.FlagUpdateTs
	//队伍1得分
	this.Game[round].part1score = game.Part1score
	//占旗时间
	this.Game[round].part1Ts = game.Part1Ts
	//队伍2得分
	this.Game[round].part2score = game.Part2score
	//占旗时间
	this.Game[round].part2Ts = game.Part2Ts
	//用户名队列
	this.Game[round].userIdx = userIdx

	v := make(map[string]interface{})
	common2.ConvertRoundData(v, game)
	cache.BattleCache.NewRoundGameData(round, "", v)
	cache.BattleCache.NewRoundUserIdx(round, userIdx)

	for key, obj := range userObj {
		//初始化用户
		maps := common2.ConvertStructToMap(obj)
		cache.BattleCache.NewRoundPlayer(round, userObj[key].User, maps)
		//装载道具
		// await this.itemInit(round, key)

		//装载道具
		// await this.equipInit(round, userObj[key].user, userObj[key].equip)
	}

	// console.log(this.Game[round]);
	this.Game[round].user = userObj
	return this.Game[round]
}

func (this *ServerGame) delGame(round int) {
	if this.Game[round] != nil {
		//退出战斗房间
		var users []string
		SocketLeaveBattleRoom(users, round)
		delete(this.Game, round)
	}
}

//根据战斗场次获取全部用户
func (this *ServerGame) getUsersByRound(battleId int) map[string]*data.UserGameAttrData {
	return this.Game[battleId].user
}

//战斗结算
func (this *ServerGame) endGame(round int) *GameData {
	if this.Game[round] != nil {
		this.Game[round].endTs = time.Now().Unix()
		cache.BattleCache.BattleEnd(round, float64(time.Now().Unix()))
		return this.Game[round]
	}
	return nil
}

//根据回合获取站场数据
func (this *ServerGame) getRoundGameData(round int) *GameData {
	if this.Game[round] != nil {
		return this.Game[round]
	}
	return nil
}

//添加机器人
func (this *ServerGame) addRobot(round int, user string) {
	//key := common.GetRoundKey(user, round)
	//let role = this.User[user].role
	//let part = this.User[user].part
}

//play=============================================================================================
//新建玩家数据
func (this *ServerGame) newRoundPlayer(round int, user string, player *data.UserGameAttrData) {
	if this.Game[round] != nil && this.Game[round].user != nil {
		this.Game[round].user[user] = player
		data := common2.ConvertStructToMap(player)
		cache.BattleCache.NewRoundPlayer(round, user, data)
	}
}

//TODO:玩家复活，同步缓存
func (this *ServerGame) relivePlayer(round int, user string, player *data.UserGameAttrData) *data.UserGameAttrData {
	if this.Game[round] != nil && this.Game[round].user != nil && this.Game[round].user[user] != nil {
		now := time.Now().Unix()
		this.Game[round].user[user].Hp = player.Hp
		this.Game[round].user[user].Hp = player.Hp
		this.Game[round].user[user].X = player.X
		this.Game[round].user[user].Y = player.Y
		this.Game[round].user[user].Shield = 0
		this.Game[round].user[user].ShieldTs = 0
		this.Game[round].user[user].Immune = 0
		this.Game[round].user[user].ImmuneTs = 0
		this.Game[round].user[user].Thorns = 0
		this.Game[round].user[user].ThornsTs = 0
		this.Game[round].user[user].StopmoveTs = 0
		this.Game[round].user[user].ReduceSpeedTs = 0
		this.Game[round].user[user].Stopmove = 0
		this.Game[round].user[user].AddDef = 0
		this.Game[round].user[user].AddDefTs = 0
		this.Game[round].user[user].AddAtk = 0
		this.Game[round].user[user].PosUpdateTs = now
		this.Game[round].user[user].UpdateTs = now
		//重制玩家数据
		data := common2.ConvertStructToMap(this.Game[round].user[user])
		cache.BattleCache.NewRoundPlayer(round, user, data)
		return this.Game[round].user[user]
	}
	return nil
}

//更新技能cd
func (this *ServerGame) updateSkillCD(round int, user string, skillId int32, cd int64) {
	player := this.Game[round].user[user]

	if player != nil {
		player.UpdateTs = cd
		fieldName := fmt.Sprintf("skill_%d_cd", skillId)
		field := reflect.ValueOf(player).Elem().FieldByName(fieldName)
		field.SetInt(cd)
		cache.BattleCache.UpdateSkillCd(round, user, skillId, float64(cd))
	}
}

//获取血量
func (this *ServerGame) getHp(round int, id string) float64 {
	userData := this.GetUserById(round, id)
	return userData.Hp
}

//加血
func (this *ServerGame) addHp(round int, id string, addhp float64, time int) float64 {
	var hp float64 = 0
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Hp += addhp
		//缓存数据
		cache.BattleCache.AddHp(round, id, addhp)
		hp = this.Game[round].user[id].Hp
	}
	return hp
}

//添加全属性
func (this *ServerGame) addAllAttr(round int, id string, addhp float64, time int) float64 {
	var hp float64 = 0
	userData := this.Game[round].user[id]
	if userData != nil {
		userData.Hp += addhp
		userData.AllAttr = 1
		//缓存数据
		cache.BattleCache.AddAllAttr(round, id, float64(addhp))
		hp = userData.Hp
	}
	return hp
}

//扣血
func (this *ServerGame) desHp(round int, from string, to string, damage float64, time0 int) bool {
	isDie := false
	dieTs := 0
	toUserData := this.Game[round].user[to]
	fromUserData := this.Game[round].user[from]
	if toUserData != nil {
		var realDamage float64 = 0
		if toUserData.Hp < damage {
			realDamage = toUserData.Hp
			dieTs = int(time.Now().Unix())
		} else {
			realDamage = damage
		}
		toUserData.Hp -= realDamage
		//缓存更新血量
		cache.BattleCache.DesHp(round, to, float64(realDamage), float64(dieTs))
		//更新dps
		this.addDps(round, from, to, (realDamage))

		if toUserData.Hp < 0 {
			//死亡
			isDie = true
			toUserData.Hp = 0
			toUserData.DieTs = time.Now().Unix()

			//dvt 易手 任何一方到达上限都不做易手
			//A < maxKill && B>maxDeath
			if fromUserData.GetDvt < datafnc.Max_Kill && toUserData.GetDvt > datafnc.Max_Death {
				this.Game[round].user[from].GetDvt += 1
				this.Game[round].user[from].Dvt += int32(datafnc.Kill_Reward_Devt)
				this.Game[round].user[to].GetDvt -= 1
				this.Game[round].user[to].Dvt -= int32(datafnc.Kill_Reward_Devt)
				//同步缓存
				cache.BattleCache.UpdataDvt(round, from, to, float64(datafnc.Kill_Reward_Devt))

				GvgBattleBroadcastAll("UPDATE_DEVT", round,
					&cmessage.UpdateDvtResp{
						PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_UpdateDvtResp), 0),
						From:       from,
						To:         to,
						Get:        int32(datafnc.Kill_Reward_Devt),
					},
				)
			}

			//更新击杀死亡数
			fromUserData.Kill += 1
			toUserData.Die += 1
			cache.BattleCache.UpdataKillAndDie(round, from, to)
			//TODO:排行榜（总榜，周榜，单职业）
		}
	}

	return isDie
}

func (this *ServerGame) addDps(round int, from string, to string, dps float64) {
	fromUserData := this.Game[round].user[from]
	fromUserData.Dps += dps
	cache.BattleCache.AddDps(round, from, float64(dps))
	this.captureFlagBreak2(round, to)
}

//打断夺旗
//async captureFlagBreak(round: number, user: string) {
//if (this.Game[round].flagUser == user) {
//let part = this.Game[round].user[user].part
////在夺旗时间内
//if (this.Game[round][`part${part}Ts`] + Score_Time_Interval > Date.now()) {
//let flagPart = this.Game[round].flagOwner
//let otherPart = 0
//if (flagPart == Constants.Part.part_1) {
//otherPart = Constants.Part.part_2
//} else {
//otherPart = Constants.Part.part_1
//}
//if (this.Game[round][`part${otherPart}Ts`] > 0) {
//this.Game[round].flagOwner = otherPart
//} else {
//this.Game[round][`part${flagPart}Ts`] = 0
//this.Game[round].flagOwner = 0
//this.Game[round].flagUpdateTs = 0
//}
//
//await BattleCache.updataFlagPart(round, part)
//}
//}
//}
//打断夺旗
func (this *ServerGame) captureFlagBreak2(round int, user string) {
	if this.Game[round].flagUser == user {
		part := this.Game[round].user[user].Part
		//在夺旗时间内
		fieldName := fmt.Sprintf("part%dTs", 0)
		field := reflect.ValueOf(this.Game[round]).Elem().FieldByName(fieldName)
		field.SetInt(0)
		this.Game[round].flagUser = ""
		cache.BattleCache.UpdataFlagPart(round, part)
	}
}

//眩晕
func (this *ServerGame) addDizzy(round int, id string, time0 int64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Dizzy = 1
		this.Game[round].user[id].DizzyTs = int64(time0)
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.AddDizzy(round, id, float64(time0))
	}
}

//移除眩晕
func (this *ServerGame) removeDizzy(round int, id string, damage int, time0 int) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Dizzy = 0
		this.Game[round].user[id].DizzyTs = 0
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.RemoveDizzy(round, id)
	}
}

//无敌
func (this *ServerGame) addImmune(round int, id string, num float64, time0 int64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Immune = num
		this.Game[round].user[id].ImmuneTs = time0
		cache.BattleCache.AddImmune(round, id, float64(time0), 0)
	}
}

//移除眩晕
func (this *ServerGame) removeImmune(round int, id string, damage int, time0 int) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Dizzy = 0
		this.Game[round].user[id].DizzyTs = 0
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.RemoveImmune(round, id)
	}
}

//禁锢
func (this *ServerGame) addStopMove(round int, id string, time0 int64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Stopmove = 1
		this.Game[round].user[id].StopmoveTs = time0
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.AddStopMove(round, id, float64(time0))
	}
}

//移除禁锢
func (this *ServerGame) removeStopMove(round int, id string, damage float64, time0 int) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Stopmove = 0
		this.Game[round].user[id].StopmoveTs = 0
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.RemoveStopMove(round, id)
	}
}

//变s身 + 禁止移动 + 攻击
func (this *ServerGame) addDeformation(round int, id string, damage float64, time0 int64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Stopmove = 1
		this.Game[round].user[id].StopmoveTs = time0
		this.Game[round].user[id].AddAtk = damage
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.AddDeformation(round, id, float64(damage), float64(time0))
	}
}

//移除变身
func (this *ServerGame) removeDeformation(round int, id string, damage float64, time0 int) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Stopmove = 0
		this.Game[round].user[id].StopmoveTs = 0
		this.Game[round].user[id].AddAtk = 0
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.RemoveDeformation(round, id)
	}
}

//减速
func (this *ServerGame) desSpeed(round int, id string, percent int, time0 int) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].ReduceSpeedTs = int64(time0)
		cache.BattleCache.DesSpeed(round, id, float64(time0))
	}
}

//设置速度
func (this *ServerGame) setSpeed(round int, id string, percent int, time0 int) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].ReduceSpeedTs = int64(time0)
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.DesSpeed(round, id, float64(time0))
	}
}

//恢复速度
func (this *ServerGame) recoverSpeed(round int, id string, speed int64, time0 int) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Speed = speed
		this.Game[round].user[id].ReduceSpeedTs = int64(time0)
		this.Game[round].user[id].PosUpdateTs = time.Now().Unix()
		cache.BattleCache.RecoverSpeed(round, id, float64(speed), 0)
	}
}

//获取防御
func (this *ServerGame) getDefPercent(round int, id string) int {
	userData := this.GetUserById(round, id)
	if userData != nil {
		return userData.DefPercent
	}
	return 0
}

//加防
func (this *ServerGame) addDefPercent(round int, id string) int {
	userData := this.GetUserById(round, id)
	if userData != nil {
		return userData.DefPercent
	}
	return 0
}

//减防
func (this *ServerGame) desDefPercent(round int, id string) int {
	userData := this.GetUserById(round, id)
	if userData != nil {
		return userData.DefPercent
	}
	return 0
}

//加盾
func (this *ServerGame) addShield(round int, id string, num float64, times int64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Shield = num
		this.Game[round].user[id].ShieldTs = times
		//缓存数据
		cache.BattleCache.AddShield(round, id, float64(num), float64(times))
	}
}

//减盾
func (this *ServerGame) desShield(round int, from string, id string, num float64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Shield -= num
		if this.Game[round].user[id].Shield <= 0 {
			this.Game[round].user[id].Shield = 0
		}
		//盾伤害 dps
		this.addDps(round, from, id, num)
		//缓存数据
		cache.BattleCache.DesShield(round, id, float64(num))
	}
}

//移除盾
func (this *ServerGame) removeShield(round int, id string) int {
	userData := this.GetUserById(round, id)
	if userData != nil {
		return userData.DefPercent
	}
	return 0
}

//反伤
func (this *ServerGame) addThorns(round int, id string, num float64, times int64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Thorns = num
		this.Game[round].user[id].ThornsTs = times
		//缓存数据
		cache.BattleCache.AddThorns(round, id, float64(num), float64(times))
	}
}

//移除反伤
func (this *ServerGame) removeThorns(round int, id string, num float64, times float64) {
	if this.Game[round].user[id] != nil {
		this.Game[round].user[id].Thorns = 0
		this.Game[round].user[id].ThornsTs = 0
		//缓存数据
		cache.BattleCache.RemoveThorns(round, id, float64(num), float64(times))
	}
}

//扣防
func (this *ServerGame) reduceDef(round int, uuid string, def int, time0 int) {
	if this.Game[round].user[uuid] != nil {
		this.Game[round].user[uuid].DefPercent -= def
	}
}

//扣速度
func (this *ServerGame) reduceSpeed(round int, uuid string, speed int, time0 int) {
	if this.Game[round].user[uuid] != nil {
		this.Game[round].user[uuid].Speed -= int64(speed)
	}
}

//type Pos struct {
//	x             int
//	y             int
//	speed         int64
//	direction     int
//	reduceSpeedTs int
//	barrier       int
//	dizzy         int
//	dizzyTs       int64
//	stopmove      int
//	stopmoveTs    int
//	posUpdateTs   int64
//}

//pos====================================================================================================
//获取用户最新坐标
func (this *ServerGame) updateGameUserPos(round int, uuid string, pos data.UserPositionData) {
	if this.Game[round].user[uuid] != nil {
		this.Game[round].user[uuid].X = pos.X
		this.Game[round].user[uuid].Y = pos.Y
		this.Game[round].user[uuid].Speed = pos.Speed
		this.Game[round].user[uuid].ReduceSpeedTs = int64(pos.ReduceSpeedTs)
		this.Game[round].user[uuid].Direction = pos.Direction
		this.Game[round].user[uuid].Barrier = pos.Barrier
		this.Game[round].user[uuid].Dizzy = pos.Dizzy
		this.Game[round].user[uuid].DizzyTs = pos.DizzyTs
		this.Game[round].user[uuid].Stopmove = pos.StopMove
		this.Game[round].user[uuid].StopmoveTs = pos.StopMoveTs
		this.Game[round].user[uuid].PosUpdateTs = pos.PosUpdateTs
		key := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(uuid, round)
		posMap := common2.ConvertStructToMap(pos)
		redisInstnace.M_pRedisClient.HMSet(key, posMap)
	}
}

//更新技能坐标
func (this *ServerGame) updateGameUserSkillPos(round int, uuid string, pos data.UserPositionData) *data.UserPositionData {
	if this.Game[round].user[uuid] != nil {
		this.Game[round].user[uuid].X = pos.X
		this.Game[round].user[uuid].Y = pos.Y
		this.Game[round].user[uuid].Speed = 0
		this.Game[round].user[uuid].PosUpdateTs = time.Now().Unix()
		// pos.speed = 0
		key := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(uuid, round)
		posMap := common2.ConvertStructToMap(pos)
		redisInstnace.M_pRedisClient.HMSet(key, posMap)
	}

	var resData *data.UserPositionData = &data.UserPositionData{}
	resData.X = this.Game[round].user[uuid].X
	resData.Y = this.Game[round].user[uuid].Y
	resData.Speed = this.Game[round].user[uuid].Speed
	resData.ReduceSpeedTs = this.Game[round].user[uuid].ReduceSpeedTs
	resData.Direction = this.Game[round].user[uuid].Direction
	resData.Barrier = this.Game[round].user[uuid].Barrier
	resData.DizzyTs = this.Game[round].user[uuid].DizzyTs
	resData.PosUpdateTs = this.Game[round].user[uuid].PosUpdateTs
	return resData
}

//item====================================================================================================
//hash  id => {id,count,cd}
func (this *ServerGame) itemInit(round int, user string) map[int]data.ItemData {
	items := this.Game[round].user[user].Item
	if len(items) == 0 {
		items = cache.BattleCache.GetItems(user)
	}
	return items
}

func (this *ServerGame) equipInit(round int, user string, equip []int32) map[int]int {
	equipData := this.Game[round].user[user].Equip
	if equipData == nil {
		eData := make(map[int]int)
		if len(equip) >= 0 {
			/*
			   当Type=2时，为装备，填（属性编号，附加的属性值）
			   A-属性编号
			   1-攻击
			   2-防御
			   3-血量
			   4-暴击
			   5-爆伤
			   6-闪避
			   9-全属性
			   B-属性增加的值，当A为9时，填全属性增加的百分比
			*/
			for _, eId := range equip {
				itemTableCfg := (*table.ITEM_CONFIG)[eId]
				if eData[itemTableCfg.Attribute[0]] != 0 {
					eData[itemTableCfg.Attribute[0]] += eData[itemTableCfg.Attribute[1]]
				} else {
					eData[itemTableCfg.Attribute[0]] = eData[itemTableCfg.Attribute[1]]
				}
			}
		}
		this.Game[round].user[user].Equip = eData
		cache.BattleCache.NewRoundPlayerEquip(round, user, eData)
	}

	return this.Game[round].user[user].Equip
}

func (this *ServerGame) skillInit(user string) map[int]data.ItemData {
	return cache.BattleCache.GetItems(user)
}

func (this *ServerGame) getItems(round int, user string) map[int]data.ItemData {
	if this.Game[round].user[user].Item == nil {
		this.Game[round].user[user].Item = cache.BattleCache.GetItems(user)
	}
	return this.Game[round].user[user].Item
}

//添加道具
func (this *ServerGame) addItem(round int, user string, itemList []*datafnc.RewardData) {
	return
}

func (this *ServerGame) reduceItem(round int, user string, itemId int, count int) *data.ItemData {
	itemData := this.Game[round].user[user].Item[itemId]
	itemData.Count -= count
	if itemData.Count < 0 {
		itemData.Count = 0
	}
	fieldName := fmt.Sprintf("item_%d_cd", itemId)
	field := reflect.ValueOf(this.Game[round].user[user]).Elem().FieldByName(fieldName)
	field.SetInt(time.Now().Unix())

	cache.BattleCache.UpdateItemData(round, user, this.Game[round].user[user].Item)
	return &data.ItemData{ItemId: itemId, Count: itemData.Count, Cd: field.Int()}
}

//flag====================================================================================
//归属
func (this *ServerGame) updateFlagOwner(round int, user string, part int32, partTs int64, flagTs int64) {
	if this.Game[round] != nil {
		if part == 1 {
			this.Game[round].part1Ts = partTs
		}
		if part == 2 {
			this.Game[round].part2Ts = int64(partTs)
		}

		this.Game[round].flagUpdateTs = flagTs
		this.Game[round].flagUser = user
		//缓存数据
		cache.BattleCache.UpdataFlagOwner(round, user, part, float64(partTs), float64(flagTs))
	}
}

func (this *ServerGame) updateFlagPart(round int, part int32, flagUpdateTs int, user string) {
	if this.Game[round] != nil {
		this.Game[round].flagOwner = part
		this.Game[round].flagUser = user
		this.Game[round].flagUpdateTs = time.Now().Unix()
		//缓存数据
		cache.BattleCache.UpdataFlagPart(round, part)
	}
}

//加积分
func (this *ServerGame) addPartScore(round int, part int32, score float64) {
	if this.Game[round] != nil {
		fieldName := fmt.Sprintf("part%dscore", part)
		field := reflect.ValueOf(this.Game[round]).Elem().FieldByName(fieldName)
		field.SetInt(field.Int() + int64(score))
		//缓存数据
		cache.BattleCache.UpdataPartScore(round, part, float64(score))
	}
}

//更新累计时间
func (this *ServerGame) updateFlagTs(round int, part int32, time0 int64) {
	if this.Game[round] != nil {
		this.Game[round].flagUpdateTs = time0
		this.Game[round].flagOwner = part
		//缓存数据
		cache.BattleCache.UpdataFlagTs(round, part, float64(time0))
	}
}
