package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthCheck godoc
// @Summary healthCheck example
// @Schemes
// @Description healthCheck
// @Tags Health Check
// @Success 200 {string} OK
// @Router /health [get]
func health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func HealthCheckAppRoute(r *gin.RouterGroup) {
	// health check
	r.GET("/health", health)
}
