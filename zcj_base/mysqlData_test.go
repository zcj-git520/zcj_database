package zcj_base

import (
	"fmt"
	"testing"
)

func TestNewMysql(t *testing.T) {
	c := &ConnDataConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		UserName: "root",
		PassWd:   "123456",
		DbName:   "zcj_database",
		TimeOut:  10,
	}
	mysqlDb, err := NewMysql(c)
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = mysqlDb.Connect()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("mysql 连接成功……")
}
