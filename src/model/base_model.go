package model

//cjproj:Cjproj_123@tcp(119.29.96.253:3306)/policy_test_bak20171030?charset=utf8
import (
	"conf"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/jinzhu/gorm"
	"log"
)

var (
	masterDb  *gorm.DB
	slaveDb   *gorm.DB
	dsnMaster string
	dsnSlave  string
)

type baseModel struct {
}

func init() {
	var masterMaxIdle, masterMaxConn, slaveMaxIdle, slaveMaxConn int
	if conf.Mode == gin.DebugMode {
		//测试环境
		dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s",
			conf.DevMasterMysqlUserName, conf.DevMasterMysqlPassword, conf.DevMasterMysqlHost, conf.DevMasterMysqlPort, conf.DevMasterMysqlDatabase, conf.DevMasterMysqlCharset, conf.DevMasterMysqlTimeout)
		dsnSlave = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s",
			conf.DevSlaveMysqlUserName, conf.DevSlaveMysqlPassword, conf.DevSlaveMysqlHost, conf.DevSlaveMysqlPort, conf.DevSlaveMysqlDatabase, conf.DevSlaveMysqlCharset, conf.DevSlaveMysqlTimeout)
		masterMaxIdle = conf.DevMasterMaxIdle
		masterMaxConn = conf.DevMasterMaxConn
		slaveMaxIdle = conf.DevSlaveMaxIdle
		slaveMaxConn = conf.DevSlaveMaxConn
	} else {
		//生产环境
		dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local",
			conf.MasterMysqlUserName, conf.MasterMysqlPassword, conf.MasterMysqlHost, conf.MasterMysqlPort, conf.MasterMysqlDatabase, conf.MasterMysqlCharset, conf.MasterMysqlTimeout)
		dsnSlave = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&parseTime=True&loc=Local",
			conf.SlaveMysqlUserName, conf.SlaveMysqlPassword, conf.SlaveMysqlHost, conf.SlaveMysqlPort, conf.SlaveMysqlDatabase, conf.SlaveMysqlCharset, conf.SlaveMysqlTimeout)
		masterMaxIdle = conf.MasterMaxIdle
		masterMaxConn = conf.MasterMaxConn
		slaveMaxIdle = conf.SlaveMaxIdle
		slaveMaxConn = conf.SlaveMaxConn
	}
	// todo 这里的mysql长连接可能会出现断连, 之后看是否加入重连机制
	var errM, errS error
	masterDb, errM = gorm.Open("mysql", dsnMaster)
	slaveDb, errS = gorm.Open("mysql", dsnSlave)
	if errM != nil || errS != nil {
		log.Println(errM)
		log.Println(errS)
	}
	masterDb.DB().SetMaxIdleConns(masterMaxIdle)
	masterDb.DB().SetMaxOpenConns(masterMaxConn)
	slaveDb.DB().SetMaxIdleConns(slaveMaxIdle)
	slaveDb.DB().SetMaxOpenConns(slaveMaxConn)

}

//初始化map数据, 可以在新建或者更新数据的时候使用
func (*baseModel) initMapData(data map[string]string) (cols []string, values []interface{}) {
	var ks []string
	var vs []interface{}
	for k, v := range data {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return ks, vs
}
