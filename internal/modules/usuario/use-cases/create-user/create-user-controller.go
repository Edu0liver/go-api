package create_user

import (
	"go-api/internal/modules/usuario/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	createUserService *CreateUserService
}

func (createUserontroller *CreateUserController) Execute(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user, err := createUserontroller.createUserService.Execute(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}
