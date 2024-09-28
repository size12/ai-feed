package generator

type Config struct {
	OpenAiEndpoint  string `yaml:"open_ai_endpoint"`
	OpenAiAuthToken string `yaml:"open_ai_auth_token"`
	TextModel       string `yaml:"text_model"`
	ImageModel      string `yaml:"image_model"`

	ImagePromptPath string `yaml:"image_prompt_path"`
	TextPromptPath  string `yaml:"text_prompt_path"`
	TitlePromptPath string `yaml:"title_prompt_path"`
}
