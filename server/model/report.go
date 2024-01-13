package model

import (
	"HttpTest-Server/global"
	"go.uber.org/zap"
)

// Report 测试报告表
type Report struct {
	Base
	// 数据ID
	UID string `json:"uid" gorm:"column:uid;primaryKey"`
	// 自动化测试数据表 UID | HT_test_data --> uid
	TestUid string `json:"test_uid" gorm:"column:test_uid"`
	// 状态
	Status int8 `json:"status" gorm:"column:status"`
	// 请求数
	Total int64 `json:"total" gorm:"column:total"`
	// 成功数
	Success int64 `json:"success" gorm:"column:success"`
	// 失败数
	Fail int64 `json:"fail" gorm:"column:fail"`
	// 耗时 | ms
	SpeedTime int64 `json:"speed_time" gorm:"column:speed_time"`
}

// ReportApp report 模型实例
var ReportApp = new(Report)

// New 初始化模型
func (i *Report) New(data Report) *Report {
	return &data
}

// Get 获取信息
func (i *Report) Get() (data []Report, err error) {
	//err = global.APP_DB.Model(&i).Where("test_uid = ?", i.TestUid).Find(&data).Error
	err = global.APP_DB.Model(&i).Where(&i).Find(&data).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 测试报告表#获取信息失败", zap.Error(err))
		return []Report{}, err
	}
	return
}

// Delete 删除信息
func (i *Report) Delete() (err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Unscoped().Delete(&i).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 测试报告表#删除信息失败", zap.Error(err))
	}
	return
}
