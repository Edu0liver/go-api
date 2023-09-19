package create_user

import (
	"fmt"
	"go-api/internal/modules/usuario/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

func CreateUserController(c *gin.Context) {
	body := CreateUserRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(body)

	user := entity.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	if user, err := CreateUserService(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}
