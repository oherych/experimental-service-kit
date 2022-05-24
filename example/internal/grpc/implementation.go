package grpc

import (
	"context"
	grpc "github.com/oherych/experimental-service-kit/example/internal/grpc/gen"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	grpc.UnimplementedUsersServer

	d locator.Implementation
}

func (cc Implementation) List(ctx context.Context, empty *emptypb.Empty) (*grpc.UserList, error) {
	users, err := cc.d.Users.All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*grpc.User, len(users))
	for i, item := range users {
		result[i] = cc.display(item)
	}

	return &grpc.UserList{Users: result}, nil
}

func (cc Implementation) GetByID(ctx context.Context, request *grpc.GetByIDRequest) (*grpc.User, error) {
	user, err := cc.d.Users.GetByID(ctx, request.GetId())
	if err == repository.ErrUserNotFound {
		return nil, status.Error(codes.NotFound, "user found")
	}
	if err != nil {
		return nil, err
	}

	return cc.display(*user), nil
}

func (cc Implementation) display(in repository.User) *grpc.User {
	return &grpc.User{
		Id:       toPointID(in.ID),
		Username: &in.Username,
		Email:    &in.Email,
	}
}

// TODO: move me
func toPointID(in int) *int32 {
	val := int32(in)
	return &val
}
