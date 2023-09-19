package user_routes

import (
	create_user "go-api/internal/modules/usuario/use-cases/create-user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/", create_user.CreateUserController)
	}
}
