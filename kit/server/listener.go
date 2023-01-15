package server

import (
	"context"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/rs/zerolog"
)

type Listener[Conf dependencies.Config, Dep dependencies.Locator] interface {
	Server(ctx context.Context, log *zerolog.Logger, dep Dep, conf Conf) error
}
