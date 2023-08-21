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

	stringUuid := (gallery.TravelID).String()
	existingImages, err := gs.galleryRepository.FetchAllByTravelID(stringUuid)
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

func (gs *galleryService) FetchAllByTravelID(travelID string) ([]models.Gallery, error) {
	galleryImages, err := gs.galleryRepository.FetchAllByTravelID(travelID)
	if err != nil {
		return galleryImages, err
	}
	return galleryImages, nil
}

func (gs *galleryService) DeleteImageByTravelID(travelID, imageID string) error {
	err := gs.galleryRepository.DeleteImageByTravelID(travelID, imageID)
	if err != nil {
		return err
	}
	return nil
}
