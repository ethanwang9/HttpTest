package model

import (
	"HttpTest-Server/global"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

// ReportData 测试报告数据表
type ReportData struct {
	Base
	// 数据ID
	UID string `json:"uid" gorm:"column:uid;primaryKey"`
	// 测试报告表 UID | HT_report --> uid
	ReportUid string `json:"report_uid" gorm:"column:report_uid"`
	// 请求状态
	Status int8 `json:"status" gorm:"column:status"`
	// HTTP请求类型
	Method string `json:"method" gorm:"column:method"`
	// HTTP请求地址
	Url string `json:"url" gorm:"column:url"`
	// 接口名称
	Name string `json:"name" gorm:"column:name"`
	// 请求状态码
	StatusCode int `json:"status_code" gorm:"column:status_code"`
	// 耗时 ms
	SpeedTime int64 `json:"speed_time" gorm:"column:speed_time"`
	// 请求返回请求头
	Header datatypes.JSON `json:"header" gorm:"column:header"`
	// 请求返回内容
	Body datatypes.JSON `json:"body" gorm:"column:body"`
	// 请求返回 cookie
	Cookie datatypes.JSON `json:"cookie" gorm:"column:cookie"`
}

// ReportDataApp report 模型实例
var ReportDataApp = new(ReportData)

// New 初始化模型
func (i *ReportData) New(data ReportData) *ReportData {
	return &data
}

// Get 获取信息
func (i *ReportData) Get(size, page int) (data []ReportData, err error) {
	o := (page - 1) * size
	err = global.APP_DB.Model(&i).Limit(size).Offset(o).Where("report_uid = ?", i.ReportUid).Find(&data).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 测试报告数据表#获取信息失败", zap.Error(err))
		return []ReportData{}, err
	}
	return
}

// Get2 通过UID获取请求信息
func (i *ReportData) Get2() (data ReportData, err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Find(&data).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 测试报告数据表#获取信息失败", zap.Error(err))
		return ReportData{}, err
	}
	return
}

// Count 获取信息总数
func (i *ReportData) Count() (count int64, err error) {
	if err = global.APP_DB.Model(&i).Where("report_uid = ?", i.ReportUid).Count(&count).Error; err != nil {
		global.APP_LOG.Error("[数据库] 测试报告数据表#获取信息总数失败", zap.Error(err))
		return -1, err
	}
	return
}
