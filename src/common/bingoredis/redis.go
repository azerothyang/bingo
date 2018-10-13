package bingoredis

import (
	"conf"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var Redis     *redis.Client    //redis连接池

func init()  {
	switch {
	case conf.Mode == gin.ReleaseMode:
		Redis = redis.NewClient(&redis.Options{
			Addr:     conf.RedisHost + ":" + conf.RedisPort,
			Password: conf.RedisAuth, // no password set
			DB:       conf.RedisDb,   // use default DB
			ReadTimeout: conf.RedisTimeout,
			WriteTimeout: conf.RedisTimeout,
			DialTimeout: conf.RedisTimeout,
		})
	case conf.Mode == gin.DebugMode:
		Redis = redis.NewClient(&redis.Options{
			Addr:     conf.DevRedisHost + ":" + conf.DevRedisPort,
			Password: conf.DevRedisAuth, // no password set
			DB:       conf.DevRedisDb,   // use default DB
			ReadTimeout: conf.DevRedisTimeout,
			WriteTimeout: conf.DevRedisTimeout,
			DialTimeout: conf.DevRedisTimeout,
		})
	}
}

