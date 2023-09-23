package user_repository

import (
	"go-api/internal/modules/usuario/entity"
	"go-api/internal/shared/database"

	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func CreateUser(user *entity.User) (*entity.User, error) {
	err := database.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func ListAllUsers() (*[]entity.User, error) {
	var users []entity.User

	err := database.Db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func GetUserByEmail(email *string) *entity.User {
	var user entity.User

	database.Db.Where("email = ?", email).First(&user)

	return &user
}
