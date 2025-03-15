package handlers

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/entity/request"
	"ai-feed/internal/storage"
	"ai-feed/templates/views"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// @Summary		CreateSocialNetwork
// @Description	Creates social network
// @Security		header
// @Tags			social
// @ID				create-social
// @Accept			json
// @Produce			text/plain
// @Param			input	body	entity.SocialNetwork	true	"social information"
// @Router			/social [post]
func (h *HTTP) CreateSocial(c fiber.Ctx) error {
	social := &entity.SocialNetwork{}

	if err := c.Bind().JSON(social); err != nil {
		return err
	}

	err := h.service.CreateSocial(c.UserContext(), social)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed create personality")
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}

// @Summary		ReadAllSocials
// @Description	Read all personalities
// @Security		header
// @Tags			social
// @ID				read-socials
// @Produce		json
// @Router			/social [get]
func (h *HTTP) ReadAllSocials(c fiber.Ctx) error {
	socials, err := h.service.ReadAllSocials(c.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed get socials")
	}

	return c.Status(fiber.StatusOK).JSON(socials)
}

// @Summary		UpdateSocial
// @Description	Updates social
// @Security		header
// @Tags			social
// @ID				update-social
// @Accept			json
// @Produce		text/plain
// @Param			input	body	entity.Social	true	"social updated information"
// @Router			/social [put]
func (h *HTTP) UpdateSocial(c fiber.Ctx) error {
	social := &entity.SocialNetwork{}

	if err := c.Bind().JSON(social); err != nil {
		return err
	}

	err := h.service.UpdateSocial(c.UserContext(), social)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed update social")
	}

	return c.Status(fiber.StatusAccepted).Send(nil)
}

// @Summary		DeleteSocial
// @Description	Deletes social
// @Security		header
// @Tags			social
// @ID				delete-social
// @Accept			text/plain
// @Produce		text/plain
// @Param			input	body	string	true	"social ID"
// @Router			/social [delete]
func (h *HTTP) DeleteSocial(c fiber.Ctx) error {
	id, err := uuid.ParseBytes(c.Body())
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.service.DeleteSocial(c.UserContext(), id)

	if errors.Is(err, storage.ErrNotFound) {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed delete social")
	}

	return c.Status(fiber.StatusOK).Send(nil)
}

// @Summary		PublishSocial
// @Description	Publishes to social
// @Security		header
// @Tags			social
// @ID				publish-social
// @Accept			json
// @Produce		    json
// @Router			/social/publish [post]
func (h *HTTP) PublishSocial(c fiber.Ctx) error {
	social := &request.PublishSocial{}

	if err := c.Bind().JSON(social); err != nil {
		return err
	}

	err := h.service.PublishToSocial(c.UserContext(), social.SocialID, social.ArticleID, social.PublishTimestamp)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusAccepted).Send(nil)
}

func (h *HTTP) GetSocialsPage(c fiber.Ctx) error {
	socials, err := h.service.ReadAllSocials(c.UserContext())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	templSocials := make([]*views.SocialNetwork, 0, len(socials))

	for _, el := range socials {
		templSocials = append(templSocials, &views.SocialNetwork{
			ID:         el.ID,
			Name:       el.Name,
			OwnerLogin: el.OwnerLogin,
			Type:       el.Type,
			APIKey:     el.APIKey,
			ChannelID:  el.ChannelID,
		})
	}

	c.Set("Content-Type", "text/html")

	return views.NewSocials(templSocials).Render(c.Context(), c.Response().BodyWriter())
}
