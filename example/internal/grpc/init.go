package grpc

import (
	"github.com/oherych/experimental-service-kit/example/internal/grpc/generated"
	"github.com/oherych/experimental-service-kit/kit/server"
	"github.com/oherych/experimental-service-kit/pkg/grpc-listener"

	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"google.golang.org/grpc"
)

func New() server.Listener[locator.Config, *locator.Container] {
	return grpc_listener.GRPC[locator.Config, *locator.Container]{
		InitConfig: func(conf locator.Config) grpc_listener.Config {
			return conf.GRPC
		},
		Builder: func(sr grpc.ServiceRegistrar, dep *locator.Container) error {
			generated.RegisterUsersServiceServer(sr, Implementation{d: dep})

			return nil
		},
	}
}
