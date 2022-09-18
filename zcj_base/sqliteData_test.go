package zcj_base

import (
	"fmt"
	"testing"
)

func TestConnectRedis(t *testing.T) {
	f := &FileDataConfig{
		Path: "./sqliteDbPath",
		Name: "test1.db",
	}
	sqlLitDb, err := NewCreatSqLite(f)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = sqlLitDb.Connect()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("sqlite 连接成功 ……")
}

