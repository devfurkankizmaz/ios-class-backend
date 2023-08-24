package handlers

import (
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GalleryHandler struct {
	GalleryService models.GalleryService
}

func NewGalleryHandler(galleryService models.GalleryService) *GalleryHandler {
	return &GalleryHandler{
		GalleryService: galleryService,
	}
}

func (gh *GalleryHandler) AddImageToTravel(c echo.Context) error {
	var payload models.GalleryInput
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": err.Error()})
	}

	image := &models.Gallery{
		PlaceID:  payload.PlaceID,
		ImageURL: payload.ImageURL,
	}

	err := gh.GalleryService.Create(image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "Image added to gallery"})
}

func (gh *GalleryHandler) GetImagesByPlaceID(c echo.Context) error {
	placeID := c.Param("placeId")
	if placeID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": "travelId param not found"})
	}

	images, err := gh.GalleryService.FetchAllByPlaceID(placeID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	response := make([]models.GalleryResponse, len(images))
	for i, img := range images {
		response[i] = models.GalleryResponse{
			ID:        img.ID,
			PlaceID:   img.PlaceID,
			ImageURL:  img.ImageURL,
			CreatedAt: img.CreatedAt,
			UpdatedAt: img.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "data": echo.Map{"count": len(images), "images": response}})
}

func (gh *GalleryHandler) DeleteImageByPlaceID(c echo.Context) error {
	placeID := c.Param("placeId")
	if placeID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": "placeId param not found"})
	}

	imageID := c.Param("imageId")
	if imageID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"status": "fail", "message": "imageId param not found"})
	}

	err := gh.GalleryService.DeleteImageByPlaceID(placeID, imageID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"status": "error", "message": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": "success", "message": "Image deleted from gallery"})
}
