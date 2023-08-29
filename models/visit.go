package models

import (
	"time"

	"github.com/google/uuid"
)

type Visit struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	PlaceID   uuid.UUID  `gorm:"type:uuid;not null"`
	VisitedAt *time.Time `gorm:"type:timestamptz;not null"`
	CreatedAt *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type VisitInput struct {
	PlaceID   uuid.UUID  `json:"place_id" validate:"required"`
	VisitedAt *time.Time `json:"visited_at" validate:"required"`
}

type VisitResponse struct {
	ID        *uuid.UUID    `json:"id"`
	PlaceID   uuid.UUID     `json:"place_id"`
	VisitedAt *time.Time    `json:"visited_at"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
	Place     PlaceResponse `json:"place"`
}

type VisitRepository interface {
	Create(visit *Visit) error
	FetchAllByUserID(id string, limit int, offset int) ([]Visit, error)
	FetchByPlaceIDAndUserID(placeID string, userID string) (Visit, error)
	FetchByID(id string) (Visit, error)
	DeleteByID(id string) error
}

type VisitService interface {
	Create(visit *Visit) error
	FetchAllByUserID(id string, limit int, page int) ([]Visit, error)
	FetchByPlaceIDAndUserID(placeID string, userID string) (Visit, error)
	FetchByID(id string) (Visit, error)
	DeleteByID(id string) error
}
