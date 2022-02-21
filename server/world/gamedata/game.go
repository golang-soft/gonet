package gamedata

import (
	"container/list"
	"fmt"
	"gonet/server/cmessage"
	"gonet/server/common"
	"gonet/server/common/data"
	"gonet/server/common/mredis"
	"gonet/server/world/datafnc"
	"gonet/server/world/helper"
	"gonet/server/world/param"
	"gonet/server/world/redisInstnace"
	"gonet/server/world/socket"
	"math"
	"strconv"
	"time"
)

type (
	SGameCtrl struct {
	}

	ISGameCtrl interface {
	}
)

var GameCtrl = &SGameCtrl{}

//游戏逻辑处理

//TODO:游戏局数
func (this *SGameCtrl) currRound() (int, error) {
	round := redisInstnace.M_pRedisClient.HGet(mredis.REDIS_KEYS[mredis.KEYS_game_global], "round")
	return strconv.Atoi(round.Val())
}

// TODO:初始化
func (this *SGameCtrl) Init() {

}
func (this *SGameCtrl) AddUser(user string, data *socket.SocketData) {
	GGame.addUser(user, data)
}

func (this *SGameCtrl) CheckUserEmpty(user string) bool {
	return GGame.CheckUserEmpty(user)
}

//TODO:游戏数据校验，判断每局游戏状态
func (this *SGameCtrl) CheckGame() {
	now := time.Now().Unix()
	gameData := GGame.getGameData()

	for key, gamedata := range gameData {
		if this.checkEndGame(gamedata) {
			continue
		}
		this.gameScore2(key, gamedata)
		if gamedata.user != nil {
			for _, usergamedata := range gamedata.user {
				//判断用户复活时间
				if usergamedata != nil && usergamedata.Hp == 0 && usergamedata.DieTs+int64(datafnc.Resurrection_Time*common.TIME.Second) < now {
					this.relivePlayer(key, usergamedata.User, usergamedata.Part)
				}
			}
		}
	}
}

func (this *SGameCtrl) checkGameStart(round int) bool {
	gameData := GGame.getGameData()
	if gameData[round] != nil {
		if gameData[round].startTs < time.Now().Unix() && gameData[round].endTs == -1 {
			return false
		}
	}
	return true
}

//验证比赛结束
func (this *SGameCtrl) checkEndGame(gameData *GameData) bool {
	var isEnd bool = false
	//比赛时间
	if gameData.endTs > 0 {
		if int64(gameData.endTs+3*1000) < time.Now().Unix() {
			//战斗结算5秒删除内存数据
			//console.log("/战斗结算5秒删除内存数据");
			GGame.delGame(gameData.round)
		}
		return true
	}

	if gameData.part1score >= datafnc.Max_Score || gameData.part2score >= datafnc.Max_Score {
		//比赛积分结算
		//console.log("//比赛积分结算");
		//console.log(Max_Score);
		//console.log(gameData.part1score);
		//console.log(gameData.part2score);

		this.endGame(gameData.round)
		isEnd = true
	}
	if int64(gameData.startTs)+int64(datafnc.Battle_Time) < time.Now().Unix() {
		//比赛时长结算
		//console.log("比赛时长结算");
		this.endGame(gameData.round)
		isEnd = true
	}

	// let userList = await getRoomUsers(gameData.round)
	// if (userList.length == 0) {
	//     //验证单局游戏人数
	//     await this.endGame(gameData.round)
	//     isEnd = true
	// }

	return isEnd
}

