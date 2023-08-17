package handlers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RefreshHandler struct {
	RefreshService models.RefreshService
}

func (rh *RefreshHandler) Refresh(c echo.Context) error {
	validate := validator.New()
	var payload models.RefreshInput

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	jwtExpiresInStr := os.Getenv("JWT_EXPIRED_IN")
	jwtRefreshExpiresInStr := os.Getenv("JWT_REFRESH_EXPIRED_IN")
	jwtRefreshExpiresIn, err := strconv.Atoi(jwtRefreshExpiresInStr)
	jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)

	err = c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "Refresh Token Required", Status: "fail"})
	}

	id, err := rh.RefreshService.ExtractIDFromToken(payload.RefreshToken, jwtRefreshSecret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "User not found!", Status: "fail"})
	}
	user, err := rh.RefreshService.FetchByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: "User not found!", Status: "fail"})
	}
	at, err := rh.RefreshService.GenerateAccessToken(&user, jwtSecret, jwtExpiresIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}
	rt, err := rh.RefreshService.GenerateRefreshToken(&user, jwtSecret, jwtRefreshExpiresIn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Status: "fail"})
	}
	return c.JSON(http.StatusOK, models.RefreshResponse{AccessToken: at, RefreshToken: rt})
}
