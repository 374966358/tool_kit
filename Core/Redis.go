package Core

import (
	"github.com/go-redis/redis"
)

//保存redis连接
var (
	RedisDb *redis.Client
)

//创建到redis的连接
func init() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "47.57.146.4:6379",
		Password: "Luchunyu870529", // no password set
		DB:       0,                // use default DB
	})

	RedisDb.Ping().Result()
}
