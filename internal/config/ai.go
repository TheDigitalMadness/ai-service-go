package config

type AI struct {
	BaseURL  string `env:"OPENAI_BASE_URL"`
	ApiKey   string `env:"OPENAI_API_KEY"`
	Project  string `env:"OPENAI_PROJECT__TOUR_FOUNDER"`
	PromptID string `env:"OPENAI_PROMPT_ID__TOUR_FOUNDER"`
}
