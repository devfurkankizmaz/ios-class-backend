package models

import "time"

type Profile struct {
	FullName  string     `json:"full_name"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	PPUrl     string     `json:"pp_url"`
	CreatedAt *time.Time `json:"created_at"`
}

type EditProfileInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	PPUrl    string `json:"pp_url"`
}

type ChangePasswordInput struct {
	NewPassword string `json:"new_password"`
}

type ProfileService interface {
	FetchProfileByID(id string) (*Profile, error)
	ChangePassword(id string, newPassword string) error
	EditProfile(id string, updatedProfile *User) error
}
