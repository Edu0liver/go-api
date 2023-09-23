package auth_routes

import (
	"go-api/internal/modules/auth/use-cases/login"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	userRoutes := r.Group("/auth")
	{
		userRoutes.POST("/login", login.LoginController)
		// userRoutes.GET("/profile", get_profile.GetProfileController)
	}
}
