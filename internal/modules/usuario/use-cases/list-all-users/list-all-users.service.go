package list_all_users

import (
	"go-api/internal/modules/usuario/entity"
	user_repository "go-api/internal/modules/usuario/repository"
)

func ListAllUsersService() (*[]entity.User, error) {
	if users, err := user_repository.ListAllUsers(); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}
