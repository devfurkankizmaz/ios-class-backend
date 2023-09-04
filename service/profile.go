package service

import (
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"strings"
)

type profileService struct {
	userRepository models.UserRepository
}

func NewProfileService(repo models.UserRepository) models.ProfileService {
	return &profileService{userRepository: repo}
}

func (ps *profileService) FetchProfileByID(id string) (*models.Profile, error) {
	user, err := ps.userRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}
	capitalizedRole := capitalizeFirstLetter(user.Role)

	return &models.Profile{FullName: user.FullName, Email: user.Email, Role: capitalizedRole, PPUrl: user.PPUrl, CreatedAt: user.CreatedAt}, nil
}

func (ps *profileService) ChangePassword(id string, newPassword string) error {
	// You can add validation logic here if needed

	// Call the UserRepository to change the password
	err := ps.userRepository.ChangePassword(id, newPassword)
	if err != nil {
		return err // Handle the error, such as user not found or a database error
	}

	return nil
}

func (ps *profileService) EditProfile(id string, updatedProfile *models.User) error {

	updatedUser := &models.User{
		FullName: updatedProfile.FullName,
		Email:    updatedProfile.Email,
		PPUrl:    updatedProfile.PPUrl,
	}

	// Call the UserRepository to edit the profile
	err := ps.userRepository.EditProfile(id, updatedUser)
	if err != nil {
		return err // Handle the error, such as user not found or a database error
	}

	return nil
}

func capitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}
	return strings.ToUpper(input[0:1]) + input[1:]
}
