package router

import (
	"HttpTest-Server/global"
	"HttpTest-Server/middleware"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() *gin.Engine {
	// gin 模式
	switch global.SysMode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// 路由
	router := gin.New()

	// 中间件
	router.Use(middleware.GinRecovery(true))
	router.Use(middleware.DefaultLogger())

	// 文件目录
	router.Static("/temp", "./temp")

	// HTML 模板资源
	router.Static("/js", "./template/js")
	router.Static("/img", "./template/img")
	router.Static("/fonts", "./template/fonts")
	router.Static("/css", "./template/css")
	router.Static("/assets", "./template/assets")

	// 路由
	r := router.Group("/")
	{
		// 默认提示信息
		Default(r, "/")
		// 系统信息接口
		Sys(r, "/sys")
		// 面板管理接口
		Panel(r, "/panel")
		// 测试接口
		Test(r, "/test")
	}

	return router
}
