package storage

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Article interface for CRUD storage for articles.
type Article interface {
	Create(ctx context.Context, article *entity.Article) error
	ReadAll(ctx context.Context) ([]*entity.Article, error)
	Read(ctx context.Context, ID uuid.UUID) (*entity.Article, error)
	Update(ctx context.Context, article *entity.Article) error
	Delete(ctx context.Context, ID uuid.UUID) error
}

func NewArticle(db *gorm.DB) Article {
	return newArticleImpl(db)
}

// articleImpl is implementation of Article interface
type articleImpl struct {
	db *gorm.DB
}

func newArticleImpl(db *gorm.DB) *articleImpl {
	return &articleImpl{
		db: db,
	}
}

func (a *articleImpl) Create(ctx context.Context, article *entity.Article) error {
	login := ctx.Value(UserLogin).(string)
	article.OwnerLogin = login

	if article.ID.ID() == 0 {
		article.ID = uuid.New()
	}

	result := a.db.Create(article)

	if result.Error != nil {
		log.Err(result.Error).Interface("article", article).Msg("failed add article to db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotCreated
	}

	return nil
}

func (a *articleImpl) ReadAll(ctx context.Context) ([]*entity.Article, error) {
	login := ctx.Value(UserLogin).(string)

	var articles []*entity.Article

	result := a.db.Where("owner_login = ?", login).Order("created_at DESC").Find(&articles)

	if result.Error != nil {
		log.Err(result.Error).Msg("failed read all articles from db")
		return nil, result.Error
	}

	return articles, nil
}

func (a *articleImpl) Read(ctx context.Context, ID uuid.UUID) (*entity.Article, error) {
	var article *entity.Article

	result := a.db.Where("id = ?", ID).Find(&article)

	if result.Error != nil {
		log.Err(result.Error).Str("id", ID.String()).Msg("failed read articles from db by id")
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}

	return article, nil
}

func (a *articleImpl) Update(ctx context.Context, article *entity.Article) error {
	login := ctx.Value(UserLogin).(string)

	result := a.db.Model(&entity.Article{}).Where("id = ? AND owner_login = ?", article.ID, login).Updates(article)

	if result.Error != nil {
		log.Err(result.Error).Interface("article", article).Msg("failed update article in db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrFailedUpdate
	}

	return nil
}

func (a *articleImpl) Delete(ctx context.Context, ID uuid.UUID) error {
	login := ctx.Value(UserLogin).(string)

	article := &entity.Article{}

	result := a.db.Where("id = ? AND owner_login = ?", ID, login).Delete(article)

	if result.Error != nil {
		log.Err(result.Error).Str("article_id", ID.String()).Msg("failed delete article from db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
