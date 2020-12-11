package main

import (
	"flag"
	"fmt"
	"net"

	"go-kit-demo/pb"
	controller "go-kit-demo/pkg/endpoint"
	"go-kit-demo/pkg/server"
	"go-kit-demo/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func main() {
	httpAddr := flag.String("HTTP", ":8890", "HTTP server")
	gRPCAddr := flag.String("gRPC", ":8891", "gRPC server")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	logger.Log("msg", "Server Start...")
	defer logger.Log("msg", "Closed")

	svc := services.NewAddServices()

	endpoints := controller.Endpoints{
		SumEndpoint:    controller.MakeSumEndpoint(svc),
		ConcatEndpoint: controller.MakeConcatEndpoint(svc),
	}

	// Error channel.
	errc := make(chan error)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport.
	go func() {
		logger := log.With(logger, "transport", "HTTP")
		logger.Log("addr", *httpAddr)

		handler := server.MakeHTTPHandler(endpoints)
		errc <- http.ListenAndServe(*httpAddr, handler)
	}()

	//gRPC transport.
	go func() {
		logger := log.With(logger, "transport", "gRPC")
		logger.Log("addr", *gRPCAddr)

		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := server.MakeGRPCServer(endpoints)
		s := grpc.NewServer()
		pb.RegisterAddServer(s, srv)
		errc <- s.Serve(listener)
	}()

	logger.Log("exit", <-errc)
}
