package grpc

import (
	"context"
	_ "embed"
	generated "github.com/oherych/experimental-service-kit/example/internal/grpc/proto/business_domain/v1"
	"github.com/oherych/experimental-service-kit/pkg/grpc-listener"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"google.golang.org/grpc"
)

//go:embed generated/users.swagger.json
var swaggerContent []byte

var Router = grpc_listener.GRPC[locator.Config, *locator.Container]{
	Swagger: swaggerContent,
	HTTPHandlers: []func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error{
		generated.RegisterUsersServiceHandlerFromEndpoint,
	},
	Builder: func(sr grpc.ServiceRegistrar, dep *locator.Container) error {
		generated.RegisterUsersServiceServer(sr, Implementation{d: dep})

		return nil
	},
}
