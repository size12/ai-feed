package storage

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Social is interface for CRUD storage for social.
type Social interface {
	Create(ctx context.Context, social *entity.SocialNetwork) error
	Read(ctx context.Context, ID uuid.UUID) (*entity.SocialNetwork, error)
	ReadAll(ctx context.Context) ([]*entity.SocialNetwork, error)
	Update(ctx context.Context, social *entity.SocialNetwork) error
	Delete(ctx context.Context, ID uuid.UUID) error
}

func NewSocial(db *gorm.DB) Social {
	return newSocialImpl(db)
}

// socialImpl is implementation of Social interface.
type socialImpl struct {
	db *gorm.DB
}

func newSocialImpl(db *gorm.DB) *socialImpl {
	return &socialImpl{
		db: db,
	}
}

func (s *socialImpl) Create(ctx context.Context, social *entity.SocialNetwork) error {
	login := ctx.Value(UserLogin).(string)
	social.OwnerLogin = login

	if social.ID.ID() == 0 {
		social.ID = uuid.New()
	}

	result := s.db.Create(social)

	if result.Error != nil {
		log.Err(result.Error).Interface("social", social).Msg("failed add social to db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotCreated
	}

	return nil
}

func (s *socialImpl) Read(ctx context.Context, ID uuid.UUID) (*entity.SocialNetwork, error) {
	var social *entity.SocialNetwork

	result := s.db.Where("id = ?", ID).Find(&social)

	if result.Error != nil {
		log.Err(result.Error).Str("id", ID.String()).Msg("failed read socials from db by id")
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrNotFound
	}

	return social, nil
}

func (s *socialImpl) ReadAll(ctx context.Context) ([]*entity.SocialNetwork, error) {
	login := ctx.Value(UserLogin).(string)

	var socials []*entity.SocialNetwork

	result := s.db.Where("owner_login = ?", login).Find(&socials)

	if result.Error != nil {
		log.Err(result.Error).Msg("failed read all socials from db")
		return nil, result.Error
	}

	return socials, nil
}

func (s *socialImpl) Update(ctx context.Context, social *entity.SocialNetwork) error {
	login := ctx.Value(UserLogin).(string)

	result := s.db.Model(&entity.SocialNetwork{}).Where("id = ? AND owner_login = ?", social.ID, login).Updates(social)

	if result.Error != nil {
		log.Err(result.Error).Interface("social", social).Msg("failed update social in db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrFailedUpdate
	}

	return nil
}

func (s *socialImpl) Delete(ctx context.Context, ID uuid.UUID) error {
	login := ctx.Value(UserLogin).(string)

	social := &entity.SocialNetwork{}

	result := s.db.Where("id = ? AND owner_login = ?", ID, login).Delete(social)

	if result.Error != nil {
		log.Err(result.Error).Str("social_id", ID.String()).Msg("failed delete social from db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
