package locator

import (
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/oherych/experimental-service-kit/kit/postgres"
)

type Config struct {
	dependencies.BaseConfig

	Postgres postgres.Config `mapstructure:",squash"`

	MyArg string `env:"MYARG"`
}

func (c Config) GetBaseConfig() dependencies.BaseConfig {
	return c.BaseConfig
}
