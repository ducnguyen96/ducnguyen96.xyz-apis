package main

import (
	"context"
	"fmt"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent"
	_ "github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/runtime"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/service"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	// pg database
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort) //Build connection string

	readClient, err := ent.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	writeClient, err := ent.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	// Run the auto migration tool.
	if err := readClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(readClient)


	// listening to tcp
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &service.AuthService{
		ReadDB:                         readClient.Debug(),
		WriteDB:                        writeClient.Debug(),
	})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}