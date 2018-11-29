package chanUtil

import (
	"fmt"
	"sync"
)

// 通道工具类
type ChanUtil struct {
	// 对象名字
	name string

	// chan集合
	registerMap map[string]chan bool

	// 锁对象
	mutex sync.RWMutex
}

// 注册被通知的对象
// 参数：
// name：唯一标识
// ch：通知的通道对象
func (this *ChanUtil) Register(name string, ch chan bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	if _, exits := this.registerMap[name]; exits {
		panic(fmt.Sprintf("ChanUtil.Register-%s-%s通知对象已经存在了", this.name, name))
	}

	this.registerMap[name] = ch
	fmt.Println(fmt.Sprintf("ChanUtil.Register成功注册了-%s-%s通知对象,并且当前存在%d对象", this.name, name, len(this.registerMap)))
}

// 取消已经注册的通知对象
// 参数：
// // name：唯一标识
func (this *ChanUtil) UnRegister(name string) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	delete(this.registerMap, name)
	fmt.Println(fmt.Sprintf("ChanUtil.UnRegister成功删除-%s-%s通知对象,并且当前存在%d对象", this.name, name, len(this.registerMap)))
}

// 通知所有已经注册的通道对象
func (this *ChanUtil) Notify() {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	for name, ch := range this.registerMap {
		ch <- true
		fmt.Println(fmt.Sprintf("ChanUtil.Notify通知消息成功--%s--%s", this.name, name))
	}
}

// 创建新的工具对象
// 参数：
// name：对象名字
func NewChanUtil(name string) *ChanUtil {
	return &ChanUtil{
		name:        name,
		registerMap: map[string]chan bool{},
	}
}
