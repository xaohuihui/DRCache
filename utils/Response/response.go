package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, code int, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}

func Err(c *gin.Context, httpCode int, code int, msg string, jsonStr interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": jsonStr,
	})
	return
}
