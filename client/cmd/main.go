package main

import (
	"context"
	"flag"
	grpclient "go-kit-demo/client/grpc"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	gRPCAddr := flag.String("gRPC", ":8891", "gRPC client")
	flag.Parse()

	conn, err := grpc.Dial(
		*gRPCAddr, grpc.WithInsecure(),
		grpc.WithTimeout(time.Second),
	)

	if err != nil {
		log.Fatalln("gRPC dial error:", err)
	}
	defer conn.Close()

	addService := grpclient.New(conn)

	println("Sum Response:")
	output := addService.Sum(context.Background(), 11111, 22222)
	println(output)
	println("Concat Response:")
	outputS := addService.Concat(context.Background(), "11111", "22222")
	println(outputS)
}
