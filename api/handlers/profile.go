package handlers

import (
	"fmt"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"

	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	ProfileService models.ProfileService
}

func (ph *ProfileHandler) Fetch(c echo.Context) error {
	userID := c.Get("x-user-id")
	newUID := fmt.Sprint(userID)
	profile, err := ph.ProfileService.FetchProfileByID(newUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error(), Status: "error"})
	}
	return c.JSON(http.StatusOK, profile)
}

func (ph *ProfileHandler) ChangePassword(c echo.Context) error {
	userID := c.Get("x-user-id")
	newUID := fmt.Sprint(userID)

	var changePasswordRequest models.ChangePasswordInput
	if err := c.Bind(&changePasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "Invalid request body", Status: "error"})
	}

	// Validate the request body using the validator package
	validate := validator.New()
	if err := validate.Struct(changePasswordRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}

	err = ph.ProfileService.ChangePassword(newUID, string(hashedPassword))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error(), Status: "error"})
	}

	return c.JSON(http.StatusOK, models.Response{Message: "Password changed successfully", Status: "success"})
}

func (ph *ProfileHandler) UpdateProfile(c echo.Context) error {
	userID := c.Get("x-user-id")
	newUID := fmt.Sprint(userID)

	var payload models.EditProfileInput
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "Invalid request body", Status: "error"})
	}

	// Perform validation using the validator package
	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}

	existingProfile, err := ph.ProfileService.FetchProfileByID(newUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error(), Status: "error"})
	}

	// Update the profile fields only if they are not empty in the request payload
	if payload.FullName != "" {
		existingProfile.FullName = payload.FullName
	}
	if payload.Email != "" {
		existingProfile.Email = payload.Email
	}
	if payload.PPUrl != "" {
		existingProfile.PPUrl = payload.PPUrl
	}

	updatedAt := time.Now()
	updatedUser := &models.User{
		FullName:  payload.FullName,
		Email:     payload.Email,
		PPUrl:     payload.PPUrl,
		UpdatedAt: &updatedAt,
	}

	err = ph.ProfileService.EditProfile(newUID, updatedUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error(), Status: "error"})
	}

	return c.JSON(http.StatusOK, models.Response{Message: "Profile updated successfully", Status: "success"})
}
