package fileUtil

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	mutexs     sync.Mutex
	dataList   = make([]string, 0)
	wg         sync.WaitGroup
	InitString = `{"Name":"张三","Age":"30","Data":"我爱中国，中国爱我"}`
	ZlibBytes  []byte
)

func TestGetFileList(t *testing.T) {
	// 获取当前路径
	tempList, err := GetFileList("E:\\Project\\WorkingDirectory\\GoProject\\src\\moqikaka.com\\Test")
	if err == nil {
		fmt.Println(tempList)
	} else {
		fmt.Println(err)
	}
}

func TestReadOrWrite(t *testing.T) {
	ReadFileByIOUtil("D:\\", "20180828_slg.sql")
	tempList, err := ReadFileByIOUtil("D:\\", "20180828_slg.sql")
	if err != nil {
		fmt.Println(err)
	}

	mutexs.Lock()
	s := string(tempList)
	dataList = append(dataList, s)
	mutexs.Unlock()

	_, err = ReadFileByBufferIO("D:\\", "20180828_slg.sql")
	if err != nil {
		fmt.Println(err)
	}

	ReadFileByBufferIO("D:\\", "20180828_slg.sql")
}

func run() {
	filePath := "D:\\tt"
	fileName := "hah"
	for {
		// 先休眠，避免系统启动时就进行报警
		time.Sleep(time.Hour * time.Duration(10))
		mutexs.Lock()
		tempDataList := make([]string, 0)
		tempDataList = append(tempDataList, dataList...)
		dataList = make([]string, 0)
		mutexs.Unlock()
		fmt.Println(len(tempDataList))
		err := WriteFile(filePath, fileName, tempDataList)
		if err != nil {
			fmt.Println(err)
		}
	}
}
