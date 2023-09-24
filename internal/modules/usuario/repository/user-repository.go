package user_repository

import (
	"go-api/internal/modules/usuario/entity"
	"go-api/internal/shared/database"
)

func CreateUser(user *entity.User) (*entity.User, error) {
	err := database.DB.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func ListAllUsers() (*[]entity.User, error) {
	var users []entity.User

	err := database.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func GetUserByEmail(email *string) *entity.User {
	var user entity.User

	database.DB.Where("email = ?", email).First(&user)

	return &user
}
