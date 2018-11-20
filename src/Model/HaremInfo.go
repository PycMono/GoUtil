package Model

import "time"

type HaremInfo struct {
	// 玩家ID
	PlayerID string

	// 创建时间
	Crdate time.Time

	// 开启后宫所需道具ID集合
	OpenHaremNeedGoodsID string

	// 下一次商城刷新时间
	NextShopRefreshTime time.Time

	// 商店刷新次数
	ShopRefreshNum int
}

// 创建新的后宫对象
func NewHaremInfo() *HaremInfo {
	return &HaremInfo{}
}
