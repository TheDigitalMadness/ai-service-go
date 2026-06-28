package service

import (
	openai_client "github.com/TheDigitalMadness/ai-service-go/internal/client/openai"
)

type service struct {
	OpenAI *openai_client.OpenAI
}

func New(openai *openai_client.OpenAI) *service {
	return &service{OpenAI: openai}
}
