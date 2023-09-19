package main

import (
	"go-api/internal/shared/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	port := ":8080"

	routes.InitializerRoutes(app)

	app.Run(port)
}
