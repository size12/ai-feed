package service

import (
	"ai-feed/internal/entity"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

func (service *AiFeed) CreateSocial(ctx context.Context, social *entity.SocialNetwork) error {
	log.Info().Interface("social", social).Msg("create social")
	return service.socials.Create(ctx, social)
}

func (service *AiFeed) ReadAllSocials(ctx context.Context) ([]*entity.SocialNetwork, error) {
	log.Info().Msg("read all social")
	return service.socials.ReadAll(ctx)
}

func (service *AiFeed) UpdateSocial(ctx context.Context, personality *entity.SocialNetwork) error {
	log.Info().Interface("personality", personality).Msg("update personality")
	return service.socials.Update(ctx, personality)
}

func (service *AiFeed) DeleteSocial(ctx context.Context, ID uuid.UUID) error {
	log.Info().Str("id", ID.String()).Msg("delete social")
	return service.socials.Delete(ctx, ID)
}

func (service *AiFeed) PublishToSocial(ctx context.Context, socialID, articleID uuid.UUID, timestamp int64) error {
	article, err := service.articles.Read(ctx, articleID)
	if err != nil {
		return err
	}

	social, err := service.socials.Read(ctx, socialID)
	if err != nil {
		return err
	}

	sendTime := time.UnixMilli(timestamp)

	log.Info().Str("social_id", socialID.String()).
		Str("article_id", articleID.String()).
		Interface("social", social).
		Interface("article", article).
		Time("ts", sendTime).
		Msg("publishing to social")

	switch social.Type {
	case "VK":
		return publishVK(article, social, sendTime)
	case "Telegram":
		go publishTelegram(article, social, sendTime)
		return nil
	default:
		return errors.New("unknown social network")
	}
}

func publishVK(article *entity.Article, social *entity.SocialNetwork, sendTime time.Time) error {
	channelIDint, err := strconv.ParseInt(social.ChannelID, 10, 64)
	if err != nil {
		return errors.New("bad group id")
	}

	b := params.NewWallPostBuilder()
	b.Message(article.Content)
	b.PublishDate(int(sendTime.Unix()))
	b.OwnerID(int(channelIDint))

	vk := api.NewVK(social.APIKey)
	_, err = vk.WallPost(b.Params)
	if err != nil {
		log.Err(err).Msg("failed post to vk")
		return err
	}

	return nil
}

func publishTelegram(article *entity.Article, social *entity.SocialNetwork, sendTime time.Time) {
	b, err := bot.New(social.APIKey)
	if err != nil {
		log.Err(err).Msg("bad api key")
	}

	now := time.Now().UTC()

	time.Sleep(sendTime.Sub(now))

	if article.ImageBase64 == "" {
		_, err = b.SendMessage(context.Background(), &bot.SendMessageParams{
			ChatID: social.ChannelID,
			Text:   article.Content,
		})

		if err != nil {
			log.Err(err).Msg("failed publish article in TG")
		}

		return
	}

	imageBase64 := strings.Split(article.ImageBase64, ",")

	if len(imageBase64) == 0 {
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(imageBase64[1])
	if err != nil {
		log.Err(err).Msg("failed decode base64 image")
	}

	photoReader := bytes.NewReader(imageData)

	_, err = b.SendPhoto(context.Background(), &bot.SendPhotoParams{
		ChatID: social.ChannelID,
		Photo: &models.InputFileUpload{
			Filename: "image.png",
			Data:     photoReader,
		},
		Caption: article.Content,
	})

	if err != nil {
		log.Err(err).Msg("failed send image")
	}
}
