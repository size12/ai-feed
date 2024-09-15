package app

import (
	"ai-feed/internal/handlers"
	"ai-feed/internal/service/feeder"
	"ai-feed/templates/views"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/rs/zerolog/log"
)

type AiFeed struct {
	handlers *handlers.HTTP
	cfg      *Config

	app    *fiber.App
	feeder *feeder.Service
}

func NewAiFeed(cfg *Config, h *handlers.HTTP, f *feeder.Service) *AiFeed {
	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
		BodyLimit:       64 * 1024 * 1024,
	})

	app.Use(logger.New(), recover.New())

	app.Get("/personalities", h.GetPersonalitiesPage)
	app.Get("/themes", h.GetThemesPage)
	app.Get("/articles", h.GetArticlesPage)

	app.Get("/article/:id", h.GetArticle)

	api := app.Group("/api")

	api.Post("/generate/article", h.GenerateArticle)
	api.Post("/generate/image", h.GenerateArticleImage)

	article := api.Group("/article")

	article.Post("/", h.CreateArticle)
	article.Get("/:id?", h.ReadArticles)
	article.Put("/", h.UpdateArticle)
	article.Delete("/", h.DeleteArticle)

	personality := api.Group("/personality")

	personality.Post("/", h.CreatePersonality)
	personality.Get("/", h.ReadAllPersonalities)
	personality.Put("/", h.UpdatePersonality)
	personality.Delete("/", h.DeletePersonality)

	theme := api.Group("/theme")

	theme.Post("/", h.CreateTheme)
	theme.Get("/", h.ReadAllThemes)
	theme.Put("/", h.UpdateTheme)
	theme.Delete("/", h.DeleteTheme)

	app.Use(func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return views.NotFound().Render(c.Context(), c.Response().BodyWriter())
	})

	return &AiFeed{
		handlers: h,
		cfg:      cfg,
		app:      app,
		feeder:   f,
	}
}

func (app *AiFeed) Run(ctx context.Context) {
	go func() {
		<-ctx.Done()
		err := app.app.Shutdown()
		if err != nil {
			log.Fatal().Err(err).Msg("failed shutdown by context")
		}
	}()

	go func() {
		app.feeder.Run(ctx)
	}()

	log.Info().Msgf("Started server on %s", app.cfg.RunPort)

	err := app.app.Listen(app.cfg.RunPort)
	if err != nil {
		log.Fatal().Err(err).Msg("server shutdown")
	}
}
