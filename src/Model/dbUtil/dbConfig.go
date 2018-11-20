package dbUtil

// DB连接配置对象
type DBConfig struct {
	// 连接字符串
	ConnectionString string

	// 连接类型(mysql:mysql库连接...)
	ConnectionType string

	// 最大开启连接数量
	MaxOpenConns int

	// 最大空闲连接数量
	MaxIdleConns int
}

// 创建新的mysql连接对象
// 参数：
// connectionString：连接字符串
// connectionType：连接类型(mysql:mysql库连接...)
// maxOpenConns：最大开启连接数量
// maxIdleConns：最大空闲连接数量
// 返回值：
//  mysql配置对象
func NewDBConfig(connectionString, connectionType string, maxOpenConns, maxIdleConns int) *DBConfig {
	return &DBConfig{
		ConnectionString: connectionString,
		MaxIdleConns:     maxIdleConns,
		MaxOpenConns:     maxOpenConns,
		ConnectionType:   connectionType,
	}
}
