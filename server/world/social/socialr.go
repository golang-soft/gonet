package social

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gonet/actor"
	"gonet/base"
	"gonet/db"
	"gonet/rd"
	"gonet/server/model"
	"gonet/server/world"
	"gonet/server/world/player"
)

//分布式考虑直接数据库
type (
	SocialMgrR struct {
		actor.Actor

		m_db  *sql.DB
		m_Log *base.CLog
	}

	ISocialMgrR interface {
		actor.IActor

		makeLink(PlayerId, TargetId int64, Type int8) int       //加好友
		destoryLink(PlayerId, TargetId int64, Type int8) int    //删除好友
		addFriendValue(PlayerId, TargetId int64, Value int) int //增加好友度
		loadSocialDB(PlayerId int64, Type int8) SOCIALITEMMAP
		loadSocialById(PlayerId, TargetId int64, Type int8) *model.SocialItem
		isFriendType(Type int8) bool
		isBestFriendType(Type int8) bool
		hasMakeLink(oldType, newType int8) bool
	}
)

var (
	SOCIALMGR SocialMgr
)

func RdKey(nPlayerId int64) string {
	return fmt.Sprintf("h_%s_%d", sqlTable, nPlayerId)
}

func (this *SocialMgrR) Init() {
	this.m_db = world.SERVER.GetDB()
	this.m_Log = world.SERVER.GetLog()
	this.Actor.Init()
	actor.MGR.AddActor(this)

	this.RegisterCall("C_W_MakeLinkRequest", func(ctx context.Context, PlayerId, TargetId int64, Type int8) {
		pPlayer := player.SIMPLEMGR.GetPlayerDataById(PlayerId)
		pTarget := player.SIMPLEMGR.GetPlayerDataById(TargetId)
		if pPlayer == nil || pTarget == nil {
			this.m_Log.Printf("查询玩家id[%d][%d]数据为空", PlayerId, TargetId)
			return
		}
		this.makeLink(PlayerId, TargetId, Type)
	})

	this.Actor.Start()
}

func (this *SocialMgrR) isFriendType(Type int8) bool {
	if Type != Temp && Type != Mute && Type != Enemy {
		return true
	}
	return false
}

func (this *SocialMgrR) isBestFriendType(Type int8) bool {
	if Type != Friend && Type != Temp && Type != Mute && Type != Enemy {
		return true
	}
	return false
}

func (this *SocialMgrR) hasMakeLink(oldType, newType int8) bool {
	if oldType == newType {
		return false
	}

	if oldType != Temp && newType == Temp {
		return false
	}

	if this.isBestFriendType(oldType) || this.isBestFriendType(newType) {
		return false
	}

	if (oldType == Friend || oldType == Mute) && newType == Enemy {
		return false
	}

	return true
}

func (this *SocialMgrR) loadSocialDB(PlayerId int64, Type int8) SOCIALITEMMAP {
	SocialMap := make(SOCIALITEMMAP)
	Item := &model.SocialItem{}
	rows, err := this.m_db.Query(db.LoadSql(Item, db.WithWhere(model.SocialItem{PlayerId: PlayerId, Type: Type})))
	rs := db.Query(rows, err)
	if rs.Next() {
		db.LoadObjSql(&Item, rs.Row())
		SocialMap[Item.TargetId] = Item
	}
	return SocialMap
}

func (this *SocialMgrR) loadSocialById(PlayerId, TargetId int64, Type int8) *model.SocialItem {
	Item := &model.SocialItem{PlayerId: PlayerId, TargetId: TargetId, Type: Type}
	rows, err := this.m_db.Query(db.LoadSql(Item))
	rs := db.Query(rows, err)
	if rs.Next() {
		db.LoadObjSql(&Item, rs.Row())
		return Item
	}
	return nil
}

