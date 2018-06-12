package conf

import "time"

const (
	//监听网段
	DevAddr = "0.0.0.0:8080"

	//主mysql开发环境配置
	DevMasterMysqlHost     = "192.168.1.188"
	DevMasterMysqlPort     = "3306"
	DevMasterMysqlUserName = "cjproj"
	DevMasterMysqlPassword = "Cjproj_123"
	DevMasterMysqlDatabase = "policy"
	DevMasterMysqlCharset  = "utf8"
	DevMasterMysqlTimeout  = "3s" //带上单位
	DevMasterMaxIdle       = 20   //连接池中最多可空闲maxIdle个连接 ，这里取值为20，表示即使没有数据库连接时依然可以保持20空闲的连接，而不被清除，随时处于待命状态。设 0 为没有限制。
	DevMasterMaxConn       = 20   //连接池支持的最大连接数，这里取值为20，表示同时最多有20个数据库连接。设 0 为没有限制。

	//从mysql开发环境配置
	DevSlaveMysqlHost     = "192.168.1.188"
	DevSlaveMysqlPort     = "3306"
	DevSlaveMysqlUserName = "cjproj"
	DevSlaveMysqlPassword = "Cjproj_123"
	DevSlaveMysqlDatabase = "policy"
	DevSlaveMysqlCharset  = "utf8"
	DevSlaveMysqlTimeout  = "3s" //带上单位
	DevSlaveMaxIdle       = 20
	DevSlaveMaxConn       = 20

	DevRedisHost    = "192.168.1.188"
	DevRedisPort    = "6379"
	DevRedisAuth    = ""
	DevRedisTimeout = time.Second * 3
	DevRedisDb      = 1

	DevTokenEnable     = true
	DevTokenLength     = 64                 //token随机字符串长度
	DevTokenCookieName = "lan_yang"         //token在cookie的名称
	DevTokenName       = "lan_yang:"        //token在redis里的键前缀
	DevTokenExpire     = time.Hour * 24 * 7 //redis里的token过期时间, 及cookie过期时间
)
