package DAL

import (
	"database/sql"
	"moqikaka.com/Test/src/DB"
)

// 获取信息
func GetList(sql string)(rows *sql.Rows,err error) {
	db:=DBTool.GetDB()
	rows,err= db.Query(sql)

	return
}
