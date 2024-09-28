package config

import (
	"ai-feed/internal/app"
	"ai-feed/internal/generator"
	"ai-feed/internal/service/feeder"
	"ai-feed/internal/storage"
	"flag"
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Config struct {
	App     *app.Config       `yaml:"app"`
	AI      *generator.Config `yaml:"ai"`
	Storage *storage.Config   `yaml:"storage"`
	Feeder  *feeder.Config    `yaml:"feeder"`
}

var (
	once = &sync.Once{}
	cfg  = &Config{}
)

func ParseConfig() *Config {
	once.Do(func() {
		configFilePath := flag.String("cfg", "./config.yaml", "Configuration file. See example_config.yaml as example")
		flag.Parse()

		configFile, err := os.ReadFile(*configFilePath)
		if err != nil {
			log.Fatal().Err(err).Msg("failed open config file")
		}

		err = yaml.Unmarshal(configFile, cfg)
		if err != nil {
			log.Fatal().Err(err).Msg("failed parse config file")
		}

		if err = validateConfig(cfg); err != nil {
			log.Fatal().Err(err).Msg("fonfig is not valid")
		}
	})

	return cfg
}

func validateConfig(cfg *Config) error {
	if cfg.AI == nil {
		return fmt.Errorf("AI config is empty")
	}

	if cfg.AI.TextModel == "" {
		return fmt.Errorf("missing model type")
	}

	if cfg.AI.OpenAiEndpoint == "" {
		return fmt.Errorf("missing open ai endpoint url")
	}

	if cfg.AI.OpenAiAuthToken == "" {
		return fmt.Errorf("missing open ai auth token")
	}

	return nil
}
