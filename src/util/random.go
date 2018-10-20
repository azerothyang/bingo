package util

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//生成随机字符串
func RandomStr(length int) string {
	rand.Seed(time.Now().UnixNano())
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}
