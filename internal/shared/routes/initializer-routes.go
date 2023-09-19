package routes

import (
	user_routes "go-api/internal/modules/usuario/routes"

	"github.com/gin-gonic/gin"
)

func InitializerRoutes(router *gin.Engine) {
	user_routes.UserRoutes(router)
}
