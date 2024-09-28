package handlers

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/service"
	"ai-feed/templates/views"
	"errors"
	"github.com/gofiber/fiber/v3"
)

func (h *HTTP) AuthUser(c fiber.Ctx) error {
	credentials := &entity.User{}

	if err := c.Bind().Body(credentials); err != nil {
		return err
	}

	token, err := h.service.AuthUser(credentials)

	if errors.Is(err, service.ErrLoginAlreadyExists) {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "error",
			"msg":    "Пара логин/пароль неверная",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "error",
			"msg":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "ok",
		"token":  token,
	})
}

func (h *HTTP) GetAuthPage(c fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	return views.NewAuth().Render(c.Context(), c.Response().BodyWriter())
}
