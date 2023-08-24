package service

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/devfurkankizmaz/iosclass-backend/utils"
)

type placeService struct {
	placeRepository models.PlaceRepository
}

func NewPlaceService(repo models.PlaceRepository) models.PlaceService {
	return &placeService{placeRepository: repo}
}

func (ps *placeService) Create(place *models.Place) error {
	if !utils.IsValidLatitude(place.Latitude) || !utils.IsValidLongitude(place.Longitude) {
		return errors.New("invalid latitude or longitude")
	}

	err := ps.placeRepository.Create(place)
	if err != nil {
		return err
	}
	return nil
}

func (ps *placeService) FetchAllByUserID(id string, limit int, page int) ([]models.Place, error) {
	offset := (page - 1) * page
	places, err := ps.placeRepository.FetchAllByUserID(id, limit, offset)
	if err != nil {
		return places, err
	}
	return places, nil
}

func (ps *placeService) FetchAll(limit int, page int) ([]models.Place, error) {
	offset := (page - 1) * page
	places, err := ps.placeRepository.FetchAll(limit, offset)
	if err != nil {
		return places, err
	}
	return places, nil
}

func (ps *placeService) FetchByID(id string) (models.Place, error) {

	place, err := ps.placeRepository.FetchByID(id)
	if err != nil {
		return place, err
	}
	return place, nil

}

func (ps *placeService) UpdateByID(place *models.Place, id string) error {
	if !utils.IsValidLatitude(place.Latitude) || !utils.IsValidLongitude(place.Longitude) {
		return errors.New("invalid latitude or longitude")
	}

	err := ps.placeRepository.UpdateByID(place, id)
	if err != nil {
		return err
	}
	return nil
}

func (ps *placeService) DeleteByID(id string) error {
	err := ps.placeRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
