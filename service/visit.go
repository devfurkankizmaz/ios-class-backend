package service

import (
	"github.com/devfurkankizmaz/iosclass-backend/models"
)

type visitService struct {
	visitRepository models.VisitRepository
}

func NewVisitService(repo models.VisitRepository) models.VisitService {
	return &visitService{visitRepository: repo}
}

func (vs *visitService) Create(visit *models.Visit) error {
	err := vs.visitRepository.Create(visit)
	if err != nil {
		return err
	}
	return nil
}

func (vs *visitService) FetchByPlaceID(id string) ([]models.Visit, error) {
	visits, err := vs.visitRepository.FetchByPlaceID(id)
	if err != nil {
		return visits, err
	}
	return visits, nil
}

func (vs *visitService) FetchAllByUserID(id string, limit int, page int) ([]models.Visit, error) {
	offset := (page - 1) * page
	visits, err := vs.visitRepository.FetchAllByUserID(id, limit, offset)
	if err != nil {
		return visits, err
	}
	return visits, nil
}

func (vs *visitService) FetchByID(id string) (models.Visit, error) {

	visit, err := vs.visitRepository.FetchByID(id)
	if err != nil {
		return visit, err
	}
	return visit, nil

}

func (vs *visitService) DeleteByID(id string) error {
	err := vs.visitRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
