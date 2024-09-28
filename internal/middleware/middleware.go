package middleware

import "ai-feed/internal/service"

type Middleware struct {
	service *service.AiFeed
}

func NewMiddleware(s *service.AiFeed) *Middleware {
	return &Middleware{
		service: s,
	}
}
