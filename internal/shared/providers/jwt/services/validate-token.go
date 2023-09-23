package jwt_provider

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(signedToken string) (string, error) {
	secretKey := "secret"

	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		username := claims["username"].(string)
		return username, nil
	}

	return "", errors.New("Invalid token")
}
