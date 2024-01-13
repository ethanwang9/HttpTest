package model

import (
	"HttpTest-Server/global"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

// TestData 自动化测试数据表
type TestData struct {
	Base
	// 数据ID
	UID string `json:"uid" gorm:"column:uid;primaryKey"`
	// 场景名称
	Name string `json:"name" gorm:"column:name"`
	// 接口列表
	InterfaceList datatypes.JSON `json:"list" gorm:"column:interface_list"`
	// 测试相关参数
	Test datatypes.JSON `json:"test" gorm:"column:test"`
}

// TestDataApp test_data 模型实例
var TestDataApp = new(TestData)

// New 初始化模型
func (i *TestData) New(data TestData) *TestData {
	return &data
}

// Set 设置信息
func (i *TestData) Set() (err error) {
	err = global.APP_DB.Model(&i).Create(&i).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 自动化测试数据表#设置场景数据失败", zap.Error(err))
	}
	return
}

// Delete 删除信息
func (i *TestData) Delete() (err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Unscoped().Delete(&i).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 自动化测试数据表#删除场景数据失败", zap.Error(err))
	}
	return
}

// Get 获取信息
func (i *TestData) Get() (data TestData, err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Find(&data).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 自动化测试数据表#获取场景信息失败", zap.Error(err))
		return TestData{}, err
	}
	return
}

// Update 更新信息
func (i *TestData) Update() (err error) {
	err = global.APP_DB.Model(&i).Where("uid = ?", i.UID).Updates(i).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 自动化测试数据表#更新场景信息失败", zap.Error(err))
	}
	return
}
