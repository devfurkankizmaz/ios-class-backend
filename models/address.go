package models

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID           *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null"`
	AddressTitle string     `gorm:"type:varchar(50);not null"`
	State        string     `gorm:"type:varchar(50)"`
	City         string     `gorm:"type:varchar(50);not null"`
	Country      string     `gorm:"type:varchar(50);not null"`
	Address      string     `gorm:"type:varchar(255);not null"`
	CreatedAt    *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt    *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type AddressInput struct {
	AddressTitle string `json:"address_title" validate:"required"`
	State        string `json:"state"`
	City         string `json:"city" validate:"required"`
	Country      string `json:"country" validate:"required"`
	Address      string `json:"address" validate:"required"`
}

type AddressResponse struct {
	AddressTitle string     `json:"address_title"`
	State        string     `json:"state"`
	City         string     `json:"city"`
	Country      string     `json:"country"`
	Address      string     `json:"address"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type AddressRepository interface {
	Create(address *Address) error
	FetchAllByUserID(id string, limit int, offset int) ([]Address, error)
	FetchByID(id string) (Address, error)
	UpdateByID(address *Address, id string) error
	DeleteByID(id string) error
}

type AddressService interface {
	Create(address *Address) error
	FetchAllByUserID(id string, limit int, page int) ([]Address, error)
	FetchByID(id string) (Address, error)
	UpdateByID(address *Address, id string) error
	DeleteByID(id string) error
}
