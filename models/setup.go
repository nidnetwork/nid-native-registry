package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB gorm.DB
var DB *gorm.DB

// ConnectDatabase connect database
func ConnectDatabase(dsn string) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.Debug()
	db.AutoMigrate(&Record{})
	DB = db
}
