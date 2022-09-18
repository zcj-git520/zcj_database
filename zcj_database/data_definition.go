package zcj_database

// 连接数据库的配置文件的结构的定义
// 关系性数据库的配置
// 非关系性数据库的定义

// 定义接口
type DataBase interface {
	Connect() error
	DisConnect() error
}

// sqlite数据库(数据库以文件的形式存储)
type fileDataConfig struct {
	Path string   // 路径
	Name string   // 名称
}

// 数据库
type connDataConfig struct {
	Host  	 string
	Port  	 int
	UserName string
	PassWd 	 string
	DbName   string
	TimeOut  int
}

	// 通用配置文件
type dataBaseConfig struct {
	DataBaseType   string         // 数据库的类型
	FileDataConfig fileDataConfig // 文件型
	ConnDataConfig connDataConfig // 连接型
}


