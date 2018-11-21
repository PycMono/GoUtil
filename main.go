package main

import (
	_ "github.com/go-sql-driver/mysql"
	"moqikaka.com/Test/src/BLL"
	"sync"
	"moqikaka.com/goutil/logUtil"
	."moqikaka.com/Test/src/TestInterface/Impl"
	"moqikaka.com/Test/src/Func"
	"fmt"
	"os/exec"
	"os"
	"path/filepath"
)

var (
	wg sync.WaitGroup
	con_SEPERATOR ="------------------------------------------------------------------------------"
)

func init() {
	// 设置WaitGroup需要等待的数量，只要有一个服务器出现错误都停止服务器
	wg.Add(1)

	// 设置日志文件的存储目录
	logUtil.SetLogPath("LOG")
}

func main()  {
	BLL.InitData()

	// 设计模式
	newPeoPlesImpl:=NewPeoplesImpl("张三",1)
	newPeoPlesImpl.CaclFight(newPeoPlesImpl)

	fmt.Println(con_SEPERATOR)

	// 匿名函数
	Func.NewAnimal("猫","喜欢吃鱼").PrintDetail()

	fmt.Println(con_SEPERATOR)

	wg.Wait()
}
