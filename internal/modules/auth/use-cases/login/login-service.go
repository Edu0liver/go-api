package login

import (
	"errors"
	login "go-api/internal/modules/auth/use-cases/login/dtos"
	user_repository "go-api/internal/modules/usuario/repository"
	jwt_provider "go-api/internal/shared/providers/jwt/dtos"
	jwt "go-api/internal/shared/providers/jwt/services"

	"golang.org/x/crypto/bcrypt"
)

func LoginService(data *login.LoginDto) (string, error) {
	user := user_repository.GetUserByEmail(&data.Email)

	if user == nil {
		return "", errors.New("User not found")
	}

	isPasswordCorrect := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if isPasswordCorrect != nil {
		return "", isPasswordCorrect
	}

	payload := jwt_provider.JwtPayload{
		Email: user.Email,
	}

	token, err := jwt.GenerateToken(payload)

	if err != nil {
		return "", err
	}

	return token, nil
}
