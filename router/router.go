package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 连接数据库
	r.GET("/index", service.GetIndex)
	return r
}
