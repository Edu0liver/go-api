package database

import (
	"go-api/internal/modules/usuario/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (database *Database) InitDatabase() *gorm.DB {

	dsn := "host=localhost user=docker password=docker dbname=postgres port=5432 sslmode=disable TimeZone=America/Fortaleza"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.db = db

	db.AutoMigrate(&entity.User{})

	return db
}

var Db = new(Database).InitDatabase()
