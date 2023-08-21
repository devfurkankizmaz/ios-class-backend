package models

import (
	"github.com/google/uuid"
	"time"
)

type Gallery struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	TravelID  uuid.UUID  `gorm:"type:uuid;not null"`
	ImageURL  string     `gorm:"type:varchar(255);not null"`
	Caption   string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:current_timestamp"`
}

type GalleryInput struct {
	TravelID uuid.UUID `json:"travel_id"`
	ImageURL string    `json:"image_url"`
	Caption  string    `json:"caption"`
}

type GalleryResponse struct {
	ID        *uuid.UUID `json:"id"`
	TravelID  uuid.UUID  `json:"travel_id"`
	ImageURL  string     `json:"image_url"`
	Caption   string     `json:"caption"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type GalleryRepository interface {
	Create(gallery *Gallery) error
	FetchAllByTravelID(travelID string) ([]Gallery, error)
	DeleteImageByTravelID(travelID, imageID string) error
}

type GalleryService interface {
	Create(gallery *Gallery) error
	FetchAllByTravelID(travelID string) ([]Gallery, error)
	DeleteImageByTravelID(travelID, imageID string) error
}
