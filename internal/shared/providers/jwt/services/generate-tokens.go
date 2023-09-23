package jwt_provider

import (
	jwt_provider "go-api/internal/shared/providers/jwt/dtos"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(payload jwt_provider.JwtPayload) (string, error) {
	secretKey := "secret"
	claims := jwt.MapClaims{
		"email": payload.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
