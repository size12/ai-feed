package handlers

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/entity/request"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func (h *HTTP) GenerateArticle(c fiber.Ctx) error {
	r := &request.GenerateRequest{}

	if err := c.Bind().JSON(r); err != nil {
		return err
	}

	article, err := h.service.GenerateArticle(c.UserContext(), r.Theme, r.Personality)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed generate article")
	}

	return c.Status(fiber.StatusOK).JSON(article)
}

func (h *HTTP) GenerateArticleImage(c fiber.Ctx) error {
	article := &entity.Article{}

	if err := c.Bind().JSON(article); err != nil {
		return err
	}

	image, err := h.service.GenerateArticleImage(c.UserContext(), article)
	if err != nil {
		log.Err(err).Msg("failed generate image")
		return fiber.NewError(fiber.StatusInternalServerError, "failed generate image")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"image_base64": image,
	})
}
