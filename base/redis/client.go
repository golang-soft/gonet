package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	redisClient *redis.Client
	prefix      string
}

func NewClient(redisConfig *RedisConfig) (*Client, error) {
	prefix := redisConfig.Prefix
	host := redisConfig.Host
	port := redisConfig.Port
	pass := redisConfig.Pass
	db := redisConfig.Db

	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Client{
		redisClient: client,
		prefix:      prefix,
	}, nil
}

func (this *Client) GetKey(key string) string {
	return /*this.prefix + ":" +*/ key
}

func (this *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	key = this.GetKey(key)
	return this.redisClient.Set(key, value, expiration)
}

func (this *Client) Get(key string) *redis.StringCmd {
	key = this.GetKey(key)
	return this.redisClient.Get(key)
}

func (this *Client) HSet(key, field string, value interface{}) *redis.BoolCmd {
	key = this.GetKey(key)
	return this.redisClient.HSet(key, field, value)
}

func (this *Client) HGet(key, field string) *redis.StringCmd {
	key = this.GetKey(key)
	return this.redisClient.HGet(key, field)
}

func (this *Client) HGetAll(key string) *redis.StringStringMapCmd {
	key = this.GetKey(key)
	return this.redisClient.HGetAll(key)
}

func (this *Client) Expire(key string, time time.Duration) {
	this.redisClient.Expire(key, time)
}

func (this *Client) SMembers(key string) *redis.StringSliceCmd {
	return this.redisClient.SMembers(key)
}

func (this *Client) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return this.redisClient.HIncrBy(key, field, incr)
}

func (this *Client) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	return this.redisClient.HIncrByFloat(key, field, incr)
}

func (this *Client) HMGet(key string, fields ...string) *redis.SliceCmd {
	return this.redisClient.HMGet(key, fields...)
}

func (this *Client) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
	return this.redisClient.HMSet(key, fields)
}

func (this *Client) SADD(key string, users ...interface{}) *redis.IntCmd {
	return this.redisClient.SAdd(key, users)
}

func (this *Client) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	return this.redisClient.HSetNX(key, field, value)
}

func (this *Client) ZScore(key, member string) *redis.FloatCmd {
	return this.redisClient.ZScore(key, member)
}

func (this *Client) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return this.redisClient.ZAdd(key, members...)
}
