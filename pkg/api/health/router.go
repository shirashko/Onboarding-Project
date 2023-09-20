package health

import "github.com/gin-gonic/gin"

func (h Handler) SetUpRoutes(router *gin.Engine) {
	router.GET("/health", h.HealthCheck) // Associate the GET HTTP method and /health path with a handler function
}
