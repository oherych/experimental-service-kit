package echo_listener

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/oherych/experimental-service-kit/kit/logs"
	"github.com/rs/zerolog"
	"net/http"
)

type HttpEcho[Conf dependencies.Config, Dep dependencies.Locator] struct {
	Swagger      []byte
	Init         func(conf Conf) Config
	Builder      func(e *echo.Echo, dep Dep) error
	ErrorHandler func(error, echo.Context) error
}

func (m HttpEcho[Conf, Dep]) Server(ctx context.Context, log zerolog.Logger, dep Dep, global Conf) error {
	config := m.Init(global)

	log.Info().Str("port", config.HTTPPort).Msg("[SYS] Starting REST cmd on port")

	r, err := m.create(dep, log)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()

		log.Info().Msg("[SYS] Stop REST cmd")

		_ = r.Close()
	}()

	err = r.Start(config.HTTPPort)
	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (m HttpEcho[_, Dep]) create(dep Dep, log zerolog.Logger) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	if m.ErrorHandler != nil {
		e.HTTPErrorHandler = m.buildErrorHandler()
	}

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
		result := dep.HealthCheck(c.Request().Context())
		for _, err := range result {
			if err != nil {
				return c.JSON(http.StatusGatewayTimeout, result)
			}
		}

		return c.JSON(http.StatusOK, result)
	})

	if len(m.Swagger) > 0 {
		e.GET("_swagger.yaml/", func(c echo.Context) error {
			return c.Blob(http.StatusOK, "text/yaml", m.Swagger)
		})
	}

	return e, nil
}

func (m HttpEcho[_, _]) buildErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		err = m.ErrorHandler(err, c)
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}
