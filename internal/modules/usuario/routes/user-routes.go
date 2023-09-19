package user_routes

import (
	create_user "go-api/internal/modules/usuario/use-cases/create-user"

	"github.com/gin-gonic/gin"
)

var createUserController create_user.CreateUserController

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/", createUserController.Execute)
	}
}
