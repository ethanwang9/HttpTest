package router

import (
	"HttpTest-Server/controller/sys"
	"github.com/gin-gonic/gin"
)

func Sys(group *gin.RouterGroup, path string) {
	r := group.Group(path)
	{
		// 获取面板基本信息
		r.GET("/basic", sys.GetBasic)
	}
}
