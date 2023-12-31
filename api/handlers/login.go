package handlers

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	LoginService models.LoginService
}

func (lh *LoginHandler) Login(c echo.Context) error {
	validate := validator.New()
	var payload *models.LoginInput
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}

	// E-posta adresini küçük harfe çevirerek işleyin
	payload.Email = strings.ToLower(payload.Email)

	user, err := lh.LoginService.FetchByEmail(payload.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{Message: "User not found with the given email", Status: "not found"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, models.Response{Message: "Invalid credentials", Status: "unauthorized"})
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	jwtExpiresInStr := os.Getenv("JWT_EXPIRED_IN")
	jwtRefreshExpiresInStr := os.Getenv("JWT_REFRESH_EXPIRED_IN")
	jwtRefreshExpiresIn, err := strconv.Atoi(jwtRefreshExpiresInStr)
	jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)
	accessToken, err := lh.LoginService.GenerateAccessToken(&user, jwtSecret, jwtExpiresIn)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error(), Status: "error"})
	}
	refreshToken, err := lh.LoginService.GenerateRefreshToken(&user, jwtRefreshSecret, jwtRefreshExpiresIn)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error(), Status: "error"})
	}

	return c.JSON(http.StatusOK, models.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
