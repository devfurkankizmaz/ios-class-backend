package models

import (
	"time"

	"github.com/google/uuid"
)

type Travel struct {
	ID          *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null"`
	VisitDate   *time.Time `gorm:"type:timestamptz;not null"`
	Location    string     `gorm:"type:varchar(255);not null"`
	ImageUrl    string     `gorm:"type:varchar(255)"`
	Information string     `gorm:"type:text"`
	Latitude    string     `gorm:"type:float"`
	Longitude   string     `gorm:"type:float"`
	CreatedAt   *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt   *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type TravelInput struct {
	VisitDate   *time.Time `json:"visit_date" validate:"required"`
	Location    string     `json:"location" validate:"required"`
	Information string     `json:"information" validate:"required"`
	ImageUrl    string     `json:"image_url" validate:"required"`
	Latitude    string     `json:"latitude" validate:"required"`
	Longitude   string     `json:"longitude" validate:"required"`
}

type TravelResponse struct {
	ID          *uuid.UUID `json:"id"`
	VisitDate   *time.Time `json:"visit_date"`
	Location    string     `json:"location"`
	Information string     `json:"information"`
	ImageUrl    string     `json:"image_url"`
	Latitude    string     `json:"latitude"`
	Longitude   string     `json:"longitude"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TravelRepository interface {
	Create(travel *Travel) error
	FetchAllByUserID(id string, limit int, offset int) ([]Travel, error)
	FetchByID(id string) (Travel, error)
	UpdateByID(travel *Travel, id string) error
	DeleteByID(id string) error
}

type TravelService interface {
	Create(travel *Travel) error
	FetchAllByUserID(id string, limit int, page int) ([]Travel, error)
	FetchByID(id string) (Travel, error)
	UpdateByID(travel *Travel, id string) error
	DeleteByID(id string) error
}
