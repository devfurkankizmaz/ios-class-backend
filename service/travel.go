package service

import "github.com/devfurkankizmaz/iosclass-backend/models"

type travelService struct {
	travelRepository models.TravelRepository
}

func NewTravelService(repo models.TravelRepository) models.TravelService {
	return &travelService{travelRepository: repo}
}

func (ts *travelService) Create(travel *models.Travel) error {
	err := ts.travelRepository.Create(travel)
	if err != nil {
		return err
	}
	return nil
}

func (ts *travelService) FetchAllByUserID(id string, limit int, page int) ([]models.Travel, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}
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
