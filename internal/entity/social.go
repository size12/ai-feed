package entity

import "github.com/google/uuid"

type SocialNetwork struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id,omitempty"`

	Name       string `json:"name" validate:"required"`
	OwnerLogin string `json:"-"`

	Type string `json:"type"`

	APIKey    string `json:"apikey"`
	ChannelID string `json:"channelID"`
}
