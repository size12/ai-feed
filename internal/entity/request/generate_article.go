package request

import "ai-feed/internal/entity"

type GenerateRequest struct {
	Theme       *entity.Theme       `json:"theme" validate:"required"`
	Personality *entity.Personality `json:"personality" validate:"required"`
}
