package zcj_database

import (
	"fmt"
	"testing"
)

func TestNewRedis(t *testing.T) {
	c := &connDataConfig{
		Host:     "192.168.1.128",
		Port:     6379,
		UserName: "",
		PassWd:   "",
		DbName:   "1",
		TimeOut:  10,
	}
	redisDb, err := NewRedis(c)
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = redisDb.Connect()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("redis 链接成功 ……")

}
