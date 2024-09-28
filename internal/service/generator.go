package service

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/rs/zerolog/log"
)

func (service *AiFeed) GenerateArticle(ctx context.Context, theme *entity.Theme, personality *entity.Personality) (*entity.Article, error) {
	log.Info().Interface("theme", theme).Interface("personality", personality).Msg("generate article content")
	return service.ai.GenerateArticle(ctx, theme, personality)
}

func (service *AiFeed) GenerateArticleImage(ctx context.Context, article *entity.Article) (string, error) {
	log.Info().Interface("article_id", article.ID).Msg("generate article image")
	return service.ai.GenerateArticleImage(ctx, article)
}
