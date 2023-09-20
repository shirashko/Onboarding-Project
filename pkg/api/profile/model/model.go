package model

type BaseUserProfile struct {
	Username      string `json:"username"`        // Username is the unique username per user
	FullName      string `json:"full_name"`       // FullName is the user's full name
	Bio           string `json:"bio"`             // Bio is the user's profile biography
	ProfilePicURL string `json:"profile_pic_url"` // ProfilePicURL is the URL to the user's profile picture
}

type UserProfile struct {
	ID              int `uri:"id" binding:"required"` // ID is the auto-generated unique identifier for the user's profile
	BaseUserProfile `binding:"-"`
}

type GetProfileRequest struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateProfileRequest struct {
	Profile UserProfile `binding:"-"`
}

type CreateProfileRequest struct {
	Profile BaseUserProfile `binding:"-"`
}
