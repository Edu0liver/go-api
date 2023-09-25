package user_routes

import (
	create_user "go-api/internal/modules/usuario/use-cases/create-user"
	list_all_users "go-api/internal/modules/usuario/use-cases/list-all-users"
	middlware "go-api/internal/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/", middlware.VerifyUserAuthenticationMiddleware, list_all_users.ListAllUsersController)
		userRoutes.POST("/", create_user.CreateUserController)
	}
}
