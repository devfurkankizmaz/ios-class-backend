package service

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"github.com/devfurkankizmaz/iosclass-backend/utils"
)

type travelService struct {
	travelRepository models.TravelRepository
}

func NewTravelService(repo models.TravelRepository) models.TravelService {
	return &travelService{travelRepository: repo}
}

func (ts *travelService) Create(travel *models.Travel) error {
	if !utils.IsValidLatitude(travel.Latitude) || !utils.IsValidLongitude(travel.Longitude) {
		return errors.New("invalid latitude or longitude")
	}

	err := ts.travelRepository.Create(travel)
	if err != nil {
		return err
	}
	return nil
}

func (ts *travelService) FetchAllByUserID(id string, limit int, page int) ([]models.Travel, error) {
	offset := (page - 1) * page
	travels, err := ts.travelRepository.FetchAllByUserID(id, limit, offset)
	if err != nil {
		return travels, err
	}
	return travels, nil
}

func (ts *travelService) FetchByID(id string) (models.Travel, error) {

	travel, err := ts.travelRepository.FetchByID(id)
	if err != nil {
		return travel, err
	}
	return travel, nil

}

func (ts *travelService) UpdateByID(travel *models.Travel, id string) error {
	if !utils.IsValidLatitude(travel.Latitude) || !utils.IsValidLongitude(travel.Longitude) {
		return errors.New("invalid latitude or longitude")
	}

	err := ts.travelRepository.UpdateByID(travel, id)
	if err != nil {
		return err
	}
	return nil
}

func (ts *travelService) DeleteByID(id string) error {
	err := ts.travelRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
