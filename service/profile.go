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

func (ps *profileService) ChangePassword(userID string, newPassword string) error {
	return ps.userRepository.ChangePassword(userID, newPassword)

}

func (ps *profileService) EditProfile(userID string, newEmail string, newFullName string, newPP string) error {
	return ps.userRepository.EditProfile(userID, newEmail, newFullName, newPP)
}

func capitalizeFirstLetter(input string) string {
	if len(input) == 0 {
		return input
	}
	return strings.ToUpper(input[0:1]) + input[1:]
}