//TODO:游戏结束
func (this *SGameCtrl) endGame(round int) {
	//战斗结算
	roundData := GGame.endGame(round)
	//发放战场奖励
	if roundData != nil {
		this.battleReward(roundData)
		GvgBattleBroadcastAll("GAME_END", round,
			&cmessage.GameEndResp{
				PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_GameEndResp), 0),
				Data: &cmessage.GameData{
					Round:        int64(roundData.round),
					StartTs:      roundData.startTs,
					EndTs:        roundData.endTs,
					FlagOwner:    int64(roundData.flagOwner),
					FlagUPdateTs: roundData.flagUpdateTs,
					Part1Score:   roundData.part1score,
					Part1Ts:      roundData.part1Ts,
					Part2Score:   roundData.part2score,
					Part2Ts:      roundData.part2Ts,
				},
			},
		)
		//存储sql 清除 redis
		SaveCtrl.SaveRound(round)
	}
}

//战斗奖励
func (this *SGameCtrl) battleReward(roundData *GameData) {

	var winPart int32 = 0
	var failPart int32 = 0
	//根据积分判定
	if roundData.part1score == 0 && roundData.part2score == 0 {
		//极端情况 双方都为0
		winPart = common.Part.Part_1
		failPart = common.Part.Part_2
	} else {
		if roundData.part1score >= datafnc.Max_Score && roundData.part1score > roundData.part2score {
			winPart = common.Part.Part_1
			failPart = common.Part.Part_2
		} else if roundData.part2score >= datafnc.Max_Score && roundData.part2score > roundData.part1score {
			winPart = common.Part.Part_2
			failPart = common.Part.Part_1
		} else {
			winPart = common.Part.Part_1
			failPart = common.Part.Part_2
		}
	}

	//获取单局在线用户
	userList := make([]string, 0)
	for _, user := range roundData.user {
		userList = append(userList, user.User)
	}
	//胜方在线判断
	inGameWinUsers := GetRoomUsersByPart(roundData.round, winPart)
	var winPartUserCount = len(inGameWinUsers)
	//失败在线判断
	// let inGamefailUsers: Array<string> = await getRoomUsersByPart(roundData.round, failPart)
	// let failPartUserCount = inGamefailUsers.length
	//胜方分享门票
	var shareDvt = math.Floor(float64(datafnc.Battle_Victory_Dvt_Reward / winPartUserCount))
	for _, id := range userList {
		var reward []*datafnc.RewardData
		var victory = false
		if roundData.user[id].Part == winPart {
			//胜利方奖励 碎片奖励
			reward = append(reward, datafnc.Battle_Victory_Pieces_Reward...)
			// reward.push(...Battle_Victory_Pieces_Reward)
			//TODO:胜方门票抽成奖励
			//dvt 不做道具单独发放
			//胜方dvt奖励
			victory = true
			GameCtrl.sendUserReward(roundData.round, id, victory, reward, roundData.user[id].Dvt, shareDvt)
			//TODO
			//LeaderboardCtrl.setLeaderboard({ userid: id, role: roundData.user[id].type, kill: roundData.user[id].kill, isWin: true })
		} else if roundData.user[id].Part == failPart {
			//败方碎片奖励
			// reward = Battle_Defeated_Pieces_Reward
			reward = append(reward, datafnc.Battle_Defeated_Pieces_Reward...)
			// reward.push(...Battle_Defeated_Pieces_Reward)
			//败方dvt奖励
			GameCtrl.sendUserReward(roundData.round, id, victory, reward, roundData.user[id].Dvt, 0)
			//TODO
			//LeaderboardCtrl.setLeaderboard({ userid: id, role: roundData.user[id].type, kill: roundData.user[id].kill, isWin: false })
		}
	}
}

func (this *SGameCtrl) sendUserReward(round int, user string, victory bool, rewardList []*datafnc.RewardData, dvt int32, shareDvt float64) {
	GGame.addItem(round, user, rewardList)

	//broadcastToSelf(USER_EVENT.USER.REWAED, user, { item: rewardList, victory, dvt, shareDvt })
	var rewardItems []*cmessage.RewardItem

	for _, data := range rewardList {
		ritem := &cmessage.RewardItem{
			Id:    int32(data.Id),
			Count: int32(data.Count),
		}
		rewardItems = append(rewardItems, ritem)
	}

	BroadcastToSelf("REWAED", user,
		&cmessage.RewardResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_RewardResp), 0),
			Item:       rewardItems,
			Victory:    victory,
			Dvt:        dvt,
			ShareDvt:   shareDvt,
		},
	)
}

