package models

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID          *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	Latitude    float64    `gorm:"type:float;not null"`
	Longitude   float64    `gorm:"type:float;not null"`
	CreatedAt   *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt   *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type PlaceInput struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
}

type PlaceResponse struct {
	ID          *uuid.UUID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type PlaceRepository interface {
	Create(place *Place) error
	FetchAll(limit int, offset int) ([]Place, error)
	FetchByID(id string) (Place, error)
	UpdateByID(place *Place, id string) error
	DeleteByID(id string) error
}

type PlaceService interface {
	Create(place *Place) error
	FetchAll(limit int, page int) ([]Place, error)
	FetchByID(id string) (Place, error)
	UpdateByID(place *Place, id string) error
	DeleteByID(id string) error
}
