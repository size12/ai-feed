package service

import (
	"ai-feed/internal/generator"
	"ai-feed/internal/storage"
)

// AiFeed is service implementation for CRUD operations for Themes, Articles and Personalities, article generation and user authorization
type AiFeed struct {
	themes        storage.Theme
	articles      storage.Article
	personalities storage.Personality
	users         storage.User

	ai *generator.AI

	authCfg *AuthConfig
}

func NewAiFeed(cfg *Config) *AiFeed {
	return &AiFeed{
		themes:        cfg.Themes,
		articles:      cfg.Articles,
		personalities: cfg.Personalities,
		users:         cfg.Users,
		ai:            cfg.Ai,
		authCfg:       cfg.AuthConfig,
	}
}
