package service

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/TheDigitalMadness/ai-service-go/internal/domain"
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

func (s *service) extractPositionalFromAIResponse(props ResponseProperties) []domain.FindTourCriteria {
	positional := make([]domain.FindTourCriteria, 0)

	propsToSkip := map[string]bool{
		"AnswerType":          true,
		"ShortAnswer":         true,
		"DescriptionKeyWords": true,
	}

	v := reflect.ValueOf(props)
	t := reflect.TypeOf(props)

	for i := 0; i < v.NumField(); i++ {
		prop := t.Field(i)
		value := v.Field(i)

		if propsToSkip[prop.Name] {
			continue
		}

		if value.Kind() == reflect.Ptr && value.IsNil() {
			continue
		}

		value_Value := value.Elem().FieldByName("Value")
		if !value_Value.IsValid() {
			continue
		}

		jsonTag := prop.Tag.Get("json")

		positional = append(positional, domain.FindTourCriteria{
			Name:  jsonTag,
			Value: value_Value.Interface(),
		})
	}

	return positional
}

// GetFindToursCriteries returns criteries to find relevant tours
func (s *service) GetFindToursCriteries(ctx context.Context, userRequest string) (domain.FindToursCriteriesResponse, error) {
	var response domain.FindToursCriteriesResponse
	var lastErr error

	sleep := func(seconds int) {
		time.Sleep(time.Duration(seconds) * time.Second)
	}

	for i := 0; i < 3; i++ {
		responseFromAI, err := s.getResponseFromAI(ctx, userRequest)
		if err != nil {
			lastErr = err

			sleep(i)
			continue
		}

		responseFromAI_json := ResponseFromAI{}
		if err := json.Unmarshal([]byte(responseFromAI), &responseFromAI_json); err != nil {
			lastErr = err

			sleep(i)
			continue
		}

		response.AnswerType = domain.AnswerType(responseFromAI_json.Properties.AnswerType.Value)
		response.ShortAnswer = responseFromAI_json.Properties.ShortAnswer.Value

		if responseFromAI_json.Properties.DescriptionKeyWords != nil {
			response.KeyWords = responseFromAI_json.Properties.DescriptionKeyWords.Value
		} else {
			response.KeyWords = make([]string, 0)
		}

		response.Positional = s.extractPositionalFromAIResponse(responseFromAI_json.Properties)

		return response, nil
	}

	response.AnswerType = domain.AnswerTypeOnlyAnswer
	response.ShortAnswer = "К сожалению не получилось подобрать экскурсии по вашему запросу из-за внутренней ошибки сервиса, попробуйте, пожалуйста, позже. Вот экскурсии, которые могли бы вас заинтересовать"
	response.KeyWords = make([]string, 0)
	response.Positional = make([]domain.FindTourCriteria, 0)

	return response, lastErr
}
