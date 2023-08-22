package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type PlaceHandler struct {
	PlaceService models.PlaceService
}

func (ph *PlaceHandler) Create(c echo.Context) error {
	validate := validator.New()
	var payload *models.PlaceInput
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}

	userRole := c.Get("x-user-role")
	value := fmt.Sprint(userRole)

	if value != "admin" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"status": "fail", "message": "You are not authorized"})
	}

	newPlace := models.Place{
		Title:       payload.Title,
		Description: payload.Description,
		Latitude:    payload.Latitude,
		Longitude:   payload.Longitude,
	}

	err = ph.PlaceService.Create(&newPlace)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := fmt.Sprintf("Place successfully created.")
	return c.JSON(http.StatusCreated, echo.Map{"status": "success", "message": response})
}

func (ph *PlaceHandler) FetchAll(c echo.Context) error {
	limitParam := c.QueryParam("limit")
	pageParam := c.QueryParam("page")

	limit := 100
	if limitParam != "" {
		limit, _ = strconv.Atoi(limitParam)
	}

	page := 1
	if pageParam != "" {
		page, _ = strconv.Atoi(pageParam)
	}

	places, err := ph.PlaceService.FetchAll(limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	response := make([]models.PlaceResponse, len(places))

	for k, v := range places {
		response[k] = models.PlaceResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"count": len(places), "places": response}})
}

func (ph *PlaceHandler) FetchByID(c echo.Context) error {
	placeId := c.Param("placeId")
	if placeId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	place, err := ph.PlaceService.FetchByID(placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := models.PlaceResponse{
		ID:          place.ID,
		Title:       place.Title,
		Description: place.Description,
		Latitude:    place.Latitude,
		Longitude:   place.Longitude,
		CreatedAt:   place.CreatedAt,
		UpdatedAt:   place.UpdatedAt,
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"place": response}})
}

func (ph *PlaceHandler) UpdateByID(c echo.Context) error {
	validate := validator.New()
	placeId := c.Param("placeId")
	var payload *models.PlaceInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	if placeId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	cr := time.Now()

	userRole := c.Get("x-user-role")
	value := fmt.Sprint(userRole)

	if value != "admin" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"status": "fail", "message": "You are not authorized"})
	}

	updatedPlace := models.Place{
		Title:       payload.Title,
		Description: payload.Description,
		Latitude:    payload.Latitude,
		Longitude:   payload.Longitude,
		UpdatedAt:   &cr,
	}

	err = ph.PlaceService.UpdateByID(&updatedPlace, placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "Place successfully updated"})
}

func (ph *PlaceHandler) DeleteByID(c echo.Context) error {
	placeId := c.Param("placeId")
	if placeId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}

	userRole := c.Get("x-user-role")
	value := fmt.Sprint(userRole)

	if value != "admin" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"status": "fail", "message": "You are not authorized"})
	}

	err := ph.PlaceService.DeleteByID(placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "Place successfully deleted"})
}
