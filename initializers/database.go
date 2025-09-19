package initializers

import (
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB * gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}