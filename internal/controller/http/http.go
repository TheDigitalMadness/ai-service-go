package http

import "context"

type HttpService interface {
	GetFindToursCriteries(ctx context.Context, userRequest string) (FindToursCriteriesResponse, error)
}

type handler struct {
	service HttpService
}

func New(service HttpService) *handler {
	return &handler{service: service}
}
