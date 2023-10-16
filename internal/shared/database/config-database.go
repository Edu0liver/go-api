package database

import (
	"go-api/internal/shared/logger"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDatabase() {
	log := logger.GetLogger("Init Database")

	dsn := os.Getenv("DB_CONNNECTION")
	ok_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = ok_db

	if err != nil {
		log.Errorf("Failed to connect to database: %v", err)
	}

}
