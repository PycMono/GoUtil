package Func

import "fmt"

type Animal struct {
	// 名字
	Name string

	// 爱好
	Hobby string
}

// 打印详细信息
func (this *Animal)PrintDetail()  {
	func(){
		fmt.Println(fmt.Sprintf("名字：%s",this.Name))
	}()// 直接调用

	a:=0
	ageFun:= func(hobby string) {
		fmt.Println(fmt.Sprintf("爱好：%s",hobby))
		a:=3
		fmt.Println(a)
	}
	fmt.Println(a)
	// 匿名方法调用
	ageFun(this.Hobby)
}

// 创建新的动物实例
// 参数：
// name：名字
// hobby:爱好
// 返回值：
// Animal实例
func NewAnimal(name,hobby string)*Animal  {
	return &Animal{
		Name:name,
		Hobby:hobby,
	}
}
