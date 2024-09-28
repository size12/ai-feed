package service

import (
	"ai-feed/internal/generator"
	"ai-feed/internal/storage"
)

type Config struct {
	Themes        storage.Theme
	Articles      storage.Article
	Personalities storage.Personality
	Users         storage.User
	Ai            *generator.AI

	AuthConfig *AuthConfig `yaml:"auth"`
}
