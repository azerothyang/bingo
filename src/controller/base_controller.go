package controller

import (
	"conf"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"model"
)

type BaseController struct {
}

var (
	Redis     *redis.Client    //redis连接池
	userModel *model.UserModel //用户model, 声明go中会自动初始化一个变量
)

//初始化
func init() {
	switch {
	case conf.Mode == gin.ReleaseMode:
		Redis = redis.NewClient(&redis.Options{
			Addr:     conf.RedisHost + ":" + conf.RedisPort,
			Password: conf.RedisAuth, // no password set
			DB:       conf.RedisDb,   // use default DB
		})
	case conf.Mode == gin.DebugMode:
		Redis = redis.NewClient(&redis.Options{
			Addr:     conf.DevRedisHost + ":" + conf.DevRedisPort,
			Password: conf.DevRedisAuth, // no password set
			DB:       conf.DevRedisDb,   // use default DB
		})
	}

}

//设置缓存, cache通过url和链接里的query组成, 暂且仅支持get或者带query的 post请求,
//expire如果等于0表示使用默认配置中的cache过期时间
//func (b *BaseController) SetCache(c *gin.Context, value interface{}, expire time.Duration) error {
//	res, err := json.Marshal(value)
//	if err != nil {
//		return err
//	}
//	cacheKey := conf.CachePrefix + c.Request.URL.String()
//	ok := Redis.Set(cacheKey, res)
//	if ok {
//		if expire == 0 {
//			Redis.Expire(cacheKey, conf.CacheExpire)
//		} else {
//			Redis.Expire(cacheKey, expire)
//		}
//	}
//	return nil
//}
