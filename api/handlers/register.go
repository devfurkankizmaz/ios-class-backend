package handlers

import (
	"fmt"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type RegisterHandler struct {
	RegisterService models.RegisterService
}

func (rh *RegisterHandler) Register(c echo.Context) error {
	validate := validator.New()
	var payload *models.RegisterInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
	}

	newUser := models.User{
		FullName: payload.FullName,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
	}

	err = rh.RegisterService.Create(&newUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}
	response := fmt.Sprintf("Inserted ID: %s", newUser.ID)
	return c.JSON(http.StatusCreated, echo.Map{"status": "success", "message": response})
}