package application

import (
	"context"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/rs/zerolog"
)

type Port[Dep dependencies.Locator] interface {
	Server(ctx context.Context, log zerolog.Logger, dep Dep, bc dependencies.BaseConfig) error
}
