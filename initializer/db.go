package initializer

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
 var DB *gorm.DB
func ConnectionDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil {
		fmt.Println("failed to connect db")
	}
}