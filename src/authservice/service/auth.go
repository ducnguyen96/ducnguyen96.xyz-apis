package service

import (
	"context"
	"fmt"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	AuthClient pb.AuthServiceClient
}

func (c *AuthService) Register(ctx context.Context, input *pb.UserRegisterInput) (*pb.RegisterPayload, error) {
	out := new(pb.RegisterPayload)
	fmt.Println(input)
	return out, nil
}