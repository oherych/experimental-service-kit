package grpc

import (
	"github.com/oherych/experimental-service-kit/example/internal/grpc/gen"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/kit"

	"google.golang.org/grpc"
)

var Router = kit.GRPC[locator.Implementation]{
	Builder: func(dep locator.Implementation, sr grpc.ServiceRegistrar) error {
		gen.RegisterUsersServer(sr, Implementation{d: dep})

		return nil
	},
}
