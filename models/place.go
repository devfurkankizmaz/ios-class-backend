package models

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID            *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null"`
	Creator       string     `gorm:"type:varchar(255);not null"`
	Place         string     `gorm:"type:varchar(255);not null"`
	Title         string     `gorm:"type:varchar(255);not null"`
	Description   string     `gorm:"type:text"`
	Latitude      float64    `gorm:"type:float;not null"`
	Longitude     float64    `gorm:"type:float;not null"`
	CoverImageUrl string     `gorm:"type:varchar(255)"`
	CreatedAt     *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt     *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type PlaceInput struct {
	Place         string  `json:"place" validate:"required"`
	Title         string  `json:"title" validate:"required"`
	Description   string  `json:"description"`
	Latitude      float64 `json:"latitude" validate:"required"`
	Longitude     float64 `json:"longitude" validate:"required"`
	CoverImageUrl string  `json:"cover_image_url"`
}

type PlaceResponse struct {
	ID            *uuid.UUID `json:"id"`
	Creator       string     `json:"creator"`
	Place         string     `json:"place"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	CoverImageUrl string     `json:"cover_image_url"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	IsVisited     bool       `json:"is_visited"`
}

type PlaceRepository interface {
	Create(place *Place) error
	FetchAll(limit int, offset int) ([]Place, error)
	FetchAllByUserID(id string, limit int, offset int) ([]Place, error)
	FetchByID(id string) (Place, error)
	UpdateByID(place *Place, id string) error
	DeleteByID(id string) error
}

type PlaceService interface {
	Create(place *Place) error
	FetchAll(limit int, page int) ([]Place, error)
	FetchAllByUserID(id string, limit int, offset int) ([]Place, error)
	FetchByID(id string) (Place, error)
	UpdateByID(place *Place, id string) error
	DeleteByID(id string) error
}
