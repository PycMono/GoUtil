package TestInterface

// 父类信息类
type MessageBase struct {
	// 信息接口
	IMessage
}

// 计算战斗校验
func (this *MessageBase) CaclFight(iMessage  IMessage)  {
	this.setMessage(iMessage)
	this.PrintName()
	println()
}

// 打印数据
func (this *MessageBase)Print(){
	this.IMessage.PrintName()
}

// 设置信息
// 参数：
// iMessage：接口信息
func (this *MessageBase)setMessage(iMessage  IMessage){
	this.IMessage=iMessage
}

// 创建新的MessageBase类
// 返回值：
// MessageBase：MessageBase信息对象
func NewMessageBase()*MessageBase{
	return &MessageBase{

	}
}

