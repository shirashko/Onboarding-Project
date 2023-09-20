package profile

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api/profile/model"
	errors "github.com/shirrashko/BuildingAServer-step2/pkg/error"
)

// ProfileRepository This ProfileRepository struct will encapsulate the operations related to the user profiles using the PostgreSQL database connection.
type ProfileRepository struct {
	client *sql.DB
}

func NewRepository(client *sql.DB) ProfileRepository {
	return ProfileRepository{client: client}
}

// implementation of the methods of the ProfileRepository object, which regard to the repository contains users profile info

func (repo *ProfileRepository) UpdateProfile(newProfile model.UserProfile) error {
	query := "UPDATE userProfiles SET username = $1, full_name = $2, bio = $3, profile_pic_url = $4 WHERE id = $5"
	_, err := repo.client.Exec(query, newProfile.Username, newProfile.FullName, newProfile.Bio, newProfile.ProfilePicURL, newProfile.ID) // execute the prepared SQL query.
	if err != nil {
		fmt.Printf("Error updating profile: %v\n", err)
		return err
	}
	return nil
}

func (repo *ProfileRepository) CreateNewProfile(newProfile model.BaseUserProfile) (int, error) {
	query := "INSERT INTO userProfiles (username, full_name, bio, profile_pic_url) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := repo.client.QueryRow(query, newProfile.Username, newProfile.FullName, newProfile.Bio, newProfile.ProfilePicURL).Scan(&id)
	if err != nil {
		fmt.Printf("Error creating profile: %v\n", err)
		return 0, err
	}
	return id, nil
}

func (repo *ProfileRepository) GetProfileByID(id int) (model.UserProfile, error) {
	query := "SELECT id, username, full_name, bio, profile_pic_url FROM userProfiles WHERE id = $1"
	var userProfile model.UserProfile
	err := repo.client.QueryRow(query, id).Scan(&userProfile.ID, &userProfile.Username, &userProfile.FullName, &userProfile.Bio, &userProfile.ProfilePicURL)

	if err != nil {
		if err == sql.ErrNoRows { // Handle the case where no rows were found.
			return userProfile, errors.NewNotFoundError(err)
		}
		fmt.Printf("Error querying user profile: %v\n", err)
		return userProfile, err
	}
	return userProfile, nil
}
