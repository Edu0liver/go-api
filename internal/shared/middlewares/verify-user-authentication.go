package middlware

import (
	"fmt"
	"go-api/internal/modules/usuario/entity"
	logger_config "go-api/internal/shared/logger"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyUserAuthenticationMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	logger := logger_config.GetLogger("VerifyUserAuthenticationMiddleware")

	if err != nil {
		c.JSON(401, gin.H{"error": "Token not found"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		c.Set("user", &entity.User{
			Id:    claims["Id"].(string),
			Email: claims["Email"].(string),
		})

		c.Next()

		logger.Info(claims["Id"], claims["Email"], claims["nbf"])
	} else {
		logger.Error(err)
	}
}
