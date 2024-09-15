package storage

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Personality is interface for CRUD storage for personalities.
type Personality interface {
	Create(ctx context.Context, personality *entity.Personality) error
	ReadAll(ctx context.Context) ([]*entity.Personality, error)
	Update(ctx context.Context, personality *entity.Personality) error
	Delete(ctx context.Context, ID uuid.UUID) error
}

func NewPersonality(db *gorm.DB) Personality {
	return newPersonalityImpl(db)
}

// personalityImpl is implementation of Personality interface.
type personalityImpl struct {
	db *gorm.DB
}

func newPersonalityImpl(db *gorm.DB) *personalityImpl {
	return &personalityImpl{
		db: db,
	}
}

func (p *personalityImpl) Create(ctx context.Context, personality *entity.Personality) error {
	if personality.ID.ID() == 0 {
		personality.ID = uuid.New()
	}

	result := p.db.Create(personality)

	if result.Error != nil {
		log.Err(result.Error).Interface("personality", personality).Msg("failed add personality to db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotCreated
	}

	return nil
}

func (p *personalityImpl) ReadAll(ctx context.Context) ([]*entity.Personality, error) {
	var personalities []*entity.Personality

	result := p.db.Order("created_at DESC").Find(&personalities)

	if result.Error != nil {
		log.Err(result.Error).Msg("failed read all personalities from db")
		return nil, result.Error
	}

	return personalities, nil
}

func (p *personalityImpl) Update(ctx context.Context, personality *entity.Personality) error {
	result := p.db.Model(&entity.Personality{}).Where("id = ?", personality.ID).Updates(personality)

	if result.Error != nil {
		log.Err(result.Error).Interface("personality", personality).Msg("failed update personality in db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrFailedUpdate
	}

	return nil
}

func (p *personalityImpl) Delete(ctx context.Context, ID uuid.UUID) error {
	personality := &entity.Personality{}

	result := p.db.Where("id = ?", ID).Delete(personality)

	if result.Error != nil {
		log.Err(result.Error).Str("personality_id", ID.String()).Msg("failed delete personality from db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
