package Impl

import (
	"fmt"
	."moqikaka.com/Test/src/TestInterface"
)

// 人结构体
type PeoplesImpl struct {
	// 名字
	Name string

	// 年龄
	Age int32

	// 消息父类
	*MessageBase
}

// 实现接口
func (this *PeoplesImpl)PrintName()  {
	fmt.Println(fmt.Sprintf("Name = %s,Age = %d",this.Name,this.Age))
}

// 创建新的结构体对象
// 参数：
// name：名字
// age：年龄
func NewPeoplesImpl(name string,age int32)*PeoplesImpl{
	return &PeoplesImpl{
		Name:name,
		Age:age,
		MessageBase:NewMessageBase(),
	}
}