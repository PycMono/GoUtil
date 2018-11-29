package fileUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestZip2(context *testing.T) {
	startTime := time.Now()
	// zip压缩文件
	dataList = append(dataList, "E:\\中国.sql")
	err := Zip(dataList, "E:\\linux.zip")
	if err != nil {
		print(err)
	}
	// zip解压文件
	err = UnZip("E:\\linux.zip", "E:\\Temp")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Now().Sub(startTime))
}
