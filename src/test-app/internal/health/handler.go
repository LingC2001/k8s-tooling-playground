package health

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health check
// @Description Returns the health status of the application
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string "status: OK, app: test-app"
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
		"app":    "test-app",
	})
}
