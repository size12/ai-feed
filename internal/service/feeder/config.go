package feeder

import "time"

type Config struct {
	FeedUpdateDelay time.Duration `yaml:"feed_update_delay"`
}
