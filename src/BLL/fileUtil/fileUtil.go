package fileUtil

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
)

// 文件读取助手类
type FileUtil struct {
}

// 判断文件夹是否存在
// 参数：
// filePath：文件夹路径
// 返回值：
// 1.true：存在，反之不存在
func DirExists(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return f.IsDir()
}

// 判断文件是否存在
func FileExists(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	return f.IsDir() == false
}

// 获取程序运行路径
// 参数：
// 无
// 返回值：
// 1.路径
// 2.错误对象
func GetCurrentPath() (string, error) {
	path, _ := exec.LookPath(os.Args[0])
	fileAbsPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	i := strings.LastIndex(fileAbsPath, "/")
	if i < 0 {
		i = strings.LastIndex(fileAbsPath, "\\")
	}

	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}

	return string(fileAbsPath[0 : i+1]), nil
}

// 获取当前目录下面的所有文件(也可以写递归调用)
// 参数：
// path：目录
// 返回值：
// 1.文件集合
// 2.错误对象
func GetFileList(path string) ([]string, error) {
	// 判断文件夹是否存在
	exists := DirExists(path)
	if !exists {
		return nil, errors.New(fmt.Sprintf("path=%s文件不存在", path))
	}

	resultList := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		resultList = append(resultList, f.Name())
		return nil
	})

	return resultList, err
}

// 读取文件
// 参数：
// filePath：文件夹路径
// fileName：文件名字
// 返回值：
// 1.读取出来的数据
// 2.错误对象
func ReadFileByIOUtil(filePath, fileName string) ([]byte, error) {
	startTime := time.Now()
	fileName = filepath.Join(filePath, fileName)
	os, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func() {
		os.Close()
		fmt.Printf("[ReadFileByIOUtil] cost time %v \n", time.Now().Sub(startTime))
	}()

	// 读取数据
	tempList, err := ioutil.ReadAll(os)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tempList, nil
}

// 读取文件
// 参数：
// filePath：文件夹路径
// fileName：文件名字
// 返回值：
// 1.读取出来的数据
// 2.错误对象
func ReadFileByBufferIO(filePath, fileName string) ([]byte, error) {
	fileName = filepath.Join(filePath, fileName)
	startTime := time.Now()
	os, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func() {
		os.Close()
		fmt.Printf("[ReadFileByBufferIO] cost time %v \n", time.Now().Sub(startTime))
	}()

	r := bufio.NewReader(os)
	var resultData []byte

	readBuf := make([]byte, 1024)
	for {
		n, err := r.Read(readBuf)
		if err != nil && err != io.EOF {
			// 文件读取失败，记录日志，直接退出
			fmt.Println(err)
			return nil, err
		}
		if 0 == n { // 文件读取完成，退出循环
			break
		}

		resultData = append(resultData, readBuf[:n]...)
	}

	return resultData, nil
}

// 写文件
// 参数：
// filePath：文件夹路径
// fileName：文件名字
// content：类容字符串
// 返回值：
// 1.错误对象
func WriteFile(filePath, fileName string, content []string) error {
	fileName = filepath.Join(filePath, fileName)

	// 如果文件夹不存在就创建文件夹先
	mutex.Lock()
	if !DirExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
		if err != nil {
			fmt.Println(err)
		}
	}

	os, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm|os.ModeTemporary)
	if err != nil {
		fmt.Println(fmt.Printf("报错了%s", err))
		return err
	}

	defer func() {
		os.Close()
		mutex.Unlock()
	}()

	// 写入文件
	for _, item := range content {
		os.WriteString(item)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// 写文件
// 参数：
// filePath：文件夹路径
// fileName：文件名字
// content：类容字符串
// 返回值：
// 1.错误对象
func WriteFileByByte(filePath, fileName string, content []byte) error {
	fileName = filepath.Join(filePath, fileName)

	// 如果文件夹不存在就创建文件夹先
	mutex.Lock()
	if !DirExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm|os.ModeTemporary)
		if err != nil {
			fmt.Println(err)
		}
	}

	os, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm|os.ModeTemporary)
	if err != nil {
		fmt.Println(fmt.Printf("报错了%s", err))
		return err
	}

	defer func() {
		os.Close()
		mutex.Unlock()
	}()

	os.Write(content)

	return nil
}
