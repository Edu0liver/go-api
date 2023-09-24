package main

import (
	"go-api/internal/shared/database"
	"go-api/internal/shared/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	app := gin.Default()
	port := ":8080"

	godotenv.Load()
	database.InitDatabase()
	database.SyncDatabase()
	routes.InitializerRoutes(app)

	app.Run(port)
}
