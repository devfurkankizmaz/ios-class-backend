package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AddressHandler struct {
	AddressService models.AddressService
}

func (ah *AddressHandler) Create(c echo.Context) error {
	validate := validator.New()
	var payload *models.AddressInput
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
	newAddress := models.Address{
		UserID:       UID,
		AddressTitle: payload.AddressTitle,
		State:        payload.State,
		City:         payload.City,
		Country:      payload.Country,
		Address:      payload.Address,
	}

	err = ah.AddressService.Create(&newAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := fmt.Sprintf("Inserted ID: %s", newAddress.ID)
	return c.JSON(http.StatusCreated, echo.Map{"status": "success", "message": response})
}

func (ah *AddressHandler) FetchAllByUserID(c echo.Context) error {
	newUID := fmt.Sprint(c.Get("x-user-id"))

	addresses, err := ah.AddressService.FetchAllByUserID(newUID, 0, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	response := make([]models.AddressResponse, len(addresses))

	for k, v := range addresses {
		response[k] = models.AddressResponse{
			ID:           v.ID,
			AddressTitle: v.AddressTitle,
			State:        v.State,
			City:         v.City,
			Country:      v.Country,
			Address:      v.Address,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"count": len(addresses), "addresses": response}})
}

func (ah *AddressHandler) FetchByID(c echo.Context) error {
	addressId := c.Param("addressId")
	if addressId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	address, err := ah.AddressService.FetchByID(addressId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	response := models.AddressResponse{
		ID:           address.ID,
		AddressTitle: address.AddressTitle,
		State:        address.State,
		City:         address.City,
		Country:      address.Country,
		Address:      address.Address,
		CreatedAt:    address.CreatedAt,
		UpdatedAt:    address.UpdatedAt,
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"address": response}})
}

func (ah *AddressHandler) UpdateByID(c echo.Context) error {
	validate := validator.New()
	addressId := c.Param("addressId")
	var payload *models.AddressInput

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	err = validate.Struct(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}
	if addressId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	cr := time.Now()
	updatedAddress := models.Address{
		AddressTitle: payload.AddressTitle,
		State:        payload.State,
		City:         payload.City,
		Country:      payload.Country,
		Address:      payload.Address,
		UpdatedAt:    &cr,
	}

	err = ah.AddressService.UpdateByID(&updatedAddress, addressId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "updated_at": updatedAddress.UpdatedAt})
}

func (ah *AddressHandler) DeleteByID(c echo.Context) error {
	addressId := c.Param("addressId")
	if addressId == "" {
		return c.JSON(http.StatusNotFound, echo.Map{"status": "fail", "message": "param not found"})
	}
	err := ah.AddressService.DeleteByID(addressId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "address successfully deleted"})
}
