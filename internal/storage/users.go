package storage

import (
	"ai-feed/internal/entity"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// ContextKey is type for context key.
type ContextKey string

var UserLogin ContextKey = "user_login"

// User is interface for create and read users in storage.
type User interface {
	Create(user *entity.User) error
	Check(user *entity.User) error
}

func NewUser(db *gorm.DB, cfg *Config) User {
	return newUserImpl(db, cfg)
}

// userImpl is implementation of User interface
type userImpl struct {
	db  *gorm.DB
	cfg *Config
}

func newUserImpl(db *gorm.DB, cfg *Config) *userImpl {
	impl := &userImpl{
		db:  db,
		cfg: cfg,
	}

	return impl
}

func (u *userImpl) Create(user *entity.User) error {
	result := u.db.Model(&entity.User{}).Create(user)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return ErrAlreadyExists
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *userImpl) Check(user *entity.User) error {
	var valid bool

	result := u.db.Model(&entity.User{}).
		Select("true").
		Where("login = ? AND password = ?", user.Login, user.Password).
		Scan(&valid)

	fmt.Println(result.Error, result.RowsAffected, valid)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || !valid {
		return ErrUnauthorized
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
