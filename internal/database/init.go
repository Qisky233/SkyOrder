package database

import (
	"fmt"
	"log"
	"order/internal/model"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 初始化数据库连接
func Init(dbPath string) error {
	// 确保数据库目录存在
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create database directory failed: %v", err)
	}

	// 配置GORM日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return fmt.Errorf("connect database failed: %v", err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(
		&model.Sys_account{},
	)
	if err != nil {
		return fmt.Errorf("migrate database failed: %v", err)
	}

	DB = db
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
