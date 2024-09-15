package service

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (service *AiFeed) CreateArticle(ctx context.Context, article *entity.Article) error {
	log.Info().Interface("article", article).Msg("create article")
	return service.articles.Create(ctx, article)
}

func (service *AiFeed) ReadArticle(ctx context.Context, ID uuid.UUID) (*entity.Article, error) {
	log.Info().Interface("id", ID).Msg("read article")
	return service.articles.Read(ctx, ID)
}

func (service *AiFeed) ReadAllArticles(ctx context.Context) ([]*entity.Article, error) {
	log.Info().Msg("read all articles")
	return service.articles.ReadAll(ctx)
}

func (service *AiFeed) UpdateArticle(ctx context.Context, article *entity.Article) error {
	log.Info().Interface("article", article).Msg("update article")
	return service.articles.Update(ctx, article)
}

func (service *AiFeed) DeleteArticle(ctx context.Context, ID uuid.UUID) error {
	log.Info().Interface("id", ID).Msg("read article")
	return service.articles.Delete(ctx, ID)
}
