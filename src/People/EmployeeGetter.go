package People

// 定义员工接口
// EmployeeGetter 接口中嵌入了 PeopleGetter 接口，前者将获取后者的所有方法
type EmployeeGetter interface {
	// 人接口
	PeopleGetter

	// 获取薪水
	GetSalary() int

	// 员工ID号
	GetID()
}