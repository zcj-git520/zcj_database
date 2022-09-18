package zcj_database

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)

// 定义类型
type RedisDataBase struct {
	opts  	*redis.Options
	db	  	*redis.Client
	mu    	sync.Mutex
}

// 连接
func (d *RedisDataBase)Connect() error {
	conn := redis.NewClient(d.opts)
	_, err := conn.Do("PING").Result()
	if err != nil {
		return err
	}
	d.db = conn
	return nil
}

// 断开连接
func (d *RedisDataBase)DisConnect() error{
	return d.db.Close()
}

// 操作
func (d *RedisDataBase) Apply(f func(db *redis.Client) error) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	err := f(d.db)
	return err
}

func NewRedis(c *connDataConfig) (*RedisDataBase, error) {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	dbNum, err :=  strconv.Atoi(c.DbName)
	if err != nil {
		return nil, err
	}
	dialTimeout := time.Duration(c.TimeOut) *time.Second
	opt := &redis.Options{
		Network:            "tcp",    // 网络类型 tcp 或者 unix.
		Addr:               addr,
		OnConnect:          nil,      // 新建一个redis连接的时候，会回调这个函数
		Password:           c.PassWd, // redis数据库连接的密码
		DB:                 dbNum,       // redis数据库，序号从0开始，默认是0，可以不用设置
		MaxRetries:         5,        //  redis操作失败最大重试次数，默认不重试。
		MinRetryBackoff:    10,        // 最小重试时间间隔  默认是 8ms ; -1 表示关闭
		MaxRetryBackoff:    512,        // 最小重试时间间隔  默认是 512ms ; -1 表示关闭
		DialTimeout:        dialTimeout,        // redis连接超时时间. 默认为5秒
		//ReadTimeout:        0,        // socket读取超时时间 默认时间为3秒
		//WriteTimeout:       0,        // socket写超时时间
		//PoolSize:           0,        // redis连接池的最大连接数.默认连接池大小等于 cpu个数 * 10
		//MinIdleConns:       0,        //redis连接池最小空闲连接数.
		//MaxConnAge:         0,        // redis连接最大的存活时间，默认不会关闭过时的连接.
		//PoolTimeout:        0,        // 当你从redis连接池获取一个连接之后，连接池最多等待这个拿出去的连接多长时间。  默认是等待 ReadTimeout + 1 秒.
		//IdleTimeout:        0,        // redis连接池多久会关闭一个空闲连接. 默认是 5 分钟. -1 则表示关闭这个配置项
		//IdleCheckFrequency: 0,        // 多长时间检测一下，空闲连接  默认是 1 分钟. -1 表示关闭空闲连接检测
		//TLSConfig:          nil,      // 要使用的TLS配置。设置TLS时将进行协商

	}
	return &RedisDataBase{
		opts: opt,
		db:   nil,
		mu:   sync.Mutex{},
	}, nil
}
