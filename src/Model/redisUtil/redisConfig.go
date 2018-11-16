package redisUtil

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

// 创建redis配置文件
// 参数：
// connectionString：连接字符串
// password：密码
// database：数据库编号
// maxActive：最大活跃连接数
// maxIdle：最大空闲连接数
// idleTimeout：空闲超时
// dialConnectTimeout：连接超时
func NewRedisConfig(connectionString, password string, database, maxActive, maxIdle int, idleTimeout, dialConnectTimeout time.Duration) *RedisConfig {
	return &RedisConfig{
		ConnectionString:   connectionString,
		Password:           password,
		Database:           database,
		MaxActive:          maxActive,
		MaxIdle:            maxIdle,
		IdleTimeout:        idleTimeout,
		DialConnectTimeout: dialConnectTimeout,
	}
}
