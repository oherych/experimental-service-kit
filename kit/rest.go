package kit

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/oherych/experimental-service-kit/kit/logs"
	"github.com/rs/zerolog"
	"net/http"
)

type Rest[Dep dependencies.Locator] struct {
	Swagger []byte
	Builder func(e *echo.Echo, dep Dep) error
}

func (m Rest[Dep]) Server(ctx context.Context, log zerolog.Logger, dep Dep, bc dependencies.BaseConfig) error {
	log.Info().Str("port", bc.HttpPort).Msg("[SYS] Starting REST server on port")

	r, err := m.create(dep, log)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()

		log.Info().Msg("[SYS] Stop REST server")

		_ = r.Close()
	}()

	err = r.Start(bc.HttpPort)
	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (m Rest[Dep]) create(dep Dep, log zerolog.Logger) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.AddTrailingSlash())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			request = request.WithContext(logs.ToContext(request.Context(), &log))
			c.SetRequest(request)

			return next(c)
		}
	})
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	if err := m.Builder(e, dep); err != nil {
		return nil, err
	}

	e.GET("_health/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dep.HealthCheck(c.Request().Context()))
	})

	if len(m.Swagger) > 0 {
		e.GET("_swagger.yaml/", func(c echo.Context) error {
			return c.Blob(http.StatusOK, "text/yaml", m.Swagger)
		})
	}

	return e, nil
}
