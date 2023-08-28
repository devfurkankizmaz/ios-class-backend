package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"

	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type PlaceHandler struct {
	PlaceService models.PlaceService
	VisitService models.VisitService
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

	userID := c.Get("x-user-id")
	value := fmt.Sprint(userID)
	UID, _ := uuid.Parse(value)

	userFullName := c.Get("x-user-full-name")
	fullName := fmt.Sprint(userFullName)

	newPlace := models.Place{
		UserID:        UID,
		Creator:       fullName,
		Place:         payload.Place,
		Title:         payload.Title,
		Description:   payload.Description,
		Latitude:      payload.Latitude,
		Longitude:     payload.Longitude,
		CoverImageUrl: payload.CoverImageUrl,
	}

	err = ph.PlaceService.Create(&newPlace)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	now := time.Now()
	newVisit := models.Visit{
		UserID:    UID,
		PlaceID:   *newPlace.ID,
		VisitedAt: &now,
	}

	err = ph.VisitService.Create(&newVisit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"status": "success", "message": newPlace.ID.String()})
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
			ID:            v.ID,
			Creator:       v.Creator,
			Place:         v.Place,
			Title:         v.Title,
			Description:   v.Description,
			Latitude:      v.Latitude,
			Longitude:     v.Longitude,
			CoverImageUrl: v.CoverImageUrl,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
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
	isVisited := false
	visit, err := ph.VisitService.FetchByPlaceID(placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	if visit != nil {
		isVisited = true
	}

	response := models.PlaceResponse{
		ID:            place.ID,
		Creator:       place.Creator,
		Place:         place.Place,
		Title:         place.Title,
		Description:   place.Description,
		Latitude:      place.Latitude,
		Longitude:     place.Longitude,
		CoverImageUrl: place.CoverImageUrl,
		CreatedAt:     place.CreatedAt,
		UpdatedAt:     place.UpdatedAt,
		IsVisited:     isVisited,
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"place": response}})
}

func (ph *PlaceHandler) UpdateByID(c echo.Context) error {
	validate := validator.New()
	placeId := c.Param("placeId")
	var payload *models.PlaceInput

	if placeId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}

	userID := c.Get("x-user-id")
	value := fmt.Sprint(userID)
	UID, _ := uuid.Parse(value)

	place, err := ph.PlaceService.FetchByID(placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	if place.UserID != UID {
		return c.JSON(http.StatusUnauthorized, echo.Map{"status": "fail", "message": "You are not creator of this place"})
	}

	err = c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}

	cr := time.Now()

	updatedPlace := models.Place{
		Place:         payload.Place,
		Title:         payload.Title,
		Description:   payload.Description,
		Latitude:      payload.Latitude,
		Longitude:     payload.Longitude,
		CoverImageUrl: payload.CoverImageUrl,
		UpdatedAt:     &cr,
	}

	err = ph.PlaceService.UpdateByID(&updatedPlace, placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "Place successfully updated"})
}

func (ph *PlaceHandler) FetchAllByUserID(c echo.Context) error {
	newUID := fmt.Sprint(c.Get("x-user-id"))
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

	places, err := ph.PlaceService.FetchAllByUserID(newUID, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	response := make([]models.PlaceResponse, len(places))

	for k, v := range places {
		isVisited := false
		visit, err := ph.VisitService.FetchByPlaceID(v.ID.String())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
		}
		if visit != nil {
			isVisited = true
		}
		response[k] = models.PlaceResponse{
			ID:            v.ID,
			Creator:       v.Creator,
			Place:         v.Place,
			Title:         v.Title,
			Description:   v.Description,
			CoverImageUrl: v.CoverImageUrl,
			Latitude:      v.Latitude,
			Longitude:     v.Longitude,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
			IsVisited:     isVisited,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"count": len(places), "places": response}})
}

func (ph *PlaceHandler) DeleteByID(c echo.Context) error {
	placeId := c.Param("placeId")
	if placeId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}

	userID := c.Get("x-user-id")
	value := fmt.Sprint(userID)
	UID, _ := uuid.Parse(value)

	place, err := ph.PlaceService.FetchByID(placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	if place.UserID != UID {
		return c.JSON(http.StatusUnauthorized, echo.Map{"status": "fail", "message": "You are not creator of this place"})
	}

	err = ph.PlaceService.DeleteByID(placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "Place successfully deleted"})
}
