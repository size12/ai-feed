package storage

import (
	"ai-feed/internal/entity"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewORM returns *gorm.DB which makes working with database easier.
func NewORM(cfg *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.PostgresEndpoint), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("failed open postgres db")
	}

	err = db.AutoMigrate(entity.Personality{}, entity.Theme{}, entity.Article{}, entity.User{})

	if err != nil {
		log.Fatal().Err(err).Msg("failed auto migrate")
	}

	return db
}
