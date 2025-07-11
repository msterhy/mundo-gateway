package main

import (
	"context"
	gatewaySdk "github.com/trancecho/mundo-gateway-sdk"
	"github.com/trancecho/mundo-gateway/test/ping/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type serverB struct {
	grpcpingv1.UnimplementedPingServiceServer
}

func (s *serverB) Ping(ctx context.Context, req *grpcpingv1.PingRequest) (*grpcpingv1.PingResponse, error) {
	return &grpcpingv1.PingResponse{
		Message: "pong",
	}, nil
}

func main() {
	server := grpc.NewServer()
	grpcpingv1.RegisterPingServiceServer(server, &serverB{})
	client := gatewaySdk.NewGatewayService("ping", "grpc://localhost:50052", "grpc", "http://localhost:12388")
	log.Println("aa")
	client.GrpcConn(server)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Println("failed to listen:", err)
		return
	}
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Println("failed to serve:", err)
		return
	}
}
