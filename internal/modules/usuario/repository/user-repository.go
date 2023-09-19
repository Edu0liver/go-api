package user_repository

import (
	"go-api/internal/modules/usuario/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (userRepository *UserRepository) Create(user *entity.User) *entity.User {
	userRepository.db.Create(user)
	return user
}
