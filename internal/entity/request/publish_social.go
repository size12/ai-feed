package request

import (
	"github.com/google/uuid"
)

type PublishSocial struct {
	SocialID         uuid.UUID `json:"social_id" validate:"required"`
	ArticleID        uuid.UUID `json:"article_id" validate:"required"`
	PublishTimestamp int64     `json:"publish_timestamp"`
}
