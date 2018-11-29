package excelUtil

import (
	"fmt"
	"path"

	"regexp"

	"strings"

	"unicode"

	"github.com/tealeg/xlsx"
)

type ExcelUtil struct {
	// 文件路径
	path string

	// xlsx文件对象
	file *xlsx.File
}

// Excel成字典的形式(每个cell的奇数作为key，偶数作为value)
// 参数：无
// 返回值：
// 字典集合
func (this *ExcelUtil) ReadExcelToDict() map[string]string {
	resultDict := make(map[string]string)
	for _, sheet := range this.file.Sheets {
		for _, row := range sheet.Rows {
			i := 0
			key := ""
			for _, cell := range row.Cells {
				i++
				text := cell.String()
				// 每个cell的奇数作为key，偶数作为value
				if i%2 == 0 {
					resultDict[key] = text
				} else {
					key = text
				}
			}
		}
	}

	return resultDict
}

func (this *ExcelUtil) Replace(resultDict map[string]string, tarFilePath string) {
	for _, sheet := range this.file.Sheets {
		for _, row := range sheet.Rows {
			if row.Height == 0 { // 表头不替换
				continue
			}

			for _, cell := range row.Cells {
				text := cell.String()
				if !this.isChinese(text) {
					continue
				}

				if value, ok := resultDict[text]; ok {
					cell.Value = value
				} else {
					fmt.Println(fmt.Sprintf("resultDict不存在key=%s 的日文翻译sheet.Name=%s", text, sheet.Name))
				}
			}
		}
	}

	splitList := strings.Split(this.path, "\\")
	err := this.file.Save(path.Join(tarFilePath, splitList[len(splitList)-1]))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 判断是否包含中文字符串
func (this *ExcelUtil) isChinese(word string) bool {
	for _, r := range word {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("^[\u4e00-\u9fa5]{3,8}$").MatchString(string(r))) {
			return true
		}
	}

	return false
}

// 创建新的ExcelUtil文件助手
// 参数：
// path：Excel文件路径
// 返回值：
// 1.助手对象
// 2.错误对象
func NewExcelUtil(path string) (*ExcelUtil, error) {
	file, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &ExcelUtil{
		file: file,
		path: path,
	}, nil
}
