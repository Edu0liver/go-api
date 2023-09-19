package create_user

import (
	"fmt"
	"go-api/internal/modules/usuario/entity"
	user_repository "go-api/internal/modules/usuario/repository"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	userRepository *user_repository.UserRepository
}

func (createUserService *CreateUserService) Execute(user *entity.User) (*entity.User, error) {
	if pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 16); err != nil {
		fmt.Println(err)
		return &entity.User{}, err
	} else {
		user.Password = string(pass)
	}

	return createUserService.userRepository.Create(user), nil
}
