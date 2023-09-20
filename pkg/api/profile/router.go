package profile

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) SetUpRoutes(router *gin.Engine) {
	// all endpoints are with the same prefix "/profile/users"
	profilesRouter := router.Group("/profile/users")
	profilesRouter.GET("/:id", h.getProfileByID)    // get user profile by id
	profilesRouter.POST("", h.createProfile)        // create user profile, user data is in the body?
	profilesRouter.PUT("/:id", h.updateProfileByID) // update user profile by id
	// check if also need to support delete and patch?
}
