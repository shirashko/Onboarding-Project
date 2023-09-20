package health

import (
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl/health"
	"net/http"
)

type Handler struct {
	service *health.Service
}

func NewHandler(s *health.Service) Handler {
	return Handler{s}
}

func (h Handler) HealthCheck(c *gin.Context) {
	if h.service.HealthCheck() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
	}
}
