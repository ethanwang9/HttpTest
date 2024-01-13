package initialization

import (
	"HttpTest-Server/global"
	"HttpTest-Server/model"
	"go.uber.org/zap"
)

// GormInit 数据库初始化
func gormInit() {
	// 判断系统是否初始化
	result, err := model.SystemApp.Get()
	if err == nil && result.SID == 1 {
		global.APP_LOG.Info("[初始化] 数据库已初始化，已跳过初始化操作！")
		return
	}

	// 初始化数据库
	// 1. 初始化表
	gormInitTable()
	// 2. 初始化表记录
	gormInitData()

	global.APP_LOG.Info("[初始化] 数据库已完成初始化.")
}

// GormInitTable 初始化表
func gormInitTable() {
	_ = global.APP_DB.AutoMigrate(
		// 系统信息表
		&model.System{},
		// 接口管理数据表
		&model.InterfaceData{},
		// 自动化测试数据表
		&model.TestData{},
		// 测试报告表
		&model.Report{},
		// 测试报告数据表
		&model.ReportData{},
	)

}

// GormInitData 初始化表记录
func gormInitData() {
	// 初始化表 system
	err := model.SystemApp.New(model.System{
		GitUrl: "https://github.com/ethanwang9/HttpTest",
		DocUrl: "https://apifox.com/apidoc/shared-2f76c614-6934-4539-8c7a-186f9f77705f",
	}).Set()
	if err != nil {
		global.APP_LOG.Error("[数据库] 初始化 system 表数据失败", zap.Error(err))
	}
}
