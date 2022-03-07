package queryCache

// author: songyanhui
// datetime: 2022/3/7 14:52:47
// software: GoLand

import (
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	queryRouter := r.Group("/query")
	{
		queryRouter.GET("/get", GetQueryCache)
		queryRouter.POST("/set", SetQueryCache)
	}
}
