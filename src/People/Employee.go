package People

// 定义员工结构体
type Employee struct {
	// 名字
	Name string

	// 年龄
	Age int32

	// 薪水
	Salary int32

	// 玩家ID号
	ID int32
}

// 获取玩家名字
func (this *Employee)GetName()string  {
	return  this.Name
}

// 获取玩家年龄
func (this *Employee)GetAge()int32  {
	return  this.Age
}

// 获取玩家ID号
func (this *Employee)GetID()int32  {
	return  this.ID
}

// 获取薪水
func  (this *Employee)GetSalary()int32  {
	return  this.GetSalary()
}
