package main

import (
	"ai-feed/internal/app"
	"ai-feed/internal/config"
	"ai-feed/internal/generator"
	"ai-feed/internal/handlers"
	"ai-feed/internal/service"
	"ai-feed/internal/service/feeder"
	"ai-feed/internal/storage"
	"context"
	"github.com/rs/zerolog/log"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.ParseConfig()

	log.Info().Interface("config", cfg).Msg("current config")

	ai := generator.NewAI(cfg.AI)

	db := storage.NewORM(cfg.Storage)

	themes := storage.NewTheme(db, cfg.Storage)
	personalities := storage.NewPersonality(db)
	articles := storage.NewArticle(db)

	s := service.NewAiFeed(themes, articles, personalities, ai)
	h := handlers.NewHTTP(s)

	f := feeder.NewService(cfg.Feeder, themes,
		feeder.VcRuFeeder,
		feeder.IxbtFeeder,
	)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGKILL, syscall.SIGSTOP)
	defer cancel()

	app.NewAiFeed(cfg.App, h, f).Run(ctx)
}
