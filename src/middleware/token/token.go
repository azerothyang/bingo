package token

import (
	"conf"
	"controller"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
	"util"
)

var (
	redis            = controller.Redis
	tokenLength      int
	cookiePrefix     string
	redisTokenPrefix string
	tokenExpire      time.Duration
)

//token结构, 后续可以补充，目前只有UserId
type Token struct {
	UserId int64 `json:"user_id"`
}

//是否开启token
func enableToken() bool {
	var tokenEnable bool
	if conf.Mode == gin.DebugMode {
		tokenEnable = conf.DevTokenEnable
		tokenLength = conf.DevTokenLength
		cookiePrefix = conf.DevTokenCookieName
		redisTokenPrefix = conf.DevTokenName
		tokenExpire = conf.DevTokenExpire
	} else {
		tokenEnable = conf.TokenEnable
		tokenLength = conf.TokenLength
		cookiePrefix = conf.TokenCookieName
		redisTokenPrefix = conf.TokenName
		tokenExpire = conf.TokenExpire
	}
	//如果没有开启token, 直接返回
	return tokenEnable
}

//初始化token
func HandleToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断是否需要初始化token
		if enableToken() {
			//先判断是否客户端已经有token, 有的话就直接解析, 没有token直接初始化为空数据
			userToken, err := c.Cookie(cookiePrefix)
			if err != nil || userToken == "" {
				//如果浏览器没有cookie则, 初始化cookie
				tokenStr := initToken()
				//设置cookie，token生成成功则初始化cookie
				if tokenStr != "" {
					c.SetCookie(cookiePrefix, tokenStr, int(tokenExpire.Seconds()), "", "", false, true)
				}
			}
		}
		c.Next()
	}
}

//初始化token, 返回token
func initToken() string {
	//生成随机字符串, 如果redis中不包含
	var tokenStr, redisKey string
	for {
		tokenStr = util.RandomStr(tokenLength)
		//如果token没有重复则退出
		redisKey = redisTokenPrefix + tokenStr
		if !redis.Exists(redisKey) {
			break
		}
	}
	res := SetToken(redisKey, &Token{}, tokenExpire)
	if res {
		return tokenStr
	}
	return ""
}

//设置token
func SetToken(token string, value interface{}, expire time.Duration) bool {
	v, err := json.Marshal(value)
	if err != nil {
		return false
	}
	resSet := redis.Set(token, v)
	resExp := redis.Expire(token, expire)
	return resSet && resExp
}

//获取redis中存的用户信息
func GetTokenInfo(c *gin.Context) *Token {
	//先判断是否客户端已经有token, 有的话就直接解析, 没有token直接初始化为空数据
	userToken, err := c.Cookie(cookiePrefix)
	tokenInfo := Token{}
	if err == nil && userToken != "" {
		if tokenStr, err := redis.Gett(redisTokenPrefix + userToken); err == nil {
			if tokenStr != "" {
				err = json.Unmarshal([]byte(tokenStr), &tokenInfo)
				if err == nil {
					return &tokenInfo
				}
			}
		}
	}
	return &tokenInfo
}

//删除token
func DelToken(c *gin.Context) int {
	//先判断是否客户端已经有token, 有的话就直接解析, 没有token直接初始化为空数据
	userToken, err := c.Cookie(cookiePrefix)
	if err != nil {
		return 0
	}
	return redis.Del(redisTokenPrefix + userToken)
}
