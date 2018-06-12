package test

import (
	"common/redisConn"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestSet(T *testing.T) {
	type person struct {
		Name string
		Age  int
	}
	p := person{
		"cheng yang",
		12,
	}
	bytes, _ := json.Marshal(&p)
	redisConn := redisConn.NewPool()
	redisConn.Set("x", bytes)
}

func TestExpire(T *testing.T) {
	redisConn := redisConn.NewPool()
	res := redisConn.Expire("x", time.Hour)
	fmt.Println(res)
}

func TestGett(T *testing.T) {
	redisConn := redisConn.NewPool()
	res, err := redisConn.Gett("x")
	if err != nil {
	}
	fmt.Println(res)
}

func TestDel(T *testing.T) {
	redisConn := redisConn.NewPool()
	num := redisConn.Del("x", "y")
	fmt.Println(num)
}

func TestLpush(T *testing.T) {
	redisConn := redisConn.NewPool()
	ok := redisConn.Lpush("list", "z")
	if ok {
		fmt.Println(11)
	}
}

func TestLpop(T *testing.T) {
	redisConn := redisConn.NewPool()
	res := redisConn.Lpop("list")
	fmt.Println(res)
}

func TestKeys(T *testing.T) {
	redisConn := redisConn.NewPool()
	fmt.Println(redisConn.Keys("policy_token*"))
}

func TestSelect(T *testing.T) {
	redisConn := redisConn.NewPool()
	fmt.Println(redisConn.Select(2))
}

func TestExists(T *testing.T) {
	redisConn := redisConn.NewPool()
	fmt.Println(redisConn.Exists("x"))
}

func TestNewPool(T *testing.T) {
	c := redisConn.NewPool().Get()
	defer c.Close()
	test, err := c.Do("GET", "x")
	fmt.Println(redis.String(test, err))
}
