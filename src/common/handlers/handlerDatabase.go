package handlers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	models "api/src/common/types"
)

var DB *gorm.DB

func InitializeDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}