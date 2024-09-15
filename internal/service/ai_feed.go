package service

import (
	"ai-feed/internal/generator"
	"ai-feed/internal/storage"
)

// AiFeed is service implementation for CRUD operations for themes, articles and personalities, and article generation
type AiFeed struct {
	themes        storage.Theme
	articles      storage.Article
	personalities storage.Personality

	ai *generator.AI
}

func NewAiFeed(themes storage.Theme, articles storage.Article, personalities storage.Personality, ai *generator.AI) *AiFeed {
	return &AiFeed{
		themes:        themes,
		articles:      articles,
		personalities: personalities,
		ai:            ai,
	}
}
