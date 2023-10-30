package routes

import (
	"go-api/docs"
	auth_routes "go-api/internal/modules/auth/routes"
	user_routes "go-api/internal/modules/usuario/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializerRoutes(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"

	user_routes.UserRoutes(router)
	auth_routes.AuthRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
