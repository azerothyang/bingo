package token

import (
	"common/bingoredis"
	"conf"
	"encoding/json"
	"github.com/gin-gonic/gin"
	redisPkg "github.com/go-redis/redis"
	"log"
	"time"
	"util"
)

var (
	redis            = bingoredis.Redis
	tokenLength      int
	CookiePrefix     string
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
		CookiePrefix = conf.DevTokenCookieName
		redisTokenPrefix = conf.DevTokenName
		tokenExpire = conf.DevTokenExpire
	} else {
		tokenEnable = conf.TokenEnable
		tokenLength = conf.TokenLength
		CookiePrefix = conf.TokenCookieName
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
			userToken, err := c.Cookie(CookiePrefix)
			if err != nil || userToken == "" {
				//如果浏览器没有cookie则, 初始化cookie
				tokenStr := initToken()
				//设置cookie，token生成成功则初始化cookie
				if tokenStr != "" {
					c.SetCookie(CookiePrefix, tokenStr, int(tokenExpire.Seconds()), "", "", false, true)
				}
			}
		}
		c.Next()
	}
}

//初始化token, 返回token
func initToken() string {
	//生成随机字符串, 如果redis中不包含
	var tokenStr string
	for {
		tokenStr = util.RandomStr(tokenLength)
		//如果token没有重复则退出
		_, err := redis.Get(redisTokenPrefix + tokenStr).Result()
		if err == redisPkg.Nil {
			log.Println(err)
			break
		}
	}
	res := SetToken(tokenStr, &Token{}, tokenExpire)
	if res == nil {
		return tokenStr
	}
	return ""
}

//设置token
func SetToken(token string, value interface{}, expire time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = redis.Set(redisTokenPrefix + token, v, expire).Result()
	return err
}

//获取redis中存的用户信息
func GetToken(c *gin.Context) *Token {
	//先判断是否客户端已经有token, 有的话就直接解析, 没有token直接初始化为空数据
	userToken, err := c.Cookie(CookiePrefix)
	tokenInfo := Token{}
	if err == nil && userToken != "" {
		tokenMarshal, err := redis.Get(redisTokenPrefix + userToken).Result()
		if err == nil && tokenMarshal != "" {
			err = json.Unmarshal([]byte(tokenMarshal), &tokenInfo)
			if err == nil {
				return &tokenInfo
			}
		}
	}
	return &tokenInfo
}

//删除token
func DelToken(c *gin.Context) error {
	//先判断是否客户端已经有token, 有的话就直接解析, 没有token直接初始化为空数据
	userToken, err := c.Cookie(CookiePrefix)
	if err != nil {
		return err
	}
	_, err = redis.Del(redisTokenPrefix + userToken).Result()
	return err
}