//TODO: 游戏结算
func (this *SGameCtrl) checkIsEnd(round int) {
	//判断游戏结算
	endTs := redisInstnace.M_pRedisClient.HGet(mredis.REDIS_KEYS[mredis.KEYS_game_round]+strconv.Itoa(round), "endTs")
	//已经结束了,由于血量为0 或者断开链接均会判断,故此处再次判断
	iEndTs, err := strconv.Atoi(endTs.Val())
	if err != nil {
		return
	}
	if iEndTs == 0 || iEndTs > 0 {
		return
	}

	onlineUser := socket.GetOnlineUsers()
	partUserHP := make(map[int][]int)
	var isEnd = false
	if len(onlineUser) == 0 {
		isEnd = true
	}
	for _, user := range onlineUser {
		keyUser := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(user, round)
		part, _ := strconv.Atoi(redisInstnace.M_pRedisClient.HGet(keyUser, "part").Val())
		hp, _ := strconv.Atoi(redisInstnace.M_pRedisClient.HGet(keyUser, "hp").Val())
		if partUserHP[part] == nil {
			partUserHP[part] = []int{hp}
		} else {
			partUserHP[part] = append(partUserHP[part], hp)
		}
	}
	for _, hps := range partUserHP {
		for _, it := range hps {
			if it <= 0 {
				isEnd = true
				break
			}
		}
	}

	if isEnd {
		// 结束
		GameCtrl.endGame(1)
	}
}

//TODO:获取所有socket 链接玩家
func (this *SGameCtrl) allRoundUsers(round int) interface{} {
	sockets := getConnSockets()
	for _, sc := range sockets {
		var users = make([]string, 0)
		if round == sc.Data.Round {
			users = append(users, sc.Data.User)
		}
		return users
	}
	return nil
}

func (this *SGameCtrl) allRoundUsersFromGdata(round int) *GameData {
	return GGame.getRoundGameData(round)
}

func (this *SGameCtrl) newGame(round int) *data.RoundGameData {
	return &data.RoundGameData{Round: round, StartTs: time.Now().Unix(), EndTs: -1,
		FlagOwner:    0,
		FlagUser:     "",
		FlagUpdateTs: 0,
		Part1score:   0,
		Part1Ts:      0,
		Part2Ts:      0,
		Part2score:   0}
}

