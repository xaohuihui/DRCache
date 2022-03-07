package initialize

import (
	"DRCache/middlewares"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options =  []Option{}

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// InitRouters 初始化路由
func InitRouters() *gin.Engine {
	r := gin.Default()

	// 在rooter中加入GinLogger, GinRecovery中间件
	r.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	// 使用跨域中间件
	r.Use(middlewares.Cors())

	for _, opt := range options {
		opt(r)
	}
	return r
}
