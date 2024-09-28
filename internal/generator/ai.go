package generator

import (
	"ai-feed/internal/entity"
	"bytes"
	"context"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"text/template"
)

type AI struct {
	client *openai.Client

	textPrompt  *template.Template
	imagePrompt *template.Template
	titlePrompt *template.Template

	textModel  string
	imageModel string
}

func NewAI(cfg *Config) *AI {
	defaultConfig := openai.DefaultConfig(cfg.OpenAiAuthToken)
	defaultConfig.BaseURL = cfg.OpenAiEndpoint

	client := openai.NewClientWithConfig(defaultConfig)

	textPrompt, err := template.ParseFiles(cfg.TextPromptPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed read text prompt")
	}

	imagePrompt, err := template.ParseFiles(cfg.ImagePromptPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed read image prompt")
	}

	titlePrompt, err := template.ParseFiles(cfg.TitlePromptPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed read image prompt")
	}

	ai := &AI{
		client:      client,
		textModel:   cfg.TextModel,
		imageModel:  cfg.ImageModel,
		textPrompt:  textPrompt,
		imagePrompt: imagePrompt,
		titlePrompt: titlePrompt,
	}

	return ai
}

// GenerateArticle generates articles using OpenAI LLM textModel using theme and personality.
func (ai *AI) GenerateArticle(ctx context.Context, theme *entity.Theme, personality *entity.Personality) (*entity.Article, error) {
	log.Info().Timestamp().
		Str("personality_id", personality.ID.String()).
		Msg("creating new article")

	text, err := ai.generateArticleText(ctx, personality, theme)
	if err != nil {
		log.Err(err).Msg("failed create article text")
		return nil, err
	}

	title, err := ai.generateArticleTitle(ctx, theme)
	if err != nil {
		log.Err(err).Msg("failed create article title")
		return nil, err
	}

	return &entity.Article{
		Title:   title,
		Content: text,
	}, nil
}

func (ai *AI) generateArticleText(ctx context.Context, personality *entity.Personality, theme *entity.Theme) (string, error) {
	var textBuffer bytes.Buffer

	request := &struct {
		*entity.Personality
		*entity.Theme
	}{
		personality, theme,
	}

	err := ai.textPrompt.Execute(&textBuffer, request)
	if err != nil {
		log.Err(err).Msg("failed execute text prompt template")
		return "", err
	}

	textRequest := openai.ChatCompletionRequest{
		Model: ai.textModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: textBuffer.String(),
			},
		},
		MaxTokens:   1024,
		Temperature: 0.5,
		N:           1,
	}

	resp, err := ai.client.CreateChatCompletion(ctx, textRequest)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", ErrNoChoices
	}

	return resp.Choices[0].Message.Content, nil
}

// GenerateArticleImage generates prompt for generate article image, then generates image, using this prompt.
func (ai *AI) GenerateArticleImage(ctx context.Context, article *entity.Article) (string, error) {
	var imagePrompt bytes.Buffer

	err := ai.imagePrompt.Execute(&imagePrompt, article)
	if err != nil {
		log.Err(err).Msg("failed execute image prompt template")
		return "", err
	}

	textRequest := openai.ChatCompletionRequest{
		Model: ai.textModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: imagePrompt.String(),
			},
		},
		MaxTokens:   1024,
		Temperature: 0.5,
		N:           1,
	}

	resp, err := ai.client.CreateChatCompletion(ctx, textRequest)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", ErrNoChoices
	}

	prompt := resp.Choices[0].Message.Content
	log.Info().Msgf("generating image with prompt: %s and model %s", prompt, ai.imageModel)

	imageRequest := openai.ImageRequest{
		Model:          ai.imageModel,
		Prompt:         prompt,
		N:              1,
		Quality:        openai.CreateImageQualityStandard,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}

	respImage, err := ai.client.CreateImage(ctx, imageRequest)
	if err != nil {
		return "", err
	}

	if len(respImage.Data) == 0 {
		return "", ErrNoChoices
	}

	return "data:image/png;base64," + respImage.Data[0].B64JSON, nil
}

func (ai *AI) generateArticleTitle(ctx context.Context, theme *entity.Theme) (string, error) {
	var titleBuffer bytes.Buffer

	err := ai.titlePrompt.Execute(&titleBuffer, theme)
	if err != nil {
		log.Err(err).Msg("failed execute title prompt template")
		return "", err
	}

	titleRequest := openai.ChatCompletionRequest{
		Model: ai.textModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: titleBuffer.String(),
			},
		},
		MaxTokens:   1024,
		Temperature: 0.5,
		N:           1,
	}

	resp, err := ai.client.CreateChatCompletion(ctx, titleRequest)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", ErrNoChoices
	}

	return resp.Choices[0].Message.Content, nil
}
