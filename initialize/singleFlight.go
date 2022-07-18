package initialize

import (
	"DRCache/global"
	"DRCache/utils"
)

func InitSingleFlight() {
	// 生成redis客户端
	global.SingleGroup = &utils.Group{}

	global.Lg.Error("[InitSingleFlight] 初始化SingleFlight")

}
