package database

import (
	"example/messages_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

// Init
// Initializes the database
func Init() *gorm.DB {
	dsn := "user=postgres password=admin dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&models.Message{}, &models.User{})

	if err != nil {
		return nil
	}

	return db
}

// GetDB
// database getter
func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
	}
	return dbase
}
