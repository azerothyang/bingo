package ckcache

import (
	"github.com/gin-gonic/gin"
	"conf"
	"controller"
	"net/http"
)

var	(
	redis = controller.Redis
)

//初始化token
func CheckCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否有缓存, 如果有直接返回
		cacheKey:= conf.CachePrefix + c.Request.URL.String()
		cache, err := redis.Gett(cacheKey)
		if err == nil && cache != ""{
			//有缓存, 直接返回
			c.Abort() //abort写在后面也可以， 最终此函数会执行完毕，到return处
			c.Header("Content-Type", "application/json;charset=utf-8")
			c.String(http.StatusOK, cache)
			return
		}
		c.Next()
	}
}
