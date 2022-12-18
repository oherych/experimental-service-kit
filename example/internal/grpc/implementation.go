package grpc

import (
	"context"
	"github.com/oherych/experimental-service-kit/example/internal/grpc/generated"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
	"github.com/oherych/experimental-service-kit/kit/logs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	d *locator.Container
}

func (cc Implementation) List(ctx context.Context, empty *emptypb.Empty) (*generated.UserList, error) {
	users, err := cc.d.Users.All(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*generated.User, len(users))
	for i, item := range users {
		result[i] = cc.display(item)
	}

	return &generated.UserList{Users: result}, nil
}

func (cc Implementation) Get(ctx context.Context, request *generated.GetByIDRequest) (*generated.User, error) {
	user, err := cc.d.Users.GetByID(ctx, request.GetId())
	if err == repository.ErrUserNotFound {
		return nil, status.Error(codes.NotFound, "user found")
	}
	if err != nil {
		return nil, err
	}

	return cc.display(*user), nil
}

func (cc Implementation) Delete(ctx context.Context, request *generated.GetByIDRequest) (*emptypb.Empty, error) {
	err := cc.d.Users.Delete(ctx, request.Id)
	if err == repository.ErrUserNotFound {
		return nil, status.Error(codes.NotFound, "user found")
	}
	if err != nil {
		return nil, err
	}

	logs.For(ctx).Log().Str("user_id", request.Id).Msg("user deleted")

	return &emptypb.Empty{}, nil
}

func (cc Implementation) display(in repository.User) *generated.User {
	return &generated.User{
		Id:       int32(in.ID),
		Username: in.Username,
		Email:    in.Email,
	}
}
