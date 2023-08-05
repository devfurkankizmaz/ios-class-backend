package service

import (
	"fmt"
	"github.com/devfurkankizmaz/iosclass-backend/models"
)

type addressService struct {
	addressRepository models.AddressRepository
}

func NewAddressService(repo models.AddressRepository) models.AddressService {
	return &addressService{addressRepository: repo}
}

func (as *addressService) Create(address *models.Address) error {
	stringId := fmt.Sprint(address.UserID)
	currentAddresses, err := as.addressRepository.FetchAllByUserID(stringId, 0, 0)
	if err != nil {
		return err
	}

	if len(currentAddresses) >= 5 {
		return fmt.Errorf("cannot create more than 5 addresses")
	}

	err = as.addressRepository.Create(address)
	if err != nil {
		return err
	}
	return nil
}

func (as *addressService) FetchAllByUserID(id string, limit int, page int) ([]models.Address, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}
	offset := (page - 1) * page
	addresses, err := as.addressRepository.FetchAllByUserID(id, limit, offset)
	if err != nil {
		return addresses, err
	}
	return addresses, nil
}

func (as *addressService) FetchByID(id string) (models.Address, error) {
	address, err := as.addressRepository.FetchByID(id)
	if err != nil {
		return address, err
	}
	return address, nil
}

func (as *addressService) UpdateByID(address *models.Address, id string) error {
	err := as.addressRepository.UpdateByID(address, id)
	if err != nil {
		return err
	}
	return nil
}

func (as *addressService) DeleteByID(id string) error {
	err := as.addressRepository.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
