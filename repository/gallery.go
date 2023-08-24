package repository

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"gorm.io/gorm"
)

type galleryRepository struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) models.GalleryRepository {
	return &galleryRepository{db: db}
}

func (gr *galleryRepository) Create(gallery *models.Gallery) error {
	result := gr.db.Create(&gallery)
	if result.Error != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (gr *galleryRepository) FetchAllByPlaceID(placeID string) ([]models.Gallery, error) {
	var galleryImages = []models.Gallery{}
	result := gr.db.Where("place_id = ?", placeID).Find(&galleryImages)
	if result.Error != nil {
		return galleryImages, result.Error
	}
	return galleryImages, nil
}

func (gr *galleryRepository) DeleteImageByPlaceID(placeID, imageID string) error {
	result := gr.db.Delete(&models.Gallery{}, "place_id = ? AND id = ?", placeID, imageID)
	if result.RowsAffected == 0 {
		return result.Error
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
