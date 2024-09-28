package middleware

import (
	"ai-feed/internal/storage"
	"ai-feed/templates/views"
	"context"
	"github.com/gofiber/fiber/v3"
)

func (m *Middleware) AuthUser(c fiber.Ctx) error {
	token := c.Cookies("token")

	login, err := m.service.VerifyToken(token)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		c.Set("Content-Type", "text/html")

		return views.NewAuth().Render(c.Context(), c.Response().BodyWriter())
	}

	c.SetUserContext(context.WithValue(c.UserContext(), storage.UserLogin, login))

	return c.Next()
}
