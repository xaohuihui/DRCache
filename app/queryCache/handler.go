package queryCache

// author: songyanhui
// datetime: 2022/3/7 14:52:37
// software: GoLand

import (
	. "DRCache/common"
	"DRCache/forms"
	"DRCache/global"
	customResponse "DRCache/utils/Response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var r = RedisController{}

func GetQueryCache(c *gin.Context) {
	// 获取参数
	QueryCacheKey := forms.QueryCacheForm{}
	if err := c.ShouldBind(&QueryCacheKey); err != nil {
		HandleValidatorError(c, err)
		return
	}
	res, err := r.GetVal(QueryCacheKey.Key)
	// 本地节点获取不到 通过一致性hash 查询key属于哪个分布式节点，然后进行获取
	remoteNode := global.ConsistentHash.GetNode(QueryCacheKey.Key)
	fmt.Println(remoteNode)
	value, err := CreateRPCClient(remoteNode, QueryCacheKey.Key)
	fmt.Println(value)
	if err != nil {
		customResponse.Err(c, http.StatusBadRequest, 400, "未获取到数据", gin.H{
			"key":   QueryCacheKey.Key,
			"value": nil,
		})
		return
	}
	customResponse.Success(c, http.StatusOK, "获取数据成功", gin.H{
		"key":   QueryCacheKey.Key,
		"value": res,
	})
}

func SetQueryCache(c *gin.Context) {
	// 获取参数
	SetCacheParses := forms.SetCacheParse{}
	if err := c.ShouldBind(&SetCacheParses); err != nil {
		HandleValidatorError(c, err)
		return
	}
	err := r.SetVal(SetCacheParses.Key, SetCacheParses.Value, time.Duration(SetCacheParses.Timeout)*time.Second)
	if err != nil {
		customResponse.Err(c, http.StatusBadRequest, 400, "加入缓存失败", gin.H{
			"key":     SetCacheParses.Key,
			"value":   SetCacheParses.Value,
			"timeout": SetCacheParses.Timeout,
		})
	}
	customResponse.Success(c, http.StatusOK, "加入缓存成功", gin.H{
		"key":     SetCacheParses.Key,
		"timeout": fmt.Sprintf("%vs", SetCacheParses.Timeout),
	})
}
