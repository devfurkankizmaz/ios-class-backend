package repository

import (
	"errors"
	"github.com/devfurkankizmaz/iosclass-backend/models"
	"gorm.io/gorm"
)

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) models.AddressRepository {
	return &addressRepository{db: db}
}

func (ar *addressRepository) Create(address *models.Address) error {
	result := ar.db.Create(&address)
	if result.Error != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (ar *addressRepository) FetchAllByUserID(id string, limit int, offset int) ([]models.Address, error) {
	var addresses = []models.Address{}
	result := ar.db.Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&addresses)
	if result.Error != nil {
		return addresses, result.Error
	}
	return addresses, nil
}

func (ar *addressRepository) FetchByID(id string) (models.Address, error) {
	var address = models.Address{}
	result := ar.db.Where("id = ?", id).First(&address)

	if result.Error != nil {
		return address, result.Error
	}
	return address, nil
}

func (ar *addressRepository) UpdateByID(address *models.Address, id string) error {
	columbs := map[string]interface{}{"address_title": address.AddressTitle, "city": address.City, "state": address.State, "country": address.Country, "address": address.Address, "latitude": address.Latitude, "longitude": address.Longitude}
	result := ar.db.Model(&address).Where("id = ?", id).Updates(columbs)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ar *addressRepository) DeleteByID(id string) error {
	result := ar.db.Delete(&models.Address{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return result.Error
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
