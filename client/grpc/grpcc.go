package grpc

import (
	"go-kit-demo/pb"
	controller "go-kit-demo/pkg/endpoint"
	"go-kit-demo/pkg/server"
	"go-kit-demo/services"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func New(conn *grpc.ClientConn) services.AddService {
	sumEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Sum",
		server.EncodeGRPCSumRequest,
		server.DecodeGRPCSumResponse,
		pb.SumReply{},
	).Endpoint()

	concatEndpoint := grpctransport.NewClient(
		conn, "pb.Add", "Concat",
		server.EncodeGRPCConcatRequest,
		server.DecodeGRPCConcatResponse,
		pb.ConcatReply{},
	).Endpoint()

	return controller.Endpoints{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}
}
