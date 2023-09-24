package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	dsn := os.Getenv("DB_CONNNECTION")
	ok_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = ok_db

	if err != nil {
		panic("Failed to connect to database!")
	}

}
