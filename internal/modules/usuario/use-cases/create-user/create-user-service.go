package create_user

import (
	"go-api/internal/modules/usuario/entity"
	user_repository "go-api/internal/modules/usuario/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(user *entity.User) (*entity.User, error) {
	if pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16); err != nil {
		return &entity.User{}, err
	} else {
		user.Password = string(pass)
	}

	uuid := uuid.New()
	user.Id = uuid.String()

	if user, err := user_repository.CreateUser(user); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
