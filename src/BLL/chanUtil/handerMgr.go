package chanUtil

var (
	initSuccessObj = NewChanUtil("handerMgr")
)

// 注册初始化成功的通道
// name:模块名称
// ch:通道对象
func RegisterInitSuccess(name string, ch chan bool) {
	initSuccessObj.Register(name, ch)
}

// 通知所有已经注册的通道对象
func Notify() {
	initSuccessObj.Notify()
}
