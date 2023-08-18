package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TravelHandler struct {
	TravelService models.TravelService
}

func (th *TravelHandler) Create(c echo.Context) error {
	validate := validator.New()
	var payload *models.TravelInput
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
	newTravel := models.Travel{
		UserID:      UID,
		VisitDate:   payload.VisitDate,
		Location:    payload.Location,
		Information: payload.Information,
		Latitude:    payload.Latitude,
		Longitude:   payload.Longitude,
	}

	err = th.TravelService.Create(&newTravel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := fmt.Sprintf("Inserted ID: %s", newTravel.ID)
	return c.JSON(http.StatusCreated, echo.Map{"status": "success", "message": response})
}

func (th *TravelHandler) FetchAllByUserID(c echo.Context) error {
	newUID := fmt.Sprint(c.Get("x-user-id"))

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}

	travels, err := th.TravelService.FetchAllByUserID(newUID, limit, page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	response := make([]models.TravelResponse, len(travels))

	for k, v := range travels {
		response[k] = models.TravelResponse{
			ID:          v.ID,
			VisitDate:   v.VisitDate,
			Location:    v.Location,
			Information: v.Information,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"count": len(travels), "travels": response}})
}

func (th *TravelHandler) FetchByID(c echo.Context) error {
	travelId := c.Param("travelId")
	if travelId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	travel, err := th.TravelService.FetchByID(travelId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := models.TravelResponse{
		ID:          travel.ID,
		VisitDate:   travel.VisitDate,
		Location:    travel.Location,
		Information: travel.Information,
		Latitude:    travel.Latitude,
		Longitude:   travel.Longitude,
		CreatedAt:   travel.CreatedAt,
		UpdatedAt:   travel.UpdatedAt,
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"travel": response}})
}

func (th *TravelHandler) UpdateByID(c echo.Context) error {
	validate := validator.New()
	travelId := c.Param("travelId")
	var payload *models.TravelInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	if travelId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	cr := time.Now()
	updatedTravel := models.Travel{
		VisitDate:   payload.VisitDate,
		Location:    payload.Location,
		Information: payload.Information,
		Latitude:    payload.Latitude,
		Longitude:   payload.Longitude,
		UpdatedAt:   &cr,
	}

	err = th.TravelService.UpdateByID(&updatedTravel, travelId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "updated_at": updatedTravel.UpdatedAt})
}

func (th *TravelHandler) DeleteByID(c echo.Context) error {
	travelId := c.Param("travelId")
	if travelId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	err := th.TravelService.DeleteByID(travelId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "travel successfully deleted"})
}
