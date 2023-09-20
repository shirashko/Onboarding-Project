package api

import "github.com/gin-gonic/gin"

// IRoutesHandler IRoutesHandlers is an interface that defines the behavior of a routes handler in a web API.
// It specifies the method required to set up routes for a specific part of the API.
type IRoutesHandler interface {
	SetUpRoutes(router *gin.Engine)
}
