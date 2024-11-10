package db

import (
	"log"

	"example.com/go-fiber-crud/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

// Initializes the database connection
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Dialector{DriverName: "sqlite", DSN: "file:test.db?cache=shared&mode=rwc"}, &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
