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
	Information string     `gorm:"type:text"`
	CreatedAt   *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt   *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type TravelInput struct {
	VisitDate   *time.Time `json:"visit_date" validate:"required"`
	Location    string     `json:"location" validate:"required"`
	Information string     `json:"information"`
}

type TravelResponse struct {
	VisitDate   *time.Time `json:"visit_date"`
	Location    string     `json:"location"`
	Information string     `json:"information"`
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
