package BLL

import (
	"moqikaka.com/Test/src/Model"
	"moqikaka.com/Test/src/DAL"
	"fmt"
	"moqikaka.com/goutil/logUtil"
)

// 初始化数据
func InitData(){
	rows,err := DAL.GetList("SELECT * FROM p_harem_info")
	if err != nil{
		logUtil.DebugLog("不存在数据")
		return
	}

	list:=make([]*Model.HaremInfo,0)
	for rows.Next(){
		haremInfo:=Model.NewHaremInfo()
		err = rows.Scan(&haremInfo.PlayerID,&haremInfo.Crdate,&haremInfo.OpenHaremNeedGoodsID,&haremInfo.NextShopRefreshTime,&haremInfo.ShopRefreshNum)
		if err != nil{
			fmt.Println(err)
			return
		}

		list=append(list,haremInfo)
	}

	//for _,haremInfo:=range list{
	//	fmt.Println(fmt.Sprintf("PlayerID=%s,Crdate=%v,OpenHaremNeedGoodsID=%s,NextShopRefreshTime=%v,ShopRefreshNum=%d",haremInfo.PlayerID,haremInfo.Crdate,haremInfo.OpenHaremNeedGoodsID,haremInfo.NextShopRefreshTime,haremInfo.ShopRefreshNum))
	//}

	if len(list)<=0{
		logUtil.DebugLog("不存在数据")
	}

	rows.Close()
}
