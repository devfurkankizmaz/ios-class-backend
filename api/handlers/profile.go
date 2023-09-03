package handlers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"

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
