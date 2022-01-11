package redisInstnace

import (
	"gonet/base/redis"
	"gonet/server/common/mredis"
	"time"
)

var M_pRedisClient *redis.Client

var (
	round    = 10000
	updateTs = time.Now()
)

func InitRedis() {
	M_pRedisClient.HSetNX(mredis.REDIS_KEYS[mredis.KEYS_game_global], "round", round)
	M_pRedisClient.HSetNX(mredis.REDIS_KEYS[mredis.KEYS_game_global], "updateTs", updateTs)
}

func Init(Prefix string,
	Host string,
	Port string,
	Pass string,
	Db int) {

	var err error
	M_pRedisClient, err = redis.NewClient(&redis.RedisConfig{
		Prefix: Prefix,
		Host:   Host,
		Port:   Port,
		Pass:   Pass,
		Db:     Db,
	})
	if err != nil {
		//this.M_Log.Debug("初始化redis失败")
	}
}
