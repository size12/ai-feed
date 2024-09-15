package entity

import (
	"github.com/google/uuid"
	"time"
)

type Personality struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id,omitempty"`

	CreatedAt time.Time `gorm:"autoUpdateTime" json:"created_at"`

	Name      string `json:"name" validate:"required"`
	Biography string `json:"biography" validate:"required"`
	Keywords  string `json:"keywords" validate:"required"`
	Thematics string `json:"thematics" validate:"required"`
	TextStyle string `json:"text_style" validate:"required"`
}
