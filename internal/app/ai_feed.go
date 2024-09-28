package app

import (
	"ai-feed/internal/handlers"
	"ai-feed/internal/middleware"
	"ai-feed/internal/service/feeder"
	"ai-feed/templates/views"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "ai-feed/docs"
)

type AiFeed struct {
	cfg *Config

	app    *fiber.App
	feeder *feeder.Service
}

// @title AI-feed API
// @BasePath  /api

func NewAiFeed(cfg *Config, h *handlers.HTTP, f *feeder.Service, m *middleware.Middleware) *AiFeed {
	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
		BodyLimit:       64 * 1024 * 1024,
	})

	app.Use(
		logger.New(),
		recover.New(),
	)

	app.Get("/auth", h.GetAuthPage)
	app.Post("/api/auth", h.AuthUser)
	app.Get("/article/:id", h.GetArticlePage)
	app.Get("api/theme/feeder", h.ReadFeederThemes)

	app.Get("/swagger/*", adaptor.HTTPHandler(
		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
		)))

	auth := app.Use(m.AuthUser)

	auth.Get("/personalities", h.GetPersonalitiesPage)
	auth.Get("/themes", h.GetThemesPage)
	auth.Get("/articles", h.GetArticlesPage)

	api := auth.Group("/api")

	api.Post("/generate/article", h.GenerateArticle)
	api.Post("/generate/image", h.GenerateArticleImage)

	article := api.Group("/article")

	article.Post("/", h.CreateArticle)
	article.Put("/", h.UpdateArticle)
	article.Delete("/", h.DeleteArticle)
	article.Get("/:id?", h.ReadArticles)

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
		c.Status(fiber.StatusNotFound)

		return views.NotFound().Render(c.Context(), c.Response().BodyWriter())
	})

	return &AiFeed{
		cfg:    cfg,
		app:    app,
		feeder: f,
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
