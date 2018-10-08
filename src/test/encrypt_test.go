package test

import (
	"common/encrypt"
	"log"
	"testing"
)

func TestAesEncryptAndDecrypt(t *testing.T)  {
	var key = []byte("abcdefghijk12348") //保证key 为16,24,32位。
	var origData = []byte("hello world")
	enc, err := encrypt.AesEncrypt(origData, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("编码后: " + enc)

	org, err := encrypt.AesDecrypt(enc, key)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("解码后: " + string(org))
}