package postgres

import (
	"example/messages_api/internal/message"
	"example/messages_api/internal/user"

	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var dbase *gorm.DB

// Init
// Initializes the database
func Init() *gorm.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		return nil
	}

	// Load environment variables
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("SSLMODE")

	// Creating database source name string
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&message.Message{}, &user.User{})

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
