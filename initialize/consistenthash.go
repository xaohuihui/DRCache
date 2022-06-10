package initialize

import (
	"DRCache/global"
	"DRCache/utils"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

/**
 * @author: xaohuihui
 * @Path: DRCache/initialize/consistenthash.go
 * @Description:
 * @datetime: 2022/4/25 11:08:09
 * software: GoLand
**/

// 单例模型
var consistentHash *utils.Consistent

// GetInstance 获取实例
func GetInstance() *utils.Consistent {
	(&sync.Once{}).Do(func() {
		consistentHash = &utils.Consistent{}
	})
	return consistentHash
}

// InitConsistentHash 初始化一致性hash
func InitConsistentHash(node string, virtualNodeCount int) {
	global.ConsistentHash = GetInstance()
	err := global.ConsistentHash.Add(node, virtualNodeCount)
	if err != nil {
		zap.L().Error(fmt.Sprintf("[InitConsistentHash] 初始化一致性哈希服务失败[%v]", err.Error()))
	}
}