//匹配游戏
func (this *SGameCtrl) matchGameStart(red list.List, blue list.List) bool {
	round := redisInstnace.M_pRedisClient.HIncrBy(mredis.REDIS_KEYS[mredis.KEYS_game_global], "round", 1)
	iround := int(round.Val())
	result, err := round.Result()
	result -= 1
	if err != nil {
		fmt.Sprintf("获得转换结果出错")
	}
	redisInstnace.M_pRedisClient.HSet(mredis.REDIS_KEYS[mredis.KEYS_game_global], "updateTs", time.Now().Unix())

	//创建游戏
	game := this.newGame(int(result))

	var userResult []*data.UserGameAttrData = make([]*data.UserGameAttrData, 0)
	var userObj map[string]*data.UserGameAttrData = make(map[string]*data.UserGameAttrData)
	var users []string = make([]string, 0)
	var userPart map[string]int32 = make(map[string]int32)
	//red
	for rede := red.Front(); rede != nil; rede = rede.Next() {
		redData := rede.Value.(*param.TmpRoomPlayerData)
		userBasic := this.newRoundPlayer(iround, redData, common.Part.Part_1, 0)
		users = append(users, redData.User)
		userBasic.Hid = redData.Hero_id
		userBasic.Equip = redData.Equips
		userResult = append(userResult, userBasic)
		userObj[redData.User] = userBasic
		userPart[redData.User] = common.Part.Part_1
	}
	//blue
	for rede := blue.Front(); rede != nil; rede = rede.Next() {
		blueData := rede.Value.(*param.TmpRoomPlayerData)
		userBasic := this.newRoundPlayer(iround, blueData, common.Part.Part_2, 0)
		users = append(users, blueData.User)
		userBasic.Hid = blueData.Hero_id
		userBasic.Equip = blueData.Equips
		userResult = append(userResult, userBasic)
		userObj[blueData.User] = userBasic
		userPart[blueData.User] = common.Part.Part_2
	}

	GGame.registerGame(iround, users, userObj, *game)

	//socketJoin2BattleRoom(users, round, userPart)
	SocketJoin2BattleRoom(users, round.Val(), userPart)

	//gvgBattleBroadcastAll(USER_EVENT.GLOBAL.GAME_START, round, { game, users: userResult })
	userdatas := make([]*cmessage.UserData, 0)
	for _, data := range userResult {
		user := &cmessage.UserData{
			User:          data.User,
			Round:         int32(data.Round),
			Part:          int32(data.Part),
			Hid:           string(data.Hid),
			Type:          int32(data.Itype),
			DefPercent:    int32(data.DefPercent),
			Hp:            int32(data.Hp),
			UpdateTs:      int64(data.UpdateTs),
			X:             int32(data.X),
			Y:             int32(data.Y),
			Speed:         int32(data.Speed),
			ReduceSpeedTs: int32(data.ReduceSpeedTs),
			Direction:     int32(data.Direction),
			Barrier:       int32(data.Barrier),
			DizzyTs:       int64(data.DizzyTs),
			Dizzy:         int32(data.Dizzy),
			ShieldTs:      int64(data.ShieldTs),
			Shield:        int32(data.Shield),
			ImmuneTs:      int32(data.ImmuneTs),
			Immune:        int32(data.Immune),
			ThornsTs:      int32(data.ThornsTs),
			Thorns:        int32(data.Thorns),
			StopMoveTs:    int32(data.StopmoveTs),
			StopMove:      int32(data.Stopmove),
			AddDef:        int32(data.AddDef),
			AddDefTs:      int32(data.AddDefTs),
			AddAtk:        int32(data.AddAtk),
			PosUpdateTs:   int64(data.PosUpdateTs),
			DieTs:         int64(data.DieTs),
			AllAttr:       int32(data.AllAttr),
			Dvt:           int32(data.Dvt),
			GetDvt_:       int32(data.GetDvt),
			DesDvt:        int32(data.DesDvt),
			Kill:          int32(data.Kill),
			Die:           int32(data.Die),
			Dps:           int32(data.Dps),
			Skill_1101Cd:  data.Skill_1101Cd,
			Skill_1102Cd:  data.Skill_1102Cd,
			Skill_1103Cd:  data.Skill_1103Cd,
			Skill_1104Cd:  data.Skill_1104Cd,

			//Item_1002Cd          int32    `protobuf:"varint,41,opt,name=item_1002_cd,json=item1002Cd,proto3" json:"item_1002_cd,omitempty"`
			//Item_1003Cd          int32    `protobuf:"varint,42,opt,name=item_1003_cd,json=item1003Cd,proto3" json:"item_1003_cd,omitempty"`
			//Item_1004Cd          int32    `protobuf:"varint,43,opt,name=item_1004_cd,json=item1004Cd,proto3" json:"item_1004_cd,omitempty"`
			//Equip_1              int32    `protobuf:"varint,44,opt,name=equip_1,json=equip1,proto3" json:"equip_1,omitempty"`
			//Equip_3              int32    `protobuf:"varint,45,opt,name=equip_3,json=equip3,proto3" json:"equip_3,omitempty"`
			//Equip                []int32  `protobuf:"varint,46,rep,packed,name=equip,proto3" json:"equip,omitempty"`
		}
		userdatas = append(userdatas, user)
	}

	GvgBattleBroadcastAll("GAME_START", int(round.Val()),
		&cmessage.GameStartResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_GameStartResp), 0),
			Game: &cmessage.GameData{
				Round:        int64(game.Round),
				StartTs:      game.StartTs,
				EndTs:        game.EndTs,
				FlagOwner:    int64(game.FlagOwner),
				FlagUPdateTs: game.FlagUpdateTs,
				Part1Score:   game.Part1score,
				Part1Ts:      game.Part1Ts,
				Part2Score:   game.Part2score,
				Part2Ts:      game.Part2Ts,
			},
			Users: userdatas,
		},
	)
	return true
}

