package main

import (
	"log/slog"
	"os"

	"test-app/internal/health"
	"test-app/internal/middleware"

	"github.com/gin-gonic/gin"

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

	// Set up Gin router
	router := gin.Default()

	// Set up routes
	api_v1 := router.Group("/api/v1")
	api_v1.Use(middleware.IPRateLimitMiddleware())
	health.SetupRoutes(api_v1)

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info("Starting server on :8000")
	router.Run(":8000")
}
