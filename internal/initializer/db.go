package initializer

import (
	"mvc/internal/helpers"
	"mvc/internal/models"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.ErrorPanic(err)
}

func SyncDB() {
DB.AutoMigrate(&models.Post{})
DB.AutoMigrate(&models.Author{})
}

