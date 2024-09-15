package service

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (service *AiFeed) CreateTheme(ctx context.Context, theme *entity.Theme) error {
	log.Info().Interface("theme", theme).Msg("create theme")
	return service.themes.Create(ctx, theme)
}

func (service *AiFeed) ReadAllThemes(ctx context.Context) ([]*entity.Theme, error) {
	log.Info().Msg("read all themes")
	return service.themes.ReadAll(ctx)
}

func (service *AiFeed) UpdateTheme(ctx context.Context, theme *entity.Theme) error {
	log.Info().Interface("theme", theme).Msg("update theme")
	return service.themes.Update(ctx, theme)
}

func (service *AiFeed) DeleteTheme(ctx context.Context, ID uuid.UUID) error {
	log.Info().Str("id", ID.String()).Msg("delete theme")
	return service.themes.Delete(ctx, ID)
}
