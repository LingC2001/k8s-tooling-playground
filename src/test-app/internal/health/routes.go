package health

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(rg *gin.RouterGroup) {
	handler := NewHealthHandler()
	rg.GET("/health", handler.HealthCheck)
}
