package profile

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile/model"
	"github.com/shirrashko/BuildingAServer-step2/pkg/bl/profile"
	e "github.com/shirrashko/BuildingAServer-step2/pkg/error"
	"net/http"
)

type Handler struct {
	service *profile.Service
}

func NewHandler(s *profile.Service) Handler {
	return Handler{s}
}

func (h Handler) getProfileByID(c *gin.Context) {
	var request model.GetProfileRequest
	if err := c.ShouldBindUri(&request); err != nil { //
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	// Get the user's userProfile
	userProfile, err := h.service.GetProfileByID(request.ID)
	if err != nil {
		if errors.Is(err, e.UserNotFoundError{}) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving userProfile"})
			return
		}
	}

	// Respond with the userProfile
	c.JSON(http.StatusOK, userProfile)
}

func (h Handler) updateProfileByID(c *gin.Context) {
	var updatedProfile model.UpdateProfileRequest

	if err := c.ShouldBindUri(&(updatedProfile.Profile)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Bind(&updatedProfile.Profile.BaseUserProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateUserProfile(updatedProfile.Profile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or error updating profile"})
		return
	}

	c.JSON(http.StatusOK, updatedProfile.Profile)
}

func (h Handler) createProfile(c *gin.Context) {
	var newProfile model.CreateProfileRequest
	if err := c.Bind(&newProfile.Profile); err != nil { // bind JSON data from the request body into a Go struct
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newID, err := h.service.CreateNewProfile(newProfile.Profile)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, newID) // output to the request body the id that the new profile got in the system
}
