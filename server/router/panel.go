package router

import (
	"HttpTest-Server/controller/panel"
	"github.com/gin-gonic/gin"
)

func Panel(group *gin.RouterGroup, path string) {
	r := group.Group(path)
	{
		// ==========
		// 接口
		// ==========
		// 获取 - 文件树
		r.GET("/interface/dir", panel.GetInterfaceDir)
		// 创建/更新 - 文件树
		r.POST("/interface/dir", panel.SetInterfaceDir)
		// 创建 - 接口
		r.POST("/interface/info", panel.SetInterfaceInfo)
		// 获取 - 接口
		r.GET("/interface/info", panel.GetInterfaceInfo)
		// 更新 - 接口
		r.PUT("/interface/info", panel.UpdateInterfaceInfo)
		// 删除 - 接口
		r.DELETE("/interface/info", panel.DelInterfaceInfo)
		// ==========
		// 自动化测试
		// ==========
		// 获取 - 文件树
		r.GET("/test/dir", panel.GetTestDir)
		// 创建/更新 - 文件树
		r.POST("/test/dir", panel.SetTestDir)
		// 创建 - 场景
		r.POST("/test/info", panel.SetTestInfo)
		// 删除 - 场景
		r.DELETE("/test/info", panel.DelTestInfo)
		// 获取 - 场景
		r.GET("/test/info", panel.GetTestInfo)
		// 更新 - 场景
		r.PUT("/test/info", panel.UpdateTestInfo)
		// 获取 - 测试报告
		r.GET("/test/report", panel.GetTestReport)
		// 删除 - 测试报告
		r.DELETE("/test/report", panel.DelTestReport)
		// 获取 - 测试报告详细信息
		r.GET("/test/report/details", panel.GetTestReportDetails)
		// 获取 - 测试接口请求信息
		r.GET("/test/report/request", panel.GetTestReportRequest)
		// ==========
		// 系统设置
		// ==========
		// 创建/更新 - 系统代理
		r.POST("/system/proxy", panel.SetSystemProxy)
		// 获取 - 系统代理
		r.GET("/system/proxy", panel.GetSystemProxy)
		// 获取 - 系统版本信息
		r.GET("/system/update", panel.GetSystemUpdate)
	}
}
