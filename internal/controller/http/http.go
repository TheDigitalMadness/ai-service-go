package http

import (
	"context"

	"github.com/TheDigitalMadness/ai-service-go/internal/domain"
)

type HttpService interface {
	GetFindToursCriteries(ctx context.Context, userRequest string) (domain.FindToursCriteriesResponse, error)
}

type handler struct {
	service HttpService
}

func New(service HttpService) *handler {
	return &handler{service: service}
}
