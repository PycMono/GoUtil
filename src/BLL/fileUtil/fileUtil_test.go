package fileUtil

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	mutexs   sync.Mutex
	dataList = make([]string, 0)
	wg       sync.WaitGroup
)

func TestReflect(context *testing.T) {
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

	//fmt.Println("-----------------------------")
	//startTime = time.Now()
	//tempList, err := ReadFileByIOUtil("E:\\", "20180828_slg.sql")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//tempList1, err := Zip2(tempList, zlib.DefaultCompression)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//WriteFileByByte("E:\\", "temp", tempList1)

	//// 设置WaitGroup需要等待的数量，只要有一个服务器出现错误都停止服务器
	//wg.Add(1)
	//
	////curPath, _ := GetCurrentPath()
	//// 获取当前路径
	////fmt.Println(curPath)
	//tempList, err := GetFileList("E:\\Project\\WorkingDirectory\\GoProject\\src\\moqikaka.com\\Test")
	//if err == nil {
	//	fmt.Println(tempList)
	//} else {
	//	fmt.Println(err)
	//}

	//go ReadFileByIOUtil("D:\\", "20180828_slg.sql")
	//tempList, err := ReadFileByIOUtil("D:\\", "20180828_slg.sql")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//mutexs.Lock()
	//s := string(tempList)
	//dataList = append(dataList,s)
	//mutexs.Unlock()
	//
	//_, err = ReadFileByBufferIO("D:\\", "20180828_slg.sql")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//go ReadFileByBufferIO("D:\\", "20180828_slg.sql")
	//

	//s := string(tempList)
	//fmt.Println(s)
	//go run()
	//go run()
	//go run()
	//for i := 0; i < 100000000; i++ {
	//	mutexs.Lock()
	//	dataList = append(dataList, "222")
	//	mutexs.Unlock()
	//}

	//wg.Wait()
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
