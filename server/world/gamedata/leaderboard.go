package gamedata

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gonet/server/common"
	"gonet/server/world/datafnc"
	"gonet/server/world/redisInstnace"
	"math"
	"time"
)

//游戏逻辑处理
type (
	SLeaderboardCtrl struct {
	}
	ISLeaderboardCtrl interface {
	}
)

var LeaderboardCtrl = &SLeaderboardCtrl{}

func (this *SLeaderboardCtrl) getLeaderboardKeyPerfix() string {
	return fmt.Sprintf("gvg_%v", datafnc.GetLeaderboardPrefix())
}

func (this *SLeaderboardCtrl) getLeaderboardKeyHistoryPerfix() string {
	return fmt.Sprintf("gvg_%v", datafnc.GetLeaderboardPrefix()-1)
}

func (this *SLeaderboardCtrl) getCurrLeaderboard(leaderboardKey string) string {
	return leaderboardKey + ":curr"
}

func (this *SLeaderboardCtrl) getHistoryLeaderboard(leaderboardKey string) string {
	return leaderboardKey + ":history"
}

func (this *SLeaderboardCtrl) getLastRewardKey() string {
	return this.getLeaderboardKeyHistoryPerfix() + ":reward"
}

func (this *SLeaderboardCtrl) getLeaderboardKey(itype int, mode int, roleType int) string {
	var perfix string = this.getLeaderboardKeyPerfix()
	var roleNmae string = ""
	var modeName string = ""
	if mode == common.LeaderBoardMode.Kill {
		modeName = common.LeaderBoardModePrefix.Kill
	} else if mode == common.LeaderBoardMode.Win {
		modeName = common.LeaderBoardModePrefix.Win
	}

	if roleType == common.LeaderBoardRole.All {
		roleNmae = common.LeaderBoardRolePrefix.All
	} else if roleType == common.LeaderBoardRole.Ranger {
		roleNmae = common.LeaderBoardRolePrefix.Ranger
	} else if roleType == common.LeaderBoardRole.Alchemist {
		roleNmae = common.LeaderBoardRolePrefix.Alchemist
	} else if roleType == common.LeaderBoardRole.Warrior {
		roleNmae = common.LeaderBoardRolePrefix.Warrior
	} else if roleType == common.LeaderBoardRole.Adventurer {
		roleNmae = common.LeaderBoardRolePrefix.Adventurer
	} else if roleType == common.LeaderBoardRole.Rogue {
		roleNmae = common.LeaderBoardRolePrefix.Rogue
	}

	if itype == common.LeaderBoardType.History {
		perfix = this.getLeaderboardKeyHistoryPerfix()
	}
	return perfix + ":" + modeName + ":" + roleNmae
}

type LeaderBoardData struct {
	userid string
	role   int
	kill   float64
	isWin  bool
}

func (this *SLeaderboardCtrl) setLeaderboard(data LeaderBoardData) {
	userid := data.userid
	role := data.role
	kill := data.kill
	isWin := data.isWin

	//击杀记录
	//单职业击杀
	var roleKill = this.getLeaderboardKey(common.LeaderBoardType.Week, common.LeaderBoardMode.Kill, role)
	this.add2Leaderboard(roleKill, userid, kill)
	//全职业击杀
	var allKill = this.getLeaderboardKey(common.LeaderBoardType.Week, common.LeaderBoardMode.Kill, common.LeaderBoardRole.All)
	this.add2Leaderboard(allKill, userid, kill)
	//胜场纪录
	if isWin {
		//单职业胜场
		var roleWin = this.getLeaderboardKey(common.LeaderBoardType.Week, common.LeaderBoardMode.Win, role)
		this.add2Leaderboard(roleWin, userid, 1)
		//全职业胜场
		var allWin = this.getLeaderboardKey(common.LeaderBoardType.Week, common.LeaderBoardMode.Win, common.LeaderBoardRole.All)
		this.add2Leaderboard(allWin, userid, 1)
	}
}

func (this *SLeaderboardCtrl) add2Leaderboard(leaderboard string, userid string, score float64) {
	var rankScore = redisInstnace.M_pRedisClient.ZScore(leaderboard, userid)
	var newScroe float64 = 0
	if rankScore != nil {
		var oldScore = math.Floor(float64(rankScore.Val()) / float64(common.LeaderBoard.BaseNum)) // 对应分数
		newScroe = (oldScore+score)*float64(common.LeaderBoard.BaseNum) + float64(time.Now().Unix())
	} else {
		newScroe = score*float64(common.LeaderBoard.BaseNum) + float64(time.Now().Unix())
	}
	if newScroe > 0 {
		redisInstnace.M_pRedisClient.ZAdd(leaderboard, redis.Z{Member: userid, Score: newScroe})
	}
}

func (this *SLeaderboardCtrl) getLeaderboard(itype int, mode int, role int) interface{} {
	var data interface{}
	var leaderboardKey = this.getLeaderboardKey(itype, mode, role)
	var redisKey string = ""
	if itype == common.LeaderBoardType.History {
		redisKey = this.getHistoryLeaderboard(leaderboardKey)
	} else {
		redisKey = this.getCurrLeaderboard(leaderboardKey)
	}
	var datastr = redisInstnace.M_pRedisClient.Get(redisKey)
	if datastr != nil {
		json.Unmarshal([]byte(datastr.Val()), data)
	}

	return data
}
