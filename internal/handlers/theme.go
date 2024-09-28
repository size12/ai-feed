package handlers

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/storage"
	"ai-feed/templates/views"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

//	@Summary		CreateTheme
//	@Description	Creates theme
//	@Security		header
//	@Tags			theme
//	@ID				create-theme
//	@Accept			json
//	@Produce		text/plain
//	@Param			input	body	entity.Theme	true	"theme description"
//	@Router			/theme [post]
func (h *HTTP) CreateTheme(c fiber.Ctx) error {
	theme := &entity.Theme{}

	if err := c.Bind().JSON(theme); err != nil {
		return err
	}

	err := h.service.CreateTheme(c.UserContext(), theme)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed create theme")
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

//	@Summary		ReadUserThemes
//	@Description	Read user themes
//	@Security		header
//	@Tags			theme
//	@ID				read-themes
//	@Produce		json
//	@Router			/theme [get]
func (h *HTTP) ReadAllThemes(c fiber.Ctx) error {
	themes, err := h.service.ReadAllThemes(c.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed get themes")
	}

	return c.Status(fiber.StatusOK).JSON(themes)
}

//	@Summary		ReadFeederThemes
//	@Description	Read feeder themes
//	@Tags			theme
//	@ID				read-feeder-themes
//	@Produce		json
//	@Router			/theme/feeder [get]
func (h *HTTP) ReadFeederThemes(c fiber.Ctx) error {
	themes, err := h.service.ReadFeederThemes(c.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed get themes")
	}

	return c.Status(fiber.StatusOK).JSON(themes)
}

//	@Summary		UpdateTheme
//	@Description	Updated theme
//	@Security		header
//	@Tags			theme
//	@ID				update-theme
//	@Accept			json
//	@Produce		text/plain
//	@Param			input	body	entity.Theme	true	"theme new description"
//	@Router			/theme [put]
func (h *HTTP) UpdateTheme(c fiber.Ctx) error {
	theme := &entity.Theme{}

	if err := c.Bind().JSON(theme); err != nil {
		return err
	}

	err := h.service.UpdateTheme(c.UserContext(), theme)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed update theme")
	}

	return c.Status(fiber.StatusAccepted).Send(nil)
}

//	@Summary		DeleteTheme
//	@Description	Deletes theme
//	@Security		header
//	@Tags			theme
//	@ID				delete-theme
//	@Accept			text/plain
//	@Produce		text/plain
//	@Param			input	body	string	true	"theme ID"
//	@Router			/theme [delete]
func (h *HTTP) DeleteTheme(c fiber.Ctx) error {
	id, err := uuid.ParseBytes(c.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.DeleteTheme(c.UserContext(), id)

	if errors.Is(err, storage.ErrNotFound) {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed delete theme")
	}

	return c.Status(fiber.StatusOK).Send(nil)
}

func (h *HTTP) GetThemesPage(c fiber.Ctx) error {
	themes, err := h.service.ReadAllThemes(c.UserContext())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	feederThemes, err := h.service.ReadFeederThemes(c.UserContext())
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

	templFeederThemes := make([]*views.Theme, 0, len(feederThemes))

	for _, theme := range feederThemes {
		templFeederThemes = append(templFeederThemes, &views.Theme{
			ID:          theme.ID,
			Description: theme.Description,
		})
	}

	c.Set("Content-Type", "text/html")

	return views.NewThemes(templThemes, templFeederThemes).Render(c.UserContext(), c.Response().BodyWriter())
}
