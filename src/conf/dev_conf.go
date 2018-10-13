package conf

import "time"

const (
	//监听网段
	DevAddr = "0.0.0.0:8080"

	//主mysql开发环境配置
	DevMasterMysqlHost     = "192.168.1.188"
	DevMasterMysqlPort     = "3306"
	DevMasterMysqlUserName = "bingo"
	DevMasterMysqlPassword = "123456"
	DevMasterMysqlDatabase = "bingo"
	DevMasterMysqlCharset  = "utf8"
	DevMasterMysqlTimeout  = "0.5s" //带上单位
	DevMasterMaxIdle       = 10   //连接池中最多可空闲maxIdle个连接 ，这里取值为20，表示即使没有数据库连接时依然可以保持20空闲的连接，而不被清除，随时处于待命状态。设 0 为没有限制。
	DevMasterMaxConn       = 20   //连接池支持的最大连接数，这里取值为20，表示同时最多有20个数据库连接。设 0 为没有限制。
	DevMasterMaxLifetime   = 8 * time.Hour   //连接重置的最长时间，可能会在之前就关闭重置

	//从mysql开发环境配置
	DevSlaveMysqlHost     = "192.168.1.188"
	DevSlaveMysqlPort     = "3306"
	DevSlaveMysqlUserName = "bingo"
	DevSlaveMysqlPassword = "123456"
	DevSlaveMysqlDatabase = "bingo"
	DevSlaveMysqlCharset  = "utf8"
	DevSlaveMysqlTimeout  = "0.5s" //带上单位
	DevSlaveMaxIdle       = 10
	DevSlaveMaxConn       = 20
	DevSlaveMaxLifetime   = 8 * time.Hour

	DevRedisHost    = "192.168.1.188"
	DevRedisPort    = "6379"
	DevRedisAuth    = ""
	DevRedisDb      = 1
	DevRedisTimeout = 1 * time.Second

	DevTokenEnable     = true               //是否开启token生成, 如果是服务其实就不用开了。 web应用或者app应用需要打开
	DevTokenLength     = 64                 //token随机字符串长度
	DevTokenCookieName = "lan_yang"         //token在cookie的名称
	DevTokenName       = "lan_yang:"        //token在redis里的键前缀
	DevTokenExpire     = time.Hour * 24 * 7 //redis里的token过期时间, 及cookie过期时间

)
