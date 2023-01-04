package locator

import (
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	echo_listener "github.com/oherych/experimental-service-kit/pkg/echo-listener"
	"github.com/oherych/experimental-service-kit/pkg/postgres"
)

type Config struct {
	dependencies.Base

	// Listeners
	Echo echo_listener.Config `envconfig:"ROUTER"`

	// Dependencies
	Postgres postgres.Config

	// Custom configuration
	MyArg string `env:"MYARG"`
}

func (c Config) Validate() error {
	return nil
}

