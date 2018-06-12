package controller

import (
	"common/redisConn"
	_ "common/redisConn"
	"github.com/gin-gonic/gin"
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

func (b *BaseController) Method(c *gin.Context) {

}
