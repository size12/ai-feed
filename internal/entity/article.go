package entity

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`

	CreatedAt time.Time `gorm:"autoUpdateTime" json:"created_at"`

	Title       string `json:"title" validate:"required"`
	ImageBase64 string `json:"image_base64"`
	Content     string `json:"content" validate:"required"`

	WordsCount   int      `json:"words_count"`
	SymbolsCount int      `json:"symbols_count"`
	Keywords     Keywords `gorm:"type:BYTEA" json:"keywords"`
}
