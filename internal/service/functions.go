package service

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
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

func (s *service) extractPositionalFromAIResponse(props ResponseProperties) []http.FindTourCriteria {
	positional := make([]http.FindTourCriteria, 0)

	propsToSkip := map[string]bool{
		"AnswerType":          true,
		"ShortAnswer":         true,
		"DescriptionKeyWords": true,
	}

	stringify := func(v any) string {
		return fmt.Sprintf("%v", v)
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

		positional = append(positional, http.FindTourCriteria{
			Name:  jsonTag,
			Value: stringify(value_Value.Interface()),
		})
	}

	return positional
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

		responseFromAI_json := ResponseFromAI{}
		if err := json.Unmarshal([]byte(responseFromAI), &responseFromAI_json); err != nil {
			lastErr = err

			sleep(i + 1)
			continue
		}

		response.AnswerType = http.AnswerType(responseFromAI_json.Properties.AnswerType.Value)
		response.ShortAnswer = responseFromAI_json.Properties.ShortAnswer.Value

		if responseFromAI_json.Properties.DescriptionKeyWords != nil {
			response.KeyWords = responseFromAI_json.Properties.DescriptionKeyWords.Value
		} else {
			response.KeyWords = make([]string, 0)
		}

		response.Positional = s.extractPositionalFromAIResponse(responseFromAI_json.Properties)

		return response, nil
	}

	response.AnswerType = http.AnswerTypeOnlyAnswer
	response.ShortAnswer = "К сожалению не получилось подобрать экскурсии по вашему запросу из-за внутренней ошибки сервиса, попробуйте, пожалуйста, позже. Вот экскурсии, которые могли бы вас заинтересовать"
	response.KeyWords = make([]string, 0)
	response.Positional = make([]http.FindTourCriteria, 0)

	return response, lastErr
}
