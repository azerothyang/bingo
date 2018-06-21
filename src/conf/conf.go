package conf

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	//全局配置start
	//根据模式，读取不同的配置文件 此配置为正式环境
	Mode        = gin.DebugMode
	MaxIdle     = 1000         //redis空闲时保持的协程数
	MaxActive   = 50000        //redis工作时最大协程数
	IdleTimeout = time.Second  //redis idle连接超时等待时间, 要小于redis连接超时时间
	CachePrefix = "lan_cache:" //redis cache前缀
	CacheExpire = time.Minute  //redis cache过期时间
	//全局配置end

	//监听网段
	Addr = "0.0.0.0:8080"

	//主mysql配置
	MasterMysqlHost     = "192.168.1.188"
	MasterMysqlPort     = "3306"
	MasterMysqlUserName = "cjproj"
	MasterMysqlPassword = "Cjproj_123"
	MasterMysqlDatabase = "policy"
	MasterMysqlCharset  = "utf8"
	MasterMysqlTimeout  = "3s"
	MasterMaxIdle       = 20 //连接池中最多可空闲maxIdle个连接 ，这里取值为20，表示即使没有数据库连接时依然可以保持20空闲的连接，而不被清除，随时处于待命状态。设 0 为没有限制。
	MasterMaxConn       = 20 //连接池支持的最大连接数，这里取值为20，表示同时最多有20个数据库连接。设 0 为没有限制。

	//从mysql配置
	SlaveMysqlHost     = "192.168.1.188"
	SlaveMysqlPort     = "3306"
	SlaveMysqlUserName = "cjproj"
	SlaveMysqlPassword = "Cjproj_123"
	SlaveMysqlDatabase = "policy"
	SlaveMysqlCharset  = "utf8"
	SlaveMysqlTimeout  = "3s"
	SlaveMaxIdle       = 20
	SlaveMaxConn       = 20

	RedisHost    = "192.168.1.188"
	RedisPort    = "6379"
	RedisAuth    = ""
	RedisTimeout = time.Second * 3
	RedisDb      = 1

	TokenEnable     = true
	TokenLength     = 64                 //token随机字符串长度
	TokenCookieName = "lan_yang"         //token在cookie的名称
	TokenName       = "lan_yang:"        //token在redis里的键前缀
	TokenExpire     = time.Hour * 24 * 7 //redis里的token过期时间, 及cookie过期时间
)
