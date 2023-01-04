package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/rs/zerolog"
)

type Monitoring[Conf dependencies.Config, Dep dependencies.Locator] struct{}

func (m Monitoring[Conf, Dep]) Server(ctx context.Context, log zerolog.Logger, dep Dep, conf Conf) error {
	bc := conf.GetBaseConfig()

	r := echo.New()
	r.HideBanner = true
	r.HidePort = true

	r.GET("", m.versionHandler(bc))
	r.GET("health", m.healthHandler(dep))
	r.GET("ready", m.readyHandler(dep))

	log.Info().Str("port", bc.InternalPort).Msg("[SYS] Starting Monitoring")

	go func() {
		<-ctx.Done()

		log.Info().Msg("[SYS] Stop Monitoring")

		_ = r.Close()
	}()

	return r.Start(bc.InternalPort)
}

func (m Monitoring[Conf, Dep]) versionHandler(bc dependencies.Base) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, map[string]any{
			"name": bc.AppName,
		})
	}
}

func (m Monitoring[Conf, Dep]) healthHandler(dep Dep) echo.HandlerFunc {
	return func(c echo.Context) error {
		result := dep.HealthCheck(c.Request().Context())
		for _, err := range result {
			if err != nil {
				return c.JSON(http.StatusGatewayTimeout, result)
			}
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (m Monitoring[Conf, Dep]) readyHandler(dep Dep) echo.HandlerFunc {
	return func(c echo.Context) error {
		ready := dep.Ready(c.Request().Context())
		if ready {
			return c.NoContent(http.StatusOK)
		}

		return c.NoContent(http.StatusGone)
	}
}
