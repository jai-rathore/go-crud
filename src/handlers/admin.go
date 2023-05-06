package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jai-rathore/lists/config"
)

type AdminHandler struct{}

func NewAdminHandler(r *gin.Engine) *AdminHandler {
	handler := &AdminHandler{}
	r.GET("/health", handler.HealthCheck)
	return handler
}

// HealthCheck godoc
// @Summary Health Check
// @Schemes
// @Description health check endpoint
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {string} health
// @Router /health [get]
func (a *AdminHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, "Lists Server is up and running at port : "+config.AppConfig.Port)
}
