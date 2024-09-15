package entity

import (
	"github.com/google/uuid"
	"time"
)

type Theme struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id,omitempty"`

	CreatedAt   time.Time `gorm:"autoUpdateTime" json:"created_at"`
	Description string    `json:"description" validate:"required"`
	Deleted     bool      `json:"-" gorm:"type:bool;default:FALSE;"`
}
