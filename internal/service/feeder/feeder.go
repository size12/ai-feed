package feeder

import (
	"ai-feed/internal/entity"
	"context"
)

// Feeder is func, which parses themes from sites.
type Feeder func(ctx context.Context) ([]*entity.Theme, error)
