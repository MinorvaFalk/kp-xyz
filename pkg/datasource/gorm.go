package datasource

import (
	"kp/config"
	"kp/pkg/logger"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.ReadConfig().Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.NewGormLogger(),
	})
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql db instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping sql db: %v", err)
	}

	return db
}
