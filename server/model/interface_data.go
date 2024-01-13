package model

import (
	"HttpTest-Server/global"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

// InterfaceData 接口管理数据表
type InterfaceData struct {
	Base
	// 数据ID
	UID string `json:"uid" gorm:"column:uid;primaryKey"`
	// HTTP请求类型
	Method string `json:"method" gorm:"column:method"`
	// HTTP请求地址
	Url string `json:"url" gorm:"column:url"`
	// 接口名称
	Name string `json:"name" gorm:"column:name"`
	// 请求头
	Header datatypes.JSON `json:"header" gorm:"column:header"`
	// 测试相关参数
	Test datatypes.JSON `json:"test" gorm:"column:test"`
	// 测试报告ID
	Data string `json:"data" gorm:"column:data"`
}

// InterfaceDataApp interface_data 模型实例
var InterfaceDataApp = new(InterfaceData)

// New 初始化模型
func (i *InterfaceData) New(data InterfaceData) *InterfaceData {
	return &data
}

// Set 设置信息
func (i *InterfaceData) Set() (err error) {
	err = global.APP_DB.Model(&i).Create(&i).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 接口管理数据表#设置接口数据失败", zap.Error(err))
	}
	return
}

// Get 获取信息
func (i *InterfaceData) Get() (data InterfaceData, err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Find(&data).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 接口管理数据表#获取接口信息失败", zap.Error(err))
		return InterfaceData{}, err
	}
	return
}

// Update 更新信息
func (i *InterfaceData) Update() (err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Updates(i).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 接口管理数据表#更新接口信息失败", zap.Error(err))
	}
	return
}

// Delete 删除信息
func (i *InterfaceData) Delete() (err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Unscoped().Delete(&i).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 接口管理数据表#删除接口信息失败", zap.Error(err))
	}
	return
}
