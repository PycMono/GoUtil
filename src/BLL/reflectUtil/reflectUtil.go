package reflectUtil

import (
	"fmt"
	"moqikaka.com/Test/src/DB"
	"reflect"
	"strconv"
	"time"
)

// 反射工具对象
type ReflectUtil struct {
}

func (this *ReflectUtil) CreateInstance(v interface{}, tableName string) ([]interface{}, error) {
	// 获取数据库中的数据
	tempList, err := this.queryDB(tableName)
	if err != nil {
		return nil, err
	}

	resultList := make([]interface{}, 0)
	newReflect := reflect.ValueOf(v)
	if newReflect.Kind() == reflect.Ptr {
		newReflect = newReflect.Elem()
	}

	// 遍历获取每行数据
	for _, tempDict := range tempList {
		for column, value := range tempDict {
			fieldInfo := newReflect.FieldByName(column)
			if !fieldInfo.IsValid() || !fieldInfo.CanSet() {
				fmt.Println(fmt.Sprintf("字段不可设置column=%s", column))
			}

			// 数据转换，并且设置字段值
			fieldInfo.Set(this.convertType(fieldInfo, value))
		}

		resultList = append(resultList, newReflect.Interface())
	}

	return resultList, nil
}

// 查询数据库
// 参数：
// tableName：表名字
// 返回值：
// 数据库中的所有数据集合
// 错误对象
func (this *ReflectUtil) queryDB(tableName string) ([]map[string]string, error) {
	db := DBTool.GetDB()
	defer db.Close()

	sql := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	params := make([][]byte, len(columns))
	scans := make([]interface{}, len(params))
	for i := range params {
		scans[i] = &params[i]
	}

	// 保存最终结果到切片中
	resultList := make([]map[string]string, 0)
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			return nil, err
		}

		// 单条数据保存到字典中
		singleDict := make(map[string]string)
		for k, v := range params {
			key := columns[k]
			singleDict[key] = string(v)
		}

		resultList = append(resultList, singleDict)
	}

	return resultList, nil
}

// 类型转换
// 参数：
// fieldInfo：反射的字段信息
// value：待转换的值
// 返回值：
// 转换后的值
func (this *ReflectUtil) convertType(fieldInfo reflect.Value, value string) reflect.Value {
	var setValue reflect.Value
	switch fieldInfo.Interface().(type) {
	case float32:
		setValue = reflect.ValueOf(value)
	case int:
		atoiInt, err := strconv.Atoi(value)
		if err == nil {
			setValue = reflect.ValueOf(atoiInt)
		}
	case string:
		setValue = reflect.ValueOf(value)
	case time.Time:
		time, err := time.Parse("2006-01-02T15:04:05Z07:00", value)
		if err == nil {
			setValue = reflect.ValueOf(time)
		}

		// 后面自己加
	}

	return setValue
}

// 创造新的反射工具对象
// 参数：
// 返回值：
// 反射对象
func NewReflectUtil() *ReflectUtil {
	return &ReflectUtil{}
}
