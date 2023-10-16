package main

import (
	"go-api/internal/shared/database"
	logger_config "go-api/internal/shared/logger"
	"go-api/internal/shared/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	logger *logger_config.Logger
)

func main() {
	app := gin.Default()
	port := ":8080"
	logger = logger_config.GetLogger("main")

	godotenv.Load()
	database.InitDatabase()
	database.SyncDatabase()
	routes.InitializerRoutes(app)

	app.Run(port)
}
