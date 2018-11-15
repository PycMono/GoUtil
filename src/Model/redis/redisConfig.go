package redis

import "time"

// redis配置文件结构体
type RedisConfig struct {
	// 连接字符串
	ConnectionString string

	// 密码
	Password string

	// 数据库编号
	Database int

	// 最大活跃连接数
	MaxActive int

	// 最大空闲连接数
	MaxIdle int

	// 空闲超时
	IdleTimeout time.Duration

	// 连接超时
	DialConnectTimeout time.Duration
}
