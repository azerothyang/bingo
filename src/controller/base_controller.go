package controller

import (
	"common/redisConn"
	_ "common/redisConn"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"conf"
	"time"
)

type BaseController struct {
}

var (
	Redis *redisConn.RedisPool //redis连接池
)

//初始化
func init() {
	Redis = redisConn.NewPool()
}

//设置缓存, cache通过url和链接里的query组成, 暂且仅支持get或者带query的 post请求,
//expire如果等于0表示使用默认配置中的cache过期时间
func (b *BaseController) SetCache(c *gin.Context, value interface{}, expire time.Duration) error{
	res, err := json.Marshal(value)
	if err != nil {
		return err
	}
	cacheKey:= conf.CachePrefix + c.Request.URL.String()
	ok := Redis.Set(cacheKey, res)
	if ok {
		if expire == 0 {
			Redis.Expire(cacheKey, conf.CacheExpire)
		}else {
			Redis.Expire(cacheKey, expire)
		}
	}
	return nil
}
