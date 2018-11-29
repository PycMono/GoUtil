package stringUtil

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串切割成字典
// 参数：
// str：待处理的字符串
// 返回值：
// 1.切割后的数据
// 2.错误对象
func SplitToDict(str string) (map[int]int, error) {
	resultDict := make(map[int]int)
	if str == "" {
		return resultDict, nil
	}

	// 按照我预先制定的规则切割字符串"|",","
	firstSplitList := strings.Split(str, "|")
	for _, item := range firstSplitList {
		secondSplitList := strings.Split(item, ",")
		// 转换成整型
		key, err := strconv.Atoi(secondSplitList[0])
		if err != nil {
			return nil, err
		}

		value, err := strconv.Atoi(secondSplitList[1])
		if err != nil {
			return nil, err
		}

		if _, exists := resultDict[key]; exists {
			// 如果存在就累加
			resultDict[key] = resultDict[key] + value
		} else {
			resultDict[key] = value
		}
	}

	return resultDict, nil
}

// 字符串切割成List
// 参数：
// str：待处理的字符串
// 1.切割后的数据
// 2.错误对象
func SplitToList(str string) ([]int, error) {
	resultList := make([]int, 0)
	if str == "" {
		return resultList, nil
	}

	// 按照我预先制定的规则切割字符串","
	firstSplitList := strings.Split(str, ",")
	for _, item := range firstSplitList {
		// 转换成整型
		value, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}

		resultList = append(resultList, value)
	}

	return resultList, nil
}

// 追加字符串
// 参数：
// sourceStr：源字符串
// format：格式符号字符串
// appendStr：待追加的字符串
// 返回值：
// 1.追加后的字符串(需要去掉末尾的格式符号)
func Append(sourceStr, format, appendStr string) string {
	targetStr := fmt.Sprintf("%s%s%s", sourceStr, format, appendStr)
	targetStr = strings.Trim(sourceStr, format)

	return targetStr
}
