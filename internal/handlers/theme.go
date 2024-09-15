package handlers

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/storage"
	"ai-feed/templates/views"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *HTTP) CreateTheme(c fiber.Ctx) error {
	theme := &entity.Theme{}

	if err := c.Bind().JSON(theme); err != nil {
		return err
	}

	err := h.service.CreateTheme(c.Context(), theme)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed create theme")
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *HTTP) ReadAllThemes(c fiber.Ctx) error {
	themes, err := h.service.ReadAllThemes(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed get themes")
	}

	return c.Status(fiber.StatusOK).JSON(themes)
}

func (h *HTTP) UpdateTheme(c fiber.Ctx) error {
	theme := &entity.Theme{}

	if err := c.Bind().JSON(theme); err != nil {
		return err
	}

	err := h.service.UpdateTheme(c.Context(), theme)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed update theme")
	}

	return c.Status(fiber.StatusAccepted).Send(nil)
}

func (h *HTTP) DeleteTheme(c fiber.Ctx) error {
	id, err := uuid.ParseBytes(c.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.DeleteTheme(c.Context(), id)

	if errors.Is(err, storage.ErrNotFound) {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed delete theme")
	}

	return c.Status(fiber.StatusOK).Send(nil)
}

func (h *HTTP) GetThemesPage(c fiber.Ctx) error {
	themes, err := h.service.ReadAllThemes(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	templThemes := make([]*views.Theme, 0, len(themes))

	for _, theme := range themes {
		templThemes = append(templThemes, &views.Theme{
			ID:          theme.ID,
			Description: theme.Description,
		})
	}

	c.Set("Content-Type", "text/html")

	return views.NewThemes(templThemes).Render(c.Context(), c.Response().BodyWriter())
}
