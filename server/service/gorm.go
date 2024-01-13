package service

import (
	"HttpTest-Server/global"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func Gorm() *gorm.DB {
	// Gorm 配置
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: "ht_",
			// 禁用表复数
			SingularTable: true,
		},
	}

	// 打开数据库
	db, err := gorm.Open(sqlite.Open(global.DbPathFile), &config)
	if err != nil {
		global.APP_LOG.Panic("[初始化] Sqlite3数据库初始化失败", zap.Error(err))
	}

	// 开启 WAL模式
	db.Exec("PRAGMA journal_mode=WAL;")

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		global.APP_LOG.Panic("[初始化] Sqlite3数据库连接池初始化失败", zap.Error(err))
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
