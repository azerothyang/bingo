package conf

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	//全局配置start
	//根据模式，读取不同的配置文件 此配置为正式环境
	Mode        = gin.DebugMode
	//全局配置end

	//监听网段
	Addr = "0.0.0.0:8080"

	//主mysql配置
	MasterMysqlHost     = "192.168.1.188"
	MasterMysqlPort     = "3306"
	MasterMysqlUserName = "bingo"
	MasterMysqlPassword = "123456"
	MasterMysqlDatabase = "bingo"
	MasterMysqlCharset  = "utf8"
	MasterMysqlTimeout  = "0.5s"
	MasterMaxIdle       = 10 //连接池中最多可空闲maxIdle个连接 ，这里取值为20，表示即使没有数据库连接时依然可以保持20空闲的连接，而不被清除，随时处于待命状态。设 0 为没有限制。
	MasterMaxConn       = 20 //连接池支持的最大连接数，这里取值为20，表示同时最多有20个数据库连接。设 0 为没有限制。
	MasterMaxLifetime   = 8 * time.Hour   //连接重置的最长时间，可能会在之前就关闭重置

	//从mysql配置
	SlaveMysqlHost     = "192.168.1.188"
	SlaveMysqlPort     = "3306"
	SlaveMysqlUserName = "bingo"
	SlaveMysqlPassword = "123456"
	SlaveMysqlDatabase = "bingo"
	SlaveMysqlCharset  = "utf8"
	SlaveMysqlTimeout  = "0.5s"
	SlaveMaxIdle       = 10
	SlaveMaxConn       = 20
	SlaveMaxLifetime   = 8 * time.Hour

	RedisHost    = "192.168.1.188"
	RedisPort    = "6379"
	RedisAuth    = ""
	RedisDb      = 1
	RedisTimeout = 1 * time.Second

	TokenEnable     = true               //是否开启token生成, 如果是服务其实就不用开了。 web应用或者app应用需要打开
	TokenLength     = 64                 //token随机字符串长度
	TokenCookieName = "lan_yang"         //token在cookie的名称
	TokenName       = "lan_yang:"        //token在redis里的键前缀
	TokenExpire     = time.Hour * 24 * 7 //redis里的token过期时间, 及cookie过期时间
)
