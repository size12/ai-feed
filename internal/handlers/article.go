package handlers

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/storage"
	"ai-feed/templates/views"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"html/template"
	"strings"
)

//	@Summary		CreateArticle
//	@Description	Saves article's text and image
//	@Security		header
//	@Tags			article
//	@ID				create-article
//	@Accept			json
//	@Produce		text/plain
//	@Param			input	body	entity.Article	true	"article information"
//	@Router			/article [post]
func (h *HTTP) CreateArticle(c fiber.Ctx) error {
	article := &entity.Article{}

	if err := c.Bind().JSON(article); err != nil {
		return err
	}

	err := h.service.CreateArticle(c.UserContext(), article)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed create article")
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

//	@Summary		ReadArticles
//	@Description	Read all articles or one by ID
//	@Security		header
//	@Tags			article
//	@ID				read-articles
//	@Accept			text/plain
//	@Produce		json
//	@Param			id	query	string	false	"article ID"
//	@Router			/article [get]
func (h *HTTP) ReadArticles(c fiber.Ctx) error {
	idStr := c.Query("id")
	c.Set("Content-Type", "application/json")

	fmt.Println(c.UserContext().Value(storage.UserLogin))

	if idStr == "" {
		articles, err := h.service.ReadAllArticles(c.UserContext())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed get articles")
		}

		return c.Status(fiber.StatusOK).JSON(articles)
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return fiber.ErrBadRequest
	}

	article, err := h.service.ReadArticle(c.UserContext(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed get articles")
	}

	return c.Status(fiber.StatusOK).JSON(article)
}

//	@Summary		UpdateArticle
//	@Description	Updates article's text and image
//	@Security		header
//	@Tags			article
//	@ID				update-article
//	@Accept			json
//	@Produce		text/plain
//	@Param			input	body	entity.Article	true	"new article information"
//	@Router			/article [put]
func (h *HTTP) UpdateArticle(c fiber.Ctx) error {
	article := &entity.Article{}

	if err := c.Bind().JSON(article); err != nil {
		return err
	}

	err := h.service.UpdateArticle(c.UserContext(), article)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed update article")
	}

	return c.Status(fiber.StatusAccepted).Send(nil)
}

//	@Summary		DeleteArticle
//	@Description	Delete article by ID
//	@Security		header
//	@Tags			article
//	@ID				delete-articles
//	@Accept			text/plain
//	@Produce		text/plain
//	@Param			input	body	string	true	"article ID"
//	@Router			/article [delete]
func (h *HTTP) DeleteArticle(c fiber.Ctx) error {
	id, err := uuid.ParseBytes(c.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.DeleteArticle(c.UserContext(), id)

	if errors.Is(err, storage.ErrNotFound) {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed delete article")
	}

	return c.Status(fiber.StatusOK).Send(nil)
}

func (h *HTTP) GetArticlePage(c fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return views.NotFoundArticle().Render(c.UserContext(), c.Response().BodyWriter())
	}

	article, err := h.service.ReadArticle(c.UserContext(), id)

	if errors.Is(err, storage.ErrNotFound) {
		return views.NotFoundArticle().Render(c.UserContext(), c.Response().BodyWriter())
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed delete article")
	}

	content := template.Must(template.New("content").Parse(strings.ReplaceAll(article.Content, "\n", "<br>")))

	articleTempl := &views.ShownArticle{
		Title:       article.Title,
		ImageBase64: article.ImageBase64,
		Content:     content,
	}

	return views.NewShownArticle(articleTempl).Render(c.UserContext(), c.Response().BodyWriter())
}

func (h *HTTP) GetArticlesPage(c fiber.Ctx) error {
	articles, err := h.service.ReadAllArticles(c.UserContext())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	themes, err := h.service.ReadAllThemes(c.UserContext())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	feederThemes, err := h.service.ReadFeederThemes(c.UserContext())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	personalities, err := h.service.ReadAllPersonalities(c.UserContext())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	templArticles := make([]*views.Article, 0, len(articles))

	for _, article := range articles {
		keywords := make([]*views.Keyword, 0, len(article.Keywords))

		for _, el := range article.Keywords {
			keywords = append(keywords, &views.Keyword{
				Name:  el.Name,
				Count: fmt.Sprint(el.Count),
			})
		}

		templArticles = append(templArticles, &views.Article{
			ID:           article.ID,
			Title:        article.Title,
			ImageBase64:  article.ImageBase64,
			Content:      article.Content,
			WordsCount:   fmt.Sprint(article.WordsCount),
			SymbolsCount: fmt.Sprint(article.SymbolsCount),
			Keywords:     keywords,
		})
	}

	templThemes := make([]*views.Theme, 0, len(themes))

	for _, theme := range themes {
		templThemes = append(templThemes, &views.Theme{
			ID:          theme.ID,
			Description: theme.Description,
		})
	}

	templFeederThemes := make([]*views.Theme, 0, len(feederThemes))

	for _, theme := range feederThemes {
		templFeederThemes = append(templFeederThemes, &views.Theme{
			ID:          theme.ID,
			Description: theme.Description,
		})
	}

	templPersonalities := make([]*views.Personality, 0, len(personalities))

	for _, el := range personalities {
		templPersonalities = append(templPersonalities, &views.Personality{
			ID:        el.ID,
			Name:      el.Name,
			Biography: el.Biography,
			Keywords:  el.Keywords,
			Thematics: el.Thematics,
			TextStyle: el.TextStyle,
		})
	}

	c.Set("Content-Type", "text/html")

	return views.NewArticles(templArticles, templThemes, templFeederThemes, templPersonalities).Render(c.UserContext(), c.Response().BodyWriter())
}
