package feeder

import (
	"ai-feed/internal/storage"
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

// Service requests actual themes from []Feeder and saves them in storage.
type Service struct {
	Feeders []Feeder
	Themes  storage.Theme

	cfg *Config
}

func NewService(cfg *Config, themes storage.Theme, feeders ...Feeder) *Service {
	return &Service{
		Feeders: feeders,
		Themes:  themes,
		cfg:     cfg,
	}
}

func (s *Service) Run(ctx context.Context) {
	log.Info().Msgf("Started feeder with %v sources", len(s.Feeders))

	ctx = context.WithValue(ctx, storage.UserLogin, "feeder")
	ticker := time.NewTicker(s.cfg.FeedUpdateDelay)
	s.updateThemes(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.updateThemes(ctx)
		}
	}
}

func (s *Service) updateThemes(ctx context.Context) {
	log.Info().Msg("Fetching new themes")
	for _, f := range s.Feeders {
		themes, err := f(ctx)
		if err != nil {
			log.Err(err).Msg("Failed fetch theme from feeder")
			continue
		}

		for _, theme := range themes {
			// CreateWithCheckDescription helps not to put themes with same description.
			err = s.Themes.CreateWithCheckDescription(ctx, theme)
			if err != nil {
				log.Err(err).Msg("failed put theme in storage")
			}
		}
	}
}
