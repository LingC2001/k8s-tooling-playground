package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "test-app/docs" // swag doc generation

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title test-app API
// @version 1.0
// @description This is a simple test application API.
// @BasePath /api/v1
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Load environment variables from .env
	if err := godotenv.Load("../.env"); err != nil {
		logger.Warn("No .env file found, using system environment variables")
	}

	// Set up Gin router
	router := gin.Default()

	// Set up routes
	api_v1 := router.Group("/api/v1")
	{
		api_v1.GET("/health", func(c *gin.Context) {
			logger.Info("Health check requested")
			c.JSON(200, gin.H{
				"status": "OK",
				"app":    "test-app",
			})
		})
	}

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info("Starting server on :8080")
	router.Run(":8080")
}
