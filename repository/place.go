package repository

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"gorm.io/gorm"
)

type placeRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) models.PlaceRepository {
	return &placeRepository{db: db}
}

func (pr *placeRepository) Create(place *models.Place) error {
	result := pr.db.Create(&place)
	if result.Error != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (pr *placeRepository) FetchAllByUserID(id string, limit int, offset int) ([]models.Place, error) {
	var places = []models.Place{}
	result := pr.db.Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&places)
	if result.Error != nil {
		return places, result.Error
	}
	return places, nil
}

func (pr *placeRepository) FetchAll(limit int, offset int) ([]models.Place, error) {
	var places = []models.Place{}
	result := pr.db.Order("created_at desc").Limit(limit).Offset(offset).Find(&places)
	if result.Error != nil {
		return places, result.Error
	}
	return places, nil
}

func (pr *placeRepository) FetchByID(id string) (models.Place, error) {
	var place = models.Place{}
	result := pr.db.Where("id = ?", id).First(&place)

	if result.Error != nil {
		return place, result.Error
	}
	return place, nil
}

func (pr *placeRepository) UpdateByID(place *models.Place, id string) error {
	columbs := map[string]interface{}{"place": place.Place, "title": place.Title, "description": place.Description, "cover_image_url": place.CoverImageUrl, "latitude": place.Latitude, "longitude": place.Longitude}
	result := pr.db.Model(&place).Where("id = ?", id).Updates(columbs)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *placeRepository) DeleteByID(id string) error {
	result := pr.db.Delete(&models.Place{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return result.Error
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
