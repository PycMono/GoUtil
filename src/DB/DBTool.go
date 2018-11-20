package DBTool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"moqikaka.com/goutil/logUtil"
	"strings"
)

// Db数据库连接池
var DB *sql.DB

// 注意方法名大写，就是public
func init() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{"root", ":", "moqikaka3309", "@tcp(", "10.1.0.30", ":", "3309", ")/", "player_feature_fight_report", "?charset=utf8&loc=Local&parseTime=true"}, "")

	// 打开数据库,前者是驱动名，所以要导入
	DB, _ = sql.Open("mysql", path)
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		panic("opon database fail")
	}

	logUtil.DebugLog("connnect success")
}

// 返回数据库连接对象
func GetDB() *sql.DB {
	return DB
}
