package holdup

import (
	"github.com/gin-gonic/gin"
	"middleware/token"
	"net/http"
	"strings"
)

var paths map[string]bool

func init() {
	paths = map[string]bool{
		"/member/getinfo": true,
	}
}

//拦截url验证
func CheckHold() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		//处理为全部小写
		path = strings.ToLower(path)
		_, ok := paths[path]
		if ok {
			//如果有这个键. 则此url需要进行拦截
			//TODO implement holdup check
			userInfo := token.GetTokenInfo(c)
			//这里如用户id等于0表示没登录， 且没有权限
			if userInfo.UserId == 0 {
				c.AbortWithStatusJSON(200, gin.H{
					"code": http.StatusForbidden,
					"msg":  http.StatusText(http.StatusForbidden),
					"data": nil,
				}) //立马终止 。后续中间件不会再执行, 同时后续请求也不在分发到控制器, 直接返回。
			}
		}
		// Set example variable
		//c.AbortWithStatusJSON(200, "data") //立马终止 。后续请求不在分发到控制器, 直接返回。
		//c.Next() //进入下一个中间件, 等后续中间件全部或者部分执行完, 再继续执行下面的代码
	}
}
