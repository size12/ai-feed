package storage

import (
	"ai-feed/internal/entity"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
)

// Theme is interface for CRUD storage for themes.
type Theme interface {
	Create(ctx context.Context, theme *entity.Theme) error
	ReadAll(ctx context.Context) ([]*entity.Theme, error)
	Update(ctx context.Context, theme *entity.Theme) error
	Delete(ctx context.Context, ID uuid.UUID) error

	// CreateWithCheckDescription helps not to put themes with same description from feeder service.
	CreateWithCheckDescription(ctx context.Context, theme *entity.Theme) error
}

func NewTheme(db *gorm.DB, cfg *Config) Theme {
	return newThemeImpl(db, cfg)
}

// themeImpl is implementation of Theme interface
type themeImpl struct {
	db  *gorm.DB
	cfg *Config
}

func newThemeImpl(db *gorm.DB, cfg *Config) *themeImpl {
	impl := &themeImpl{
		db:  db,
		cfg: cfg,
	}

	go impl.run()

	return impl
}

func (t *themeImpl) Create(ctx context.Context, theme *entity.Theme) error {
	if theme.ID.ID() == 0 {
		theme.ID = uuid.New()
	}

	result := t.db.Model(&entity.Theme{}).Create(theme)

	if result.Error != nil {
		log.Err(result.Error).Interface("theme", theme).Msg("failed add theme to db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotCreated
	}

	return nil
}

func (t *themeImpl) ReadAll(ctx context.Context) ([]*entity.Theme, error) {
	var themes []*entity.Theme

	result := t.db.Model(&entity.Theme{}).Order("created_at DESC").Where("deleted = FALSE").Find(&themes)

	if result.Error != nil {
		log.Err(result.Error).Msg("failed read all themes from db")
		return nil, result.Error
	}

	return themes, nil
}

func (t *themeImpl) Update(ctx context.Context, theme *entity.Theme) error {
	result := t.db.Model(&entity.Theme{}).Where("id = ?", theme.ID).Updates(theme)

	if result.Error != nil {
		log.Err(result.Error).Interface("theme", theme).Msg("failed update theme in db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrFailedUpdate
	}

	return nil
}

func (t *themeImpl) Delete(ctx context.Context, ID uuid.UUID) error {
	result := t.db.Model(&entity.Theme{}).Where("id = ?", ID).Update("deleted", true)

	if result.Error != nil {
		log.Err(result.Error).Str("theme_id", ID.String()).Msg("failed delete theme from db")
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (t *themeImpl) CreateWithCheckDescription(ctx context.Context, theme *entity.Theme) error {
	var count int64 = 0

	result := t.db.Model(&entity.Theme{}).Where("description = ?", theme.Description).Count(&count)
	if result.Error != nil {
		log.Err(result.Error).Interface("theme", theme).Msg("failed count themes")
		return result.Error
	}

	if count >= 1 {
		return ErrAlreadyExists
	}

	return t.Create(ctx, theme)
}

// run starts worker, which deletes old themes
func (t *themeImpl) run() {
	ticker := time.NewTicker(t.cfg.WorkerActualUpdate)

	for {
		<-ticker.C

		var ids []uuid.UUID

		err := t.db.Model(&entity.Theme{}).
			Where("deleted = FALSE").
			Order("created_at DESC").
			Limit(int(t.cfg.ThemesActualCount)).Pluck("id", &ids).Error

		if err != nil {
			log.Err(err).Msg("failed get newest themes to delete old ones")
			continue
		}

		fmt.Println(ids, t.cfg.ThemesActualCount)

		result := t.db.Model(&entity.Theme{}).
			Where("id NOT IN ?", ids).
			Where("deleted = FALSE").
			UpdateColumn("deleted", true)

		if result.Error != nil {
			log.Err(result.Error).Msg("failed delete old themes")
			continue
		}

		if result.RowsAffected > 0 {
			log.Info().Msgf("marked %v themes as old", result.RowsAffected)
		}
	}
}
