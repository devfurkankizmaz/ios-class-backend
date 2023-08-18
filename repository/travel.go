package repository

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"gorm.io/gorm"
)

type travelRepository struct {
	db *gorm.DB
}

func NewTravelRepository(db *gorm.DB) models.TravelRepository {
	return &travelRepository{db: db}
}

func (tr *travelRepository) Create(travel *models.Travel) error {
	result := tr.db.Create(&travel)
	if result.Error != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (tr *travelRepository) FetchAllByUserID(id string, limit int, offset int) ([]models.Travel, error) {
	var travels = []models.Travel{}
	result := tr.db.Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&travels)
	if result.Error != nil {
		return travels, result.Error
	}
	return travels, nil
}

func (tr *travelRepository) FetchByID(id string) (models.Travel, error) {
	var travel = models.Travel{}
	result := tr.db.Where("id = ?", id).First(&travel)

	if result.Error != nil {
		return travel, result.Error
	}
	return travel, nil
}

func (tr *travelRepository) UpdateByID(travel *models.Travel, id string) error {
	columbs := map[string]interface{}{"visit_date": travel.VisitDate, "location": travel.Location, "information": travel.Information, "image_url": travel.ImageUrl, "latitude": travel.Latitude, "longitude": travel.Longitude}
	result := tr.db.Model(&travel).Where("id = ?", id).Updates(columbs)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *travelRepository) DeleteByID(id string) error {
	result := tr.db.Delete(&models.Travel{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return result.Error
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
