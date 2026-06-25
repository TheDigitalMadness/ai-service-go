package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/TheDigitalMadness/ai-service-go/internal/controller/http"
	"github.com/openai/openai-go/v3/packages/param"
	"github.com/openai/openai-go/v3/responses"
)

// getResponseFromAI makes request to OpenAI model and returns it's response
func (s *service) getResponseFromAI(ctx context.Context, userRequest string) (string, error) {
	response, err := s.OpenAI.Client.Responses.New(
		ctx,
		responses.ResponseNewParams{
			Prompt: responses.ResponsePromptParam{
				ID: s.OpenAI.PromptID,
			},
			Input: responses.ResponseNewParamsInputUnion{
				OfString: param.Opt[string]{
					Value: userRequest,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return response.OutputText(), nil
}

// GetFindToursCriteries returns criteries to find relevant tours
func (s *service) GetFindToursCriteries(ctx context.Context, userRequest string) (http.FindToursCriteriesResponse, error) {
	var response http.FindToursCriteriesResponse
	var lastErr error

	sleep := func(seconds int) {
		time.Sleep(time.Duration(seconds) * time.Second)
	}

	for i := 0; i < 3; i++ {
		responseFromAI, err := s.getResponseFromAI(ctx, userRequest)
		if err != nil {
			lastErr = err

			sleep(i + 1)
			continue
		}

		if err := json.Unmarshal([]byte(responseFromAI), &response); err != nil {
			lastErr = err

			sleep(i + 1)
			continue
		}

		return response, nil
	}

	return response, lastErr
}
