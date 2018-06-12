package redisConn

import (
	"conf"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"time"
)

//redis set方法
type RedisPool struct {
	redis.Pool
}

//redis set
func (pool *RedisPool) Set(key string, value interface{}) bool {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("SET", key, value)
	if str, _ := redis.String(reply, err); str == "OK" {
		return true
	}
	return false
}

//redis expire, 秒为单位
func (pool *RedisPool) Expire(key string, expire time.Duration) bool {
	redisConn := pool.Get()
	defer redisConn.Close()
	realTime := expire.Seconds()
	reply, err := redisConn.Do("EXPIRE", key, realTime)
	//如果有错误, 大多情况是断开连接, 可以重连
	if answer, _ := redis.Int64(reply, err); answer == 1 {
		//过期时间设置成功
		return true
	}
	return false
}

//redis get
func (pool *RedisPool) Gett(key string) (string, error) {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("GET", key)
	str, err := redis.String(reply, err)
	return str, err
}

//redis del,返回表示删除的个数
func (pool *RedisPool) Del(key ...interface{}) int {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("DEL", key...)
	answer, _ := redis.Int(reply, err)
	return answer
}

//redis exists
func (pool *RedisPool) Exists(key string) bool {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("EXISTS", key)
	res, _ := redis.Bool(reply, err) //num为增长后的值
	return res
}

//redis incr
func (pool *RedisPool) Incr(key string) (int64, error) {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("INCR", key)
	num, err := redis.Int64(reply, err) //num为增长后的值
	return num, err
}

//redis decr
func (pool *RedisPool) Decr(key string) (int64, error) {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("DECR", key)
	num, _ := redis.Int64(reply, err) //num为自减后的值
	return num, err
}

//如 keys token:*
func (pool *RedisPool) Keys(pattern string) []string {
	redisConn := pool.Get()
	defer redisConn.Close()
	var keys []string //获得所有键
	reply, err := redisConn.Do("KEYS", pattern)
	keys, _ = redis.Strings(reply, err)
	return keys
}

//redis lpush
func (pool *RedisPool) Lpush(key string, value string) bool {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("LPUSH", key, value)
	row, _ := redis.Int64(reply, err)
	if reflect.TypeOf(row).String() == "int64" {
		return true
	}
	return false
}

//redis rpush
func (pool *RedisPool) Rpush(key string, value string) bool {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("RPUSH", key, value)
	row, _ := redis.Int64(reply, err)
	if reflect.TypeOf(row).String() == "int64" {
		return true
	}
	return false
}

//redis lpop
func (pool *RedisPool) Lpop(key string) string {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("LPOP", key)
	str, _ := redis.String(reply, err)
	return str
}

//redis rpop
func (pool *RedisPool) Rpop(key string) string {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("RPOP", key)
	str, _ := redis.String(reply, err)
	return str
}

//重新选择redis库
func (pool *RedisPool) Select(db int) bool {
	redisConn := pool.Get()
	defer redisConn.Close()
	reply, err := redisConn.Do("SELECT", db)
	str, _ := redis.String(reply, err)
	if str == "OK" {
		return true
	}
	return false
}

//新建redis协程池
func NewPool() *RedisPool {
	//todo 完善redis连接池，和常用方法
	pool := RedisPool{}
	pool.MaxIdle = conf.MaxIdle
	pool.MaxActive = conf.MaxActive
	pool.IdleTimeout = conf.IdleTimeout
	pool.Dial = func() (redis.Conn, error) {
		var c redis.Conn
		var err error
		if conf.Mode == gin.DebugMode {
			optionDb := redis.DialDatabase(conf.DevRedisDb)
			optionTimeout := redis.DialConnectTimeout(conf.DevRedisTimeout)
			optionAuth := redis.DialPassword(conf.DevRedisAuth)
			c, err = redis.Dial("tcp", conf.DevRedisHost+":"+conf.DevRedisPort, optionDb, optionTimeout, optionAuth)
		} else {
			optionDb := redis.DialDatabase(conf.RedisDb)
			optionTimeout := redis.DialConnectTimeout(conf.RedisTimeout)
			optionAuth := redis.DialPassword(conf.RedisAuth)
			c, err = redis.Dial("tcp", conf.RedisHost+":"+conf.RedisPort, optionDb, optionTimeout, optionAuth)
		}
		if err != nil {
			panic(err.Error())
		}
		return c, err
	}
	return &pool

}
