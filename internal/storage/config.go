package storage

import "time"

type Config struct {
	PostgresEndpoint   string        `yaml:"postgres_endpoint"`
	ThemesActualCount  uint          `yaml:"themes_actual_count"`
	WorkerActualUpdate time.Duration `yaml:"worker_actual_update"`
}
