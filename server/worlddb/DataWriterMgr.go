package worlddb

import (
	"database/sql"
	"encoding/json"
	"gonet/actor"
	"gonet/db"
	"gonet/server/common"
	"gonet/server/common/mredis"
	"gonet/server/model"
	"strconv"
	"time"
)

type (
	DataWriterMgr struct {
		actor.Actor

		m_DataWriterMap map[int64]*DataWriter
		m_db            *sql.DB
	}

	IDataWriterMgr interface {
		actor.IActor
	}
)

var (
	DATA_MGR DataWriterMgr
)

type SettleData struct {
	Game  map[string]string   `json:"game"`
	Users []map[string]string `json:"users"`
}

func (this *DataWriterMgr) Init() {
	SERVER.m_Log.Debugf("数据库初始化 ...................................")

	this.m_db = SERVER.GetDB()
	this.Actor.Init()
	this.m_DataWriterMap = make(map[int64]*DataWriter)
	this.RegisterTimer(5*60*time.Second, this.SaveToDB) //定时器
	this.Actor.Start()
}

func (this *DataWriterMgr) SaveToDB() {
	SERVER.m_Log.Debugf("SaveToDB ...................................")
	this.SaveRound(10000)
}

func (this *DataWriterMgr) SaveRound(round int) {
	keyRound := mredis.REDIS_KEYS[mredis.KEYS_user_round_players] + strconv.Itoa(round)
	roundUser := SERVER.m_pRedisClient.SMembers(keyRound)
	mapdatas, err := roundUser.Result()
	if err != nil {
		SERVER.m_Log.Debugf("error[%v]", err)
		return
	}
	var keys []string
	var users []map[string]string

	for _, user := range mapdatas {
		keyUser := mredis.REDIS_KEYS[mredis.KEYS_user_round_basic] + common.GetRoundKey(user, round)
		keys = append(keys, keyUser)

		userMsg := SERVER.m_pRedisClient.HGetAll(keyUser)
		userMsgResult, err := userMsg.Result()
		if err != nil {
			SERVER.m_Log.Debugf("转换数据错误: %v", err)
			return
		}
		users = append(users, userMsgResult)
	}

	keyGame := mredis.REDIS_KEYS[mredis.KEYS_game_round] + strconv.Itoa(round)
	keys = append(keys, keyGame)
	gameMsg := SERVER.m_pRedisClient.HGetAll(keyGame)
	gameMstData, _ := gameMsg.Result()

	data := model.Settle{}
	data.Id = 0
	data.Round = int64(round)
	data.Ts = time.Now()

	settleData := SettleData{Game: gameMstData, Users: users}

	databyte, err := json.Marshal(settleData)
	if err != nil {
		SERVER.m_Log.Debugf("转换对象出错, %v", err)
		return
	}
	data.Data = string(databyte)
	//先查询数据是否存在
	sql := db.LoadSql(data, db.WithWhere(&model.Settle{Round: int64(round)}), db.WithLimit(10))

	rows, err := this.m_db.Query(sql)
	if err == nil {
		rs := db.Query(rows, err)
		if rs.Next() {
			roundId := rs.Row().Int("round")
			if roundId >= 1 { //创建玩家上限
				//更新
				this.m_db.Exec(db.UpdateSql(data))
			} else {
				//inserttodb
				res, err := this.m_db.Exec(db.InsertSql(data))
				if err == nil {
					SERVER.m_Log.Debugf("插入数据库错误, %v", err)
				}
				SERVER.m_Log.Debugf("插入数据库结果 %v", res)
			}
		}
	}

	//expiredata
	for _, key := range keys {
		SERVER.m_pRedisClient.Expire(key, 1*time.Second)
	}
}
