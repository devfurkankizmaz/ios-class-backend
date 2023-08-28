package repository

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"gorm.io/gorm"
)

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) models.VisitRepository {
	return &visitRepository{db: db}
}

func (vr *visitRepository) Create(visit *models.Visit) error {
	result := vr.db.Create(&visit)
	if result.Error != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (vr *visitRepository) FetchAllByUserID(id string, limit int, offset int) ([]models.Visit, error) {
	var visits = []models.Visit{}
	result := vr.db.Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&visits)
	if result.Error != nil {
		return visits, result.Error
	}
	return visits, nil
}

func (vr *visitRepository) FetchByPlaceID(id string) ([]models.Visit, error) {
	var visits = []models.Visit{}
	result := vr.db.Where("place_id = ?", id).Find(&visits)
	if result.Error != nil {
		return visits, result.Error
	}
	return visits, nil
}

func (vr *visitRepository) FetchByID(id string) (models.Visit, error) {
	var visit = models.Visit{}
	result := vr.db.Where("id = ?", id).First(&visit)

	if result.Error != nil {
		return visit, result.Error
	}
	return visit, nil
}

func (vr *visitRepository) DeleteByID(id string) error {
	result := vr.db.Delete(&models.Visit{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return result.Error
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