//匹配游戏
// async matchGameOneStart(red: any) {

//     let round = await RedisClient.hincrby(REDIS_KEYS.game_global, 'round', 1)
//     round -= 1
//     console.log('start game ', round)
//     await RedisClient.hset(REDIS_KEYS.game_global, 'updateTs', Date.now())
//     const game = {
//         round,
//         startTs: Date.now(),
//         endTs: -1,
//         flagOwner: 0,
//         flagUser: "",
//         flagUpdateTs: 0,
//         part1score: 0,
//         part1Ts: 0,
//         part2Ts: 0,
//         part2score: 0
//     }

//     let userResult: any = []
//     let userObj = {}
//     let users: Array<string> = []
//     //red
//     for (const key in red) {
//         let userBasic: any = await this.newRoundPlayer(round, red[key].user, red[key].role, Constants.Part.part_1)
//         users.push(red[key].user)
//         userResult.push(userBasic)
//         userObj[red[key].user] = userBasic
//     }
//     //blue

//     await gGame.registerGame(round, users, userObj, game)
//     await socketJoin2GvgRoom(users, round)

//     gvgBattleBroadcastAll(USER_EVENT.GLOBAL.GAME_START, round, { game, users: userResult })
//     return true
// },

//TODO:
func (this *SGameCtrl) getRoundGame(round int) *GameData {
	data := GGame.getRoundGameData(round)
	if data == nil {
		return nil
	}
	return data
}

//TODO:新建回合玩家数据
func (this *SGameCtrl) newRoundPlayer(round int, userInfo *param.TmpRoomPlayerData, part int32, attrAll int) *data.UserGameAttrData {
	userBasic := InitUserBasic(round, userInfo.User, userInfo.Hero_id, userInfo.Role, part, userInfo.Equips, attrAll, 1)
	return userBasic
}

//TODO:添加机器人
func (this *SGameCtrl) addRobot(round int, user string, hid string, role int32, part int32) *data.UserGameAttrData {
	userBasic := InitUserBasic(round, user, hid, role, part, make(map[int]int), 0, 0)
	// let { bornX, bornY } = getBornPoint(part)
	// userBasic.x = bornX
	// userBasic.y = bornY
	// userBasic.hp = bornY
	return userBasic
}

