package profile

import (
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile/model"
	"github.com/shirrashko/BuildingAServer-step2/pkg/repository/profile"
)

// Service The service has a repository with a client field, which is the connection to the database we are working with
type Service struct {
	repository *profile.ProfileRepository
}

func NewService(profileRepo *profile.ProfileRepository) Service {
	return Service{repository: profileRepo}
}

func (s *Service) UpdateUserProfile(newProfile model.UserProfile) error {
	return s.repository.UpdateProfile(newProfile)
}

func (s *Service) CreateNewProfile(newProfile model.BaseUserProfile) (int, error) {
	return s.repository.CreateNewProfile(newProfile) // return newID and error
}

func (s *Service) GetProfileByID(id int) (model.UserProfile, error) {
	return s.repository.GetProfileByID(id)
}
