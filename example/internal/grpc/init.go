package grpc

import (
	"context"
	_ "embed"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oherych/experimental-service-kit/example/internal/grpc/generated"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/kit"
	"google.golang.org/grpc"
)

//go:embed generated/users.swagger.json
var swaggerContent []byte

var Router = kit.GRPC[locator.Implementation]{
	Swagger: swaggerContent,
	HTTPHandlers: []func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error{
		generated.RegisterUsersServiceHandlerFromEndpoint,
	},
	Builder: func(sr grpc.ServiceRegistrar, dep locator.Implementation) error {
		generated.RegisterUsersServiceServer(sr, Implementation{d: dep})

		return nil
	},
}
