package handlers

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/storage"
	"ai-feed/templates/views"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (h *HTTP) CreatePersonality(c fiber.Ctx) error {
	personality := &entity.Personality{}

	if err := c.Bind().JSON(personality); err != nil {
		return err
	}

	err := h.service.CreatePersonality(c.Context(), personality)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed create personality")
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

func (h *HTTP) ReadAllPersonalities(c fiber.Ctx) error {
	personalities, err := h.service.ReadAllPersonalities(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed get personalities")
	}

	return c.Status(fiber.StatusOK).JSON(personalities)
}

func (h *HTTP) UpdatePersonality(c fiber.Ctx) error {
	personality := &entity.Personality{}

	if err := c.Bind().JSON(personality); err != nil {
		return err
	}

	err := h.service.UpdatePersonality(c.Context(), personality)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed update personality")
	}

	return c.Status(fiber.StatusAccepted).Send(nil)
}

func (h *HTTP) DeletePersonality(c fiber.Ctx) error {
	id, err := uuid.ParseBytes(c.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.DeletePersonality(c.Context(), id)

	if errors.Is(err, storage.ErrNotFound) {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed delete personality")
	}

	return c.Status(fiber.StatusOK).Send(nil)
}

func (h *HTTP) GetPersonalitiesPage(c fiber.Ctx) error {
	personalities, err := h.service.ReadAllPersonalities(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
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

	return views.NewPersonalities(templPersonalities).Render(c.Context(), c.Response().BodyWriter())
}
