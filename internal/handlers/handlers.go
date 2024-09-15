package handlers

import (
	"ai-feed/internal/service"
)

type HTTP struct {
	service *service.AiFeed
}

func NewHTTP(feed *service.AiFeed) *HTTP {
	return &HTTP{
		service: feed,
	}
}
