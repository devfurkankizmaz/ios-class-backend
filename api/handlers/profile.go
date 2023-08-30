package handlers

import (
	"fmt"
	"github.com/go-playground/validator"
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

func (ph *ProfileHandler) Update(c echo.Context) error {
	validate := validator.New()
	var payload *models.EditProfileInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}

	userID := c.Get("x-user-id")
	newUID := fmt.Sprint(userID)

	err = ph.ProfileService.EditProfile(newUID, payload.Email, payload.FullName, payload.PPUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error(), Status: "error"})
	}
	return c.JSON(http.StatusOK, models.Response{Message: "success", Status: "success"})
}
