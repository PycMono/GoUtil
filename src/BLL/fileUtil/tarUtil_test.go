package fileUtil

import (
	"fmt"
	"testing"
	"time"
)

func TestTar(t *testing.T) {
	startTime := time.Now()
	// tar压缩文件
	dataList = append(dataList, "E:\\20180828_slg.sql")
	err := Tar(dataList, "E:\\tube.tar.gz")
	if err != nil {
		print(err)
	}

	err = UnTar("E:\\tube.tar.gz", "E:\\Temp")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(time.Now().Sub(startTime))
}
