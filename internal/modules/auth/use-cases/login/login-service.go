package login

import (
	"errors"
	login "go-api/internal/modules/auth/use-cases/login/dtos"
	user_repository "go-api/internal/modules/usuario/repository"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(data *login.LoginDto) (string, error) {
	user := user_repository.GetUserByEmail(&data.Email)

	if user.Id == "" {
		return "", errors.ErrUnsupported
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return "", err
	}

	payload := jwt.MapClaims{
		"Id":    user.Id,
		"Email": user.Email,
		"nbf":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
