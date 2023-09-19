package main

import (
	"go-api/internal/shared/database"
	"go-api/internal/shared/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	port := ":8080"

	db := new(database.Database)
	db.InitConfig()
	db.InitDatabase()

	routes.InitializerRoutes(app)

	app.Run(port)
}
