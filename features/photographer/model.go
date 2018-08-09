package photographer

import "time"

type onlineProfiles struct {
	Facebook   string `json:"profile_facebook"`
	Twitter    string `json:"profile_twitter"`
	GooglePlus string `json:"profile_google"`
	Instagram  string `json:"profile_instagram"`
	Website    string `json:"profile_website"`
}

//Photographer Represents a user that uploads images on the platform
type Photographer struct {
	ID             string         `json:"user_id"`
	Username       string         `json:"username"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Email          string         `json:"email"`
	Migrate        string         `json:"migrate"` //How t
	City           string         `json:"city"`
	Country        string         `json:"country"`
	DateOfBirth    string         `json:"dob"`
	Description    string         `json:"about_me"`
	OnlineProfiles onlineProfiles `json:"online_profile"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
