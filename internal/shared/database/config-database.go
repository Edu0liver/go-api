package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (database *Database) InitConfig() *Database {
	return &Database{}
}

func (database *Database) InitDatabase() *gorm.DB {

	dsn := "host=localhost user=docker password=docker dbname=postgres port=5432 sslmode=disable TimeZone=America/Brasilia"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.db = db

	return db
}