//TODO:复活
func (this *SGameCtrl) relivePlayer(round int, user string, part int32) *data.UserGameAttrData {
	userInfo := GGame.GetUserById(round, user)
	if userInfo == nil {
		return nil
	}
	if userInfo.Hp == 0 {
		userBasic := InitUserBasic(round, user, userInfo.Hid, int32(userInfo.Itype), int32(part), userInfo.Equip, userInfo.AllAttr, 0)
		newPlayer := GGame.relivePlayer(round, user, userBasic)
		//gvgBattleBroadcastAll(USER_EVENT.USER.RELIVE, round, { ...newPlayer })

		GvgBattleBroadcastAll("RELIVE", round,
			&cmessage.ReliveResp{
				PacketHead:    common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_ReliveResp), 0),
				User:          newPlayer.User,
				Round:         int32(newPlayer.Round),
				Part:          int32(newPlayer.Part),
				Hid:           string(newPlayer.Hid),
				Type:          int32(newPlayer.Itype),
				DefPercent:    int32(newPlayer.DefPercent),
				Hp:            int64(newPlayer.Hp),
				UpdateTs:      int64(newPlayer.UpdateTs),
				X:             (newPlayer.X),
				Y:             (newPlayer.Y),
				Speed:         int32(newPlayer.Speed),
				ReduceSpeedTs: (newPlayer.ReduceSpeedTs),
				Direction:     int32(newPlayer.Direction),
				Barrier:       int32(newPlayer.Barrier),
				DizzyTs:       int64(newPlayer.DizzyTs),
				Dizzy:         int32(newPlayer.Dizzy),
				ShieldTs:      int64(newPlayer.ShieldTs),
				Shield:        int32(newPlayer.Shield),
				ImmuneTs:      (newPlayer.ImmuneTs),
				Immune:        int32(newPlayer.Immune),
				ThornsTs:      (newPlayer.ThornsTs),
				Thorns:        int32(newPlayer.Thorns),
				StopMoveTs:    (newPlayer.StopmoveTs),
				StopMove:      int32(newPlayer.Stopmove),
				AddDef:        int32(newPlayer.AddDef),
				AddDefTs:      (newPlayer.AddDefTs),
				AddAtk:        int32(newPlayer.AddAtk),
				PosUpdateTs:   int64(newPlayer.PosUpdateTs),
				DieTs:         int64(newPlayer.DieTs),
				AllAttr:       int32(newPlayer.AllAttr),
				Dvt:           int32(newPlayer.Dvt),
				GetDvt_:       int32(newPlayer.GetDvt),
				DesDvt:        int32(newPlayer.DesDvt),
				Kill:          int32(newPlayer.Kill),
				Die:           int32(newPlayer.Die),
				Dps:           int32(newPlayer.Dps),
				Skill_1101Cd:  newPlayer.Skill_1101Cd,
				Skill_1102Cd:  newPlayer.Skill_1102Cd,
				Skill_1103Cd:  newPlayer.Skill_1103Cd,
				Skill_1104Cd:  newPlayer.Skill_1104Cd,
			},
		)

		return newPlayer
	}
	return nil
}

func (this *SGameCtrl) gameScore2(round int, gameData *GameData) {

	spendTime := math.Floor(float64((time.Now().Unix() - gameData.startTs) / 1000))
	//额外奖励积分
	var extraScore float64 = 0
	var extraIndex float64 = spendTime / float64(common.TIME.Minute)

	if datafnc.Extra_Time_Reward[int(extraIndex)] != 0 {
		extraScore = float64(datafnc.Extra_Time_Reward[int(extraIndex)])
	}

	var isReward = true
	if gameData.part1Ts > 0 && gameData.part1Ts > gameData.part2Ts {

		if gameData.part1Ts+datafnc.Score_Time_Interval < time.Now().Unix() {
			if gameData.flagUpdateTs == gameData.part1Ts+datafnc.Score_Time_Interval {
				GGame.updateFlagTs(round, common.Part.Part_1, time.Now().Unix())
				isReward = false
				//gvgBattleBroadcastAll(USER_EVENT.USER.FLAG_SUCCESS, round, { part: Constants.Part.part_1, part1score: gameData.part1score, part2score: gameData.part2score })
				GvgBattleBroadcastAll("FLAG_SUCCESS", round,
					&cmessage.FlagSuccessResp{
						PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_FlagSuccessResp), 0),
						Part:       common.Part.Part_1,
						Part1Score: gameData.part1score,
						Part2Score: gameData.part2score,
					},
				)
			}
		}
	} else if gameData.part2Ts > 0 && gameData.part2Ts > gameData.part1Ts {
		if gameData.part2Ts+datafnc.Score_Time_Interval < time.Now().Unix() {
			if gameData.flagUpdateTs == gameData.part2Ts+datafnc.Score_Time_Interval {
				GGame.updateFlagTs(round, common.Part.Part_2, time.Now().Unix())
				isReward = false
				//gvgBattleBroadcastAll(USER_EVENT.USER.FLAG_SUCCESS, round, { part: Constants.Part.part_2, part1score: gameData.part1score, part2score: gameData.part2score })
				GvgBattleBroadcastAll("FLAG_SUCCESS", round,
					&cmessage.FlagSuccessResp{
						PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_FlagSuccessResp), 0),
						Part:       common.Part.Part_2,
						Part1Score: gameData.part1score,
						Part2Score: gameData.part2score,
					},
				)
			}
		}
	}

	//有归属并奖励
	if gameData.flagOwner != 0 && isReward {
		if extraScore > 0 {
			GGame.addPartScore(round, gameData.flagOwner, datafnc.Second_Score+extraScore)
			//gvgBattleBroadcastAll(USER_EVENT.USER.FLAG_SUCCESS, round, { part: gameData.flagOwner, part1score: gameData.part1score, part2score: gameData.part2score })

			GvgBattleBroadcastAll("FLAG_SUCCESS", round,
				&cmessage.FlagSuccessResp{
					PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_FlagSuccessResp), 0),
					Part:       gameData.flagOwner,
					Part1Score: gameData.part1score,
					Part2Score: gameData.part2score,
				},
			)

		} else {
			GGame.addPartScore(round, gameData.flagOwner, datafnc.Second_Score)
		}
	}
}

