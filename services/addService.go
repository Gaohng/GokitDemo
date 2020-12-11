package services

import "context"

type AddService interface {
	Sum(_ context.Context, a, b int) (v int)
	Concat(_ context.Context, a, b string) (v string)
}
type addService struct {
}

func NewAddServices() AddService {
	return addService{}
}

func (addService) Sum(_ context.Context, a, b int) (v int) {
	return a + b
}

func (addService) Concat(_ context.Context, a, b string) (v string) {
	return a + b
}
