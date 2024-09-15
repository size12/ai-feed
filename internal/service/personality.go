package service

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (service *AiFeed) CreatePersonality(ctx context.Context, personality *entity.Personality) error {
	log.Info().Interface("personality", personality).Msg("create personality")
	return service.personalities.Create(ctx, personality)
}

func (service *AiFeed) ReadAllPersonalities(ctx context.Context) ([]*entity.Personality, error) {
	log.Info().Msg("read all personality")
	return service.personalities.ReadAll(ctx)
}

func (service *AiFeed) UpdatePersonality(ctx context.Context, personality *entity.Personality) error {
	log.Info().Interface("personality", personality).Msg("update personality")
	return service.personalities.Update(ctx, personality)
}

func (service *AiFeed) DeletePersonality(ctx context.Context, ID uuid.UUID) error {
	log.Info().Str("id", ID.String()).Msg("delete personality")
	return service.personalities.Delete(ctx, ID)
}
