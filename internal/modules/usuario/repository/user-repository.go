package user_repository

import (
	"go-api/internal/modules/usuario/entity"
	"go-api/internal/shared/database"

	"gorm.io/gorm"
)

// type UserRepository interface {
// 	Create(user *entity.User) (*entity.User, error)
// }

type UserRepositoryDb struct {
	Db *gorm.DB
}

func UserRepositoryCreate(user *entity.User) (*entity.User, error) {
	err := database.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
