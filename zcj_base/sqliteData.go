// sqlite 数据库的创建
package zcj_base

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"sync"
)

// 定义类型
type SqlLiteDataBase struct {
	db   *sql.DB
	Path string
	mu   sync.Mutex
}

// 建立连接
func (d *SqlLiteDataBase)Connect() error {

	db, err := sql.Open("sqlite3", d.Path)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)
	db.SetConnMaxLifetime(-1)
	d.db = db
	return nil
}

func (d *SqlLiteDataBase) DisConnect() error {
	return d.db.Close()
}

func (d *SqlLiteDataBase) Apply(f func(db *sql.DB) error) error{
	d.mu.Lock()
	defer d.mu.Unlock()
	err := f(d.db)
	return err
}

func NewCreatSqLite(f *FileDataConfig)(*SqlLiteDataBase, error)  {
	if f.Path == "" || f.Name == ""{
		return nil, fmt.Errorf("sql data path or name is empty")
	}
	// 文件夹不存在则创建
	_, err := os.Stat(f.Path)
	ok := os.IsNotExist(err)
	if ok{
		err := os.MkdirAll(f.Path, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("creat %v failure", f.Path)
		}
	}
	pathJoin := path.Join(f.Path, f.Name)
	dbPath := fmt.Sprintf("file:%s?cache=shared&_journal=WAL&sync=2", pathJoin)
	return &SqlLiteDataBase{
		db:   nil,
		Path: dbPath,
		mu:   sync.Mutex{},
	}, nil
}
