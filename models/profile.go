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

type ProfileService interface {
	FetchProfileByID(id string) (*Profile, error)
	ChangePassword(userID string, newPassword string) error
	EditProfile(userID string, newEmail string, newFullName string, newPP string) error
}
