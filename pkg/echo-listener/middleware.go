package echo_listener

import (
	"github.com/labstack/echo/v4"
	"github.com/oherych/experimental-service-kit/kit/logs"
	"github.com/rs/zerolog"
)

func traceIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetRequest(c.Request().WithContext(logs.WithContext(
				c.Request().Context(),
				c.Request().Header.Get(logs.TraceHeaderKey),
			)))

			return next(c)
		}
	}
}

func loggerMiddleware(log *zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			var rec *zerolog.Event

			if err := next(c); err != nil {
				c.Error(err)

				rec = log.Error().Err(err)
			} else {
				rec = log.Info()
			}

			rec.
				Str("_p", req.Method+" "+req.URL.String()).
				Int("_s", res.Status).
				Str("_r", logs.TraceID(req.Context())).
				Send()

			return nil
		}
	}
}
