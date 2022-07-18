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
	var err error
	QueryCacheKey := forms.QueryCacheForm{}
	if err := c.ShouldBind(&QueryCacheKey); err != nil {
		HandleValidatorError(c, err)
		return
	}
	res, err := r.GetVal(QueryCacheKey.Key)
	// 本地节点获取不到 通过一致性hash 查询key属于哪个分布式节点，然后进行获取
	flag := 0
	if err != nil {
		remoteNode := global.ConsistentHash.GetNode(QueryCacheKey.Key)
		data, err := global.SingleGroup.Do(QueryCacheKey.Key, func() (interface{}, error) {
			data, err := CreateGetRPCClient(remoteNode, QueryCacheKey.Key)
			if err != nil {
				return nil, err
			}
			global.Lg.Info(fmt.Sprintf("获取远程节点%s缓存key%s\n", remoteNode, QueryCacheKey.Key))
			return data, nil
		})
		if err != nil {
			customResponse.Err(c, http.StatusBadRequest, 400, "未获取到数据", gin.H{
				"key":   QueryCacheKey.Key,
				"value": nil,
			})
			return
		}
		flag = 1
		res = data.([]byte)
	}
	if err != nil {
		customResponse.Err(c, http.StatusBadRequest, 400, "未获取到数据", gin.H{
			"key":   QueryCacheKey.Key,
			"value": nil,
		})
		return
	}

	// 若为远程节点获取的值，则缓存该值，过期时间设置为较短
	if flag == 1 {
		err := r.SetVal(QueryCacheKey.Key, string(res), time.Duration(100)*time.Second)
		if err != nil {
			global.Lg.Info(fmt.Sprintf("缓存数据失败key: %s, value : %s\n", QueryCacheKey.Key, string(res)))
		}
	}
	customResponse.Success(c, http.StatusOK, "获取数据成功", gin.H{
		"key":   QueryCacheKey.Key,
		"value": string(res),
	})
}

func SetQueryCache(c *gin.Context) {
	// 获取参数
	SetCacheParses := forms.SetCacheParse{}
	if err := c.ShouldBind(&SetCacheParses); err != nil {
		HandleValidatorError(c, err)
		return
	}
	// string 类型转byte数组
	remoteNode := global.ConsistentHash.GetNode(SetCacheParses.Key)
	err := CreateSetRPCClient(remoteNode, SetCacheParses.Key, []byte(SetCacheParses.Value), SetCacheParses.Timeout)
	global.Lg.Info(fmt.Sprintf("设置远程节点[%s]缓存key[%s]\n", remoteNode, SetCacheParses.Key))
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
