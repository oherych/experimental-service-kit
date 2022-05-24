package application

import (
	"context"
	"github.com/mcuadros/go-defaults"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

type Construct[Conf dependencies.Config, Dep dependencies.Locator] struct {
	Name       string
	DebBuilder dependencies.Builder[Conf, Dep]
	Ports      []Port[Dep]
	Log        zerolog.Logger
}

func (Construct[Conf, Dep]) Configuration() (Conf, error) {
	var cnf Conf

	defaults.SetDefaults(&cnf)

	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return cnf, err
		}
	}

	err := v.Unmarshal(&cnf)
	if err != nil {
		return cnf, err
	}

	return cnf, nil
}

func (c Construct[Conf, Dep]) Dependencies(cnf Conf) (Dep, error) {
	return c.DebBuilder(cnf)
}

func (c Construct[Conf, Dep]) RunServe(ctx context.Context) (err error) {
	c.Log.Info().Msg("[SYS] Read configuration")

	cnf, err := c.Configuration()
	if err != nil {
		return err
	}

	c.Log.Info().Msg("[SYS] Build dependencies")
	dep, err := c.Dependencies(cnf)
	if err != nil {
		c.Log.Error().Err(err).Send()

		return err
	}

	defer func() {
		c.Log.Info().Msg("[SYS] Destroy dependencies")

		if err := dep.Close(); err != nil {
			c.Log.Error().Err(err).Send()
		}
	}()

	g, ctx := errgroup.WithContext(ctx)

	bc := cnf.GetBaseConfig()

	for _, port := range c.Ports {
		fn := port.Server

		g.Go(func() error { return fn(ctx, c.Log, dep, bc) })
	}

	return g.Wait()
}
