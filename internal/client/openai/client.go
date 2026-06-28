package openai_client

import (
	"github.com/TheDigitalMadness/ai-service-go/internal/config"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// OpenAI contains reference on created openai.Client object and ID of built-in prompt in AI model
type OpenAI struct {
	Client   *openai.Client
	PromptID string
}

func New(cfg *config.Config) *OpenAI {
	client := openai.NewClient(
		option.WithBaseURL(cfg.AI.BaseURL),
		option.WithAPIKey(cfg.AI.ApiKey),
		option.WithProject(cfg.AI.Project),
	)

	return &OpenAI{
		Client:   &client,
		PromptID: cfg.AI.PromptID,
	}
}
