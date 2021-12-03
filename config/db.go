package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
	var err error

	logMode := map[string]logger.LogLevel{
		"Silent": logger.Silent,
		"Error":  logger.Error,
		"Warn":   logger.Warn,
		"Info":   logger.Info,
	}[Env.Database.LogLevel]

	db, err = gorm.Open(postgres.Open(Env.Database.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	return err
}

func DB() *gorm.DB {
	return db
}
