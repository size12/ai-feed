package service

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/metric"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (service *AiFeed) CreateArticle(ctx context.Context, article *entity.Article) error {
	updateMetrics(article)
	log.Info().Interface("article_id", article.ID).Msg("create article")

	return service.articles.Create(ctx, article)
}

func (service *AiFeed) ReadArticle(ctx context.Context, ID uuid.UUID) (*entity.Article, error) {
	log.Info().Interface("id", ID).Msg("read article")
	return service.articles.Read(ctx, ID)
}

func (service *AiFeed) ReadAllArticles(ctx context.Context) ([]*entity.Article, error) {
	log.Info().Msg("read all Articles")
	return service.articles.ReadAll(ctx)
}

func (service *AiFeed) UpdateArticle(ctx context.Context, article *entity.Article) error {
	updateMetrics(article)
	log.Info().Interface("article", article).Msg("update article")

	return service.articles.Update(ctx, article)
}

func (service *AiFeed) DeleteArticle(ctx context.Context, ID uuid.UUID) error {
	log.Info().Interface("id", ID).Msg("read article")
	return service.articles.Delete(ctx, ID)
}

func updateMetrics(article *entity.Article) {
	article.WordsCount = metric.WordsCount(article.Content)
	article.SymbolsCount = metric.SymbolsCount(article.Content)
	article.Keywords = metric.Keywords(article.Content)
}
