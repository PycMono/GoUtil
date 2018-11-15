package People

// 获取人的信息
type PeopleGetter interface {
	// 获取名字
	GetName() string

	// 获取年龄
	GetAge() string
}
