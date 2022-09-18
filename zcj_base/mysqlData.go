package zcj_base

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)

// 定义类型
type MysqlDataBase struct {
	db 		*sqlx.DB
	mu   	sync.Mutex
	conf 	*ConnDataConfig
}

// 连接
func (d *MysqlDataBase)Connect() error {
	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.conf.UserName, d.conf.PassWd, d.conf.Host,
		d.conf.Port,d.conf.DbName)
	db, err := sqlx.Open("mysql", dbPath)
	if err != nil {
		return nil
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

// 断开连接
func (d *MysqlDataBase)DisConnect()error{
	return d.db.Close()
}

// 执行
func (d *MysqlDataBase) Apply(f func(db *sqlx.DB) error) error{
	d.mu.Lock()
	defer d.mu.Unlock()
	err := f(d.db)
	return err
}

func NewMysql(c *ConnDataConfig) (*MysqlDataBase, error) {
	return &MysqlDataBase{
		db:   nil,
		mu:   sync.Mutex{},
		conf: c,
	},nil
}
