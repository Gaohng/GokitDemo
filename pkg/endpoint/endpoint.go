package controller

import (
	"context"
	"go-kit-demo/services"

	"github.com/go-kit/kit/endpoint"
)

type SumRequest struct {
	A int
	B int
}

type SumResponse struct {
	V int
}
type ConcatRequest struct {
	A string
	B string
}

type ConcatResponse struct {
	V string
}

//all endpoints required by AddService.
type Endpoints struct {
	SumEndpoint    endpoint.Endpoint
	ConcatEndpoint endpoint.Endpoint
}

// MakeSumEndpoint returns an endpoint that invokes Sum on the AddService
// for server
func MakeSumEndpoint(svc services.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v := svc.Sum(ctx, req.A, req.B)
		return SumResponse{v}, nil
	}
}

// Sum implements AddService
//for client
func (e Endpoints) Sum(ctx context.Context, a, b int) int {
	req := SumRequest{A: a, B: b}
	res, err := e.SumEndpoint(ctx, req)
	if err != nil {
		return SumResponse{0}.V
	}
	return res.(SumResponse).V
}

// MakeConcatEndpoint returns an endpoint that invokes Sum on the AddService
// for server
func MakeConcatEndpoint(svc services.AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatRequest)
		v := svc.Concat(ctx, req.A, req.B)
		return ConcatResponse{v}, nil
	}
}

// Concat implements AddService
//for client
func (e Endpoints) Concat(ctx context.Context, a, b string) string {
	req := ConcatRequest{a, b}
	res, err := e.ConcatEndpoint(ctx, req)
	if err != nil {
		return ConcatResponse{"err"}.V
	}
	return res.(ConcatResponse).V
}
