package service

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
)

type galleryService struct {
	galleryRepository models.GalleryRepository
}

func NewGalleryService(repo models.GalleryRepository) models.GalleryService {
	return &galleryService{galleryRepository: repo}
}

func (gs *galleryService) Create(gallery *models.Gallery) error {

	stringUuid := (gallery.PlaceID).String()
	existingImages, err := gs.galleryRepository.FetchAllByPlaceID(stringUuid)
	if err != nil {
		return err
	}

	if len(existingImages) >= 3 {
		return errors.New("maximum number of images reached for this travel")
	}

	err = gs.galleryRepository.Create(gallery)
	if err != nil {
		return err
	}
	return nil
}

func (gs *galleryService) FetchAllByPlaceID(placeID string) ([]models.Gallery, error) {
	galleryImages, err := gs.galleryRepository.FetchAllByPlaceID(placeID)
	if err != nil {
		return galleryImages, err
	}
	return galleryImages, nil
}

func (gs *galleryService) DeleteImageByPlaceID(placeID, imageID string) error {
	err := gs.galleryRepository.DeleteImageByPlaceID(placeID, imageID)
	if err != nil {
		return err
	}
	return nil
}
