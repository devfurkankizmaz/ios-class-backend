package handlers

import (
	"fmt"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type VisitHandler struct {
	VisitService models.VisitService
	PlaceService models.PlaceService
}

func (vh *VisitHandler) Create(c echo.Context) error {
	validate := validator.New()
	var payload *models.VisitInput
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

	newVisit := models.Visit{
		UserID:    UID,
		PlaceID:   payload.PlaceID,
		VisitedAt: payload.VisitedAt,
	}

	err = vh.VisitService.Create(&newVisit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := fmt.Sprintf("Visit successfully created.")
	return c.JSON(http.StatusCreated, echo.Map{"status": "success", "message": response})
}

func (vh *VisitHandler) FetchAllByUserID(c echo.Context) error {
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

	visits, err := vh.VisitService.FetchAllByUserID(newUID, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	response := make([]models.VisitResponse, len(visits))

	for k, v := range visits {
		placeIDString := v.PlaceID.String()
		place, err := vh.PlaceService.FetchByID(placeIDString)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
		}

		placeResponse := models.PlaceResponse{
			ID:            place.ID,
			Creator:       place.Creator,
			Place:         place.Place,
			Title:         place.Title,
			Description:   place.Description,
			CoverImageUrl: place.CoverImageUrl,
			Latitude:      place.Latitude,
			Longitude:     place.Longitude,
			CreatedAt:     place.CreatedAt,
			UpdatedAt:     place.UpdatedAt,
		}

		visitResponse := models.VisitResponse{
			ID:        v.ID,
			PlaceID:   v.PlaceID,
			VisitedAt: v.VisitedAt,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Place:     placeResponse,
		}

		response[k] = visitResponse

	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"count": len(visits), "visits": response}})
}

func (vh *VisitHandler) FetchByID(c echo.Context) error {
	visitId := c.Param("visitId")
	if visitId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	visit, err := vh.VisitService.FetchByID(visitId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	placeIDString := visit.PlaceID.String()
	place, err := vh.PlaceService.FetchByID(placeIDString)

	placeResponse := models.PlaceResponse{
		ID:            place.ID,
		Creator:       place.Creator,
		Place:         place.Place,
		Title:         place.Title,
		Description:   place.Description,
		CoverImageUrl: place.CoverImageUrl,
		Latitude:      place.Latitude,
		Longitude:     place.Longitude,
		CreatedAt:     place.CreatedAt,
		UpdatedAt:     place.UpdatedAt,
	}

	response := models.VisitResponse{
		ID:        visit.ID,
		PlaceID:   visit.PlaceID,
		VisitedAt: visit.VisitedAt,
		CreatedAt: visit.CreatedAt,
		UpdatedAt: visit.UpdatedAt,
		Place:     placeResponse,
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"visit": response}})
}

func (vh *VisitHandler) DeleteByID(c echo.Context) error {
	visitId := c.Param("visitId")
	if visitId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	err := vh.VisitService.DeleteByID(visitId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "Visit successfully deleted"})
}
