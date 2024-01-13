package model

import (
	"HttpTest-Server/global"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

// System 系统信息表
type System struct {
	Base
	// 系统ID
	SID uint8 `json:"sid" gorm:"column:sid;primaryKey"`
	// 接口管理目录树
	InterfaceMenu string `json:"interface_menu" gorm:"column:interface_menu"`
	// 自动化测试目录树
	TestMenu string `json:"test_menu" gorm:"column:test_menu"`
	// 项目 Github 仓库地址
	GitUrl string `json:"git_url" gorm:"column:git_url"`
	// 项目文档地址
	DocUrl string `json:"doc_url" gorm:"column:doc_url"`
	// 系统代理池列表
	ProxyList datatypes.JSON `json:"proxy_list" gorm:"column:proxy_list"`
}

// SystemApp system 模型实例
var SystemApp = new(System)

// New 初始化模型
func (s *System) New(system System) *System {
	return &system
}

// Get 获取信息
func (s *System) Get() (data System, err error) {
	err = global.APP_DB.Model(&s).Where("sid = ?", "1").Find(&data).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 系统信息表#获取系统信息失败", zap.Error(err))
		return System{}, err
	}

	return
}

// Set 设置信息
func (s *System) Set() (err error) {
	s.SID = 1
	err = global.APP_DB.Model(&s).Create(&s).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 系统信息表#设置系统信息失败", zap.Error(err))
		return err
	}

	return nil
}

// Update 更新信息
func (s *System) Update() (err error) {
	s.SID = 1
	err = global.APP_DB.Model(&s).Where("sid = ?", "1").Updates(s).Error
	if err != nil {
		global.APP_LOG.Error("[数据库] 系统信息表#更新系统信息失败", zap.Error(err))
		return err
	}

	return nil
}
