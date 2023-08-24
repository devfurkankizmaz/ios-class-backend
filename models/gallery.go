package models

import (
	"github.com/google/uuid"
	"time"
)

type Gallery struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	PlaceID   uuid.UUID  `gorm:"type:uuid;not null"`
	ImageURL  string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:current_timestamp"`
}

type GalleryInput struct {
	PlaceID  uuid.UUID `json:"place_id" validate:"required"`
	ImageURL string    `json:"image_url" validate:"required"`
}

type GalleryResponse struct {
	ID        *uuid.UUID `json:"id"`
	PlaceID   uuid.UUID  `json:"place_id"`
	ImageURL  string     `json:"image_url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type GalleryRepository interface {
	Create(gallery *Gallery) error
	FetchAllByPlaceID(placeID string) ([]Gallery, error)
	DeleteImageByPlaceID(placeID, imageID string) error
}

type GalleryService interface {
	Create(gallery *Gallery) error
	FetchAllByPlaceID(placeID string) ([]Gallery, error)
	DeleteImageByPlaceID(placeID, imageID string) error
}
