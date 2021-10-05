package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/model/mapping"
	pb "github.com/ducnguyen96/ducnguyen96.xyz-protos/protogen/v1"

	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/generated"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/gateway/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.UserRegisterInput) (*model.RegisterPayload, error) {
	fmt.Println("input :", input)
	// validation here

	res, err := r.AuthClient.Register(ctx, &pb.UserRegisterInput{
		Username:       input.Username,
		Password:       input.Password,
		RepeatPassword: input.RepeatPassword,
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &model.RegisterPayload{
		User:  mapping.UserEntity(res.User),
		Token: mapping.TokenPayloadDto(res.Token),
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.UserLoginInput) (*model.TokenPayloadDto, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