func (this *SGameCtrl) itemInit(round int, user string) map[int]data.ItemData {
	return GGame.itemInit(round, user)
}

func (this *SGameCtrl) skillInit(user string) map[int]data.ItemData {
	return GGame.skillInit(user)
}
func (this *SGameCtrl) getItems(round int, user string) map[int]data.ItemData {
	return GGame.getItems(round, user)
}

func (this *SGameCtrl) addHp(round int, user string, add float64) {
	hp := GGame.addHp(round, user, add, 0)
	//gvgBattleBroadcastAll(USER_EVENT.USER.ITEM_ADD_HP, round, { hp: hp, add: add })
	GvgBattleBroadcastAll("ITEM_ADD_HP", round,
		&cmessage.AddHpResp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_AddHpResp), 0),
			Hp:         hp,
			Add:        add,
		},
	)
}
func (this *SGameCtrl) addAllAttr(round int, user string, percent float64, play *helper.BasicInfo, userInfo *data.UserGameAttrData) {
	equipHp := 0
	var equipAttrPercent float64 = 0
	if userInfo.Equip_3 > 0 {
		equipHp = userInfo.Equip_3
	}
	addPercent := percent + equipAttrPercent
	addHp := math.Floor(float64((play.Hp + float64(equipHp)) * addPercent))
	//添加血上限
	hp := GGame.addAllAttr(round, user, addHp, 0)
	GvgBattleBroadcastAll("ITEM_1003", round,
		&cmessage.UseItem1003Resp{
			PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_UseItem1003Resp), 0),
			Hp:         hp,
			AllAttr:    1,
		},
	)
}

func (this *SGameCtrl) useItem(round int, user string, itemId int, count int) {
	GGame.reduceItem(round, user, itemId, count)
	//init.GvgBattleBroadcastAll("USE_ITEM", round,
	//	&cmessage.UseItemResp{
	//		PacketHead: common.BuildPacketHead(cmessage.MessageID(cmessage.MessageID_MSG_UseItemResp), 0),
	//		From: user,
	//		ItemId : itemId,
	//		Msg: "",
	//	},
	//)
}

//获取夺旗
func (this *SGameCtrl) UpDateFlagOwner(round int, user string, part int) {
	userData := GGame.GetUserById(round, user)
	flagData := GGame.getRoundGameData(round)
	now := time.Now().Unix()
	if userData.Part == flagData.flagOwner {
		return
	}
	GGame.updateFlagOwner(round, user, userData.Part, now, now+datafnc.Score_Time_Interval)
}
