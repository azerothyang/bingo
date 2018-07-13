package ckcache

import (
	"conf"
	"controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	redis = controller.Redis
)

//初始化token
func CheckCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断请求是否为get, 只有get请求才可能缓存
		if c.Request.Method != "GET" {
			return
		}
		//判断是否有缓存, 如果有直接返回
		cacheKey := conf.CachePrefix + c.Request.URL.String()
		cache, err := redis.Gett(cacheKey)
		if err == nil && cache != "" {
			//有缓存, 直接返回
			c.Abort() //abort写在后面也可以， 最终此函数会执行完毕，到return处
			c.Header("Content-Type", "application/json;charset=utf-8")
			c.String(http.StatusOK, cache)
			return
		}
		c.Next()
	}
}
