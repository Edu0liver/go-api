package database

import "go-api/internal/modules/usuario/entity"

func SyncDatabase() {
	DB.AutoMigrate(&entity.User{})
}
