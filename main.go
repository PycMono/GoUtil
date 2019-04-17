package main

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"

	//"moqikaka.com/Test/src/BLL"
	//"moqikaka.com/Test/src/TempClass"
	//. "moqikaka.com/Test/src/TestInterface/Impl"
	"fmt"

	"moqikaka.com/goutil/logUtil"
)

var (
	wg            sync.WaitGroup
	con_SEPERATOR = "------------------------------------------------------------------------------"
)

func init() {
	// 设置WaitGroup需要等待的数量，只要有一个服务器出现错误都停止服务器
	//wg.Add(1)

	// 设置日志文件的存储目录
	logUtil.SetLogPath("LOG")
}

func main() {
	fmt.Println("hello my go process!!")

	//BLL.InitData()

	//TempClass.NewManImpl().TT()
	//// 设计模式
	//newPeoPlesImpl := NewPeoplesImpl("张三", 1)
	//newPeoPlesImpl.CaclFight(newPeoPlesImpl)
	//
	//fmt.Println(con_SEPERATOR)
	//
	//// 匿名函数
	//Func.NewAnimal("猫", "喜欢吃鱼").PrintDetail()
	//
	//fmt.Println(con_SEPERATOR)
	//ch := make(chan int)
	//go func() {
	//	time.Sleep(10 * time.Second)
	//	ch <- 1
	//}()
	//
	//for {
	//	select {
	//	case _ = <-ch:
	//		fmt.Println("妈的，智障")
	//		break
	//	default:
	//		// 如果channel中没有数据，则休眠5秒
	//		time.Sleep(5 * time.Second)
	//	}
	//}

	//	wg.Wait()
}
