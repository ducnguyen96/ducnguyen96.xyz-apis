package main

import (
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/service"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)


func main() {
	// listening to tcp
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &service.AuthService{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}