func (this *SocialMgrR) makeLink(PlayerId, TargetId int64, Type int8) int {
	if Type >= Count {
		return SocialError_Unknown
	}

	if this.isBestFriendType(Type) {
		return SocialError_Unallowed
	}

	if PlayerId == TargetId {
		return SocialError_Self
	}

	SocialMap := SOCIALITEMMAP{}
	datas, err := redis.ByteSlices(rd.Do(world.RdID, func(c redis.Conn) (reply interface{}, err error) {
		return c.Do("HVALS", RdKey(PlayerId))
	}))

	for _, v := range datas {
		pData := &model.SocialItem{}
		json.Unmarshal(v, pData)
		SocialMap[pData.TargetId] = pData
	}

	if err == nil {

	} else {
		SocialMap = this.loadSocialDB(PlayerId, Type)
	}

	Item, exist := SocialMap[TargetId]
	if exist {
		firstPlayerId := int64(0)
		nCount := len(SocialMap)

		//找一个空的
		for _, v := range SocialMap {
			if v.TargetId != TargetId {
				if firstPlayerId == 0 {
					firstPlayerId = v.TargetId
					break
				}
			}
		}

		// type is full, send error
		if nCount > getSocialTypeMaxCount(Type) {
			if Type != Enemy && Type != Temp {
				return SocialError_MaxCount
			} else {
				// 当仇人或临时好友添加满时，删除一个仇人或临时好友
				if firstPlayerId != 0 && SocialError_None != this.destoryLink(PlayerId, firstPlayerId, Type) {
					return SocialError_DbError
				}
			}
		}

		if Item != nil {
			if Item.Type == Type {
				return SocialError_Existed
			} else {
				if !this.hasMakeLink(Item.Type, Type) {
					return SocialError_Unallowed
				}
			}

			Item.PlayerId = PlayerId
			Item.TargetId = TargetId
			Item.Type = Type
			Item.FriendValue = 0
			this.m_db.Exec(db.UpdateSql(Item))
			data, _ := json.Marshal(&Item)
			rd.Do(world.RdID, func(c redis.Conn) (reply interface{}, err error) {
				c.Send("HSET", RdKey(PlayerId), TargetId, data)
				c.Send("EXPIRE", RdKey(PlayerId), 5*60)
				c.Flush()
				return nil, nil
			})
			this.m_Log.Printf("更新社会关系playerId=%d,destPlayerId=%d,newType=%d", PlayerId, TargetId, Item.Type)
			return SocialError_None
		}
	} else {
		Item := &model.SocialItem{}
		Item.PlayerId = PlayerId
		Item.TargetId = TargetId
		Item.Type = Type
		Item.FriendValue = 0

		SocialMap[TargetId] = Item
		this.m_db.Exec(db.InsertSql(Item))
		data, _ := json.Marshal(&Item)
		rd.Do(world.RdID, func(c redis.Conn) (reply interface{}, err error) {
			c.Send("HSET", RdKey(PlayerId), TargetId, data)
			c.Send("EXPIRE", RdKey(PlayerId), 5*60)
			c.Flush()
			return nil, nil
		})
		this.m_Log.Printf("新增社会关系playerId=%d,destPlayerId=%d,type=%d", PlayerId, TargetId, Item.Type)
		return SocialError_None
	}

	return SocialError_DbError
}

func (this *SocialMgrR) destoryLink(PlayerId, TargetId int64, Type int8) int {
	if PlayerId == TargetId {
		return SocialError_Self
	}

	Item := &model.SocialItem{}
	Item.PlayerId = PlayerId
	Item.TargetId = TargetId
	Item.Type = Type
	Item.FriendValue = 0

	if Item.Type == Friend || Item.Type == Mute || Item.Type == Temp || Item.Type == Enemy {
		this.m_db.Exec(db.DeleteSql(Item))
		rd.Do(world.RdID, func(c redis.Conn) (reply interface{}, err error) {
			c.Do("HDEL", RdKey(PlayerId), TargetId)
			return nil, nil
		})
		return SocialError_None
	}

	return SocialError_Unallowed
}

func (this *SocialMgrR) addFriendValue(PlayerId, TargetId int64, Value int) int {
	Item1 := this.loadSocialById(PlayerId, TargetId, Friend)
	Item2 := this.loadSocialById(TargetId, PlayerId, Friend)

	if Item1 == nil || Item2 == nil {
		return 0
	}

	if !this.isFriendType(Item1.Type) || !this.isFriendType(Item2.Type) {
		return 0
	}

	Item1.FriendValue += Value
	Item2.FriendValue = Item1.FriendValue
	this.m_db.Exec(db.UpdateSql(Item1))
	this.m_db.Exec(db.UpdateSql(Item2))
	data, _ := json.Marshal(&Item1)
	data1, _ := json.Marshal(&Item2)
	rd.Do(world.RdID, func(c redis.Conn) (reply interface{}, err error) {
		c.Send("HSET", RdKey(PlayerId), TargetId, data)
		c.Send("EXPIRE", RdKey(PlayerId), 5*60)
		c.Send("HSET", RdKey(TargetId), PlayerId, data1)
		c.Send("EXPIRE", RdKey(TargetId), 5*60)
		c.Flush()
		return nil, nil
	})
	return Value
}
