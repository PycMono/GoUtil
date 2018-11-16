package mysqlUtil

// mysql连接配置对象
type MysqlConfig struct {
	// 连接字符串
	ConnectionString string

	// 最大开启连接数量
	MaxOpenConns int

	// 最大空闲连接数量
	MaxIdleConns int
}

// 创建新的mysql连接对象
// 参数：
// connectionString：连接字符串
// maxOpenConns：最大开启连接数量
// maxIdleConns：最大空闲连接数量
// 返回值：
//  mysql配置对象
func NewMysqlConfig(connectionString string, maxOpenConns, maxIdleConns int) *MysqlConfig {
	return &MysqlConfig{
		ConnectionString: connectionString,
		MaxIdleConns:     maxIdleConns,
		MaxOpenConns:     maxOpenConns,
	}
}
