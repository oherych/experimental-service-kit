package rest

import (
	_ "embed"
	"github.com/oherych/experimental-service-kit/pkg/echo-listener"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
)

//go:embed data/swagger.yaml
var swaggerContent []byte

var Router = echo_listener.HttpEcho[locator.Config, *locator.Container]{
	Swagger:      swaggerContent,
	ErrorHandler: errorHandling,
	Init: func(conf locator.Config) echo_listener.Config {
		return conf.Echo
	},
	Builder: func(e *echo.Echo, dep *locator.Container) error {
		e.GET("/user/", UsersController{d: dep}.List)
		e.GET("/user/:id/", UsersController{d: dep}.Get)
		e.DELETE("/user/:id/", UsersController{d: dep}.Delete)

		return nil
	},
}

func errorHandling(err error, c echo.Context) error {
	switch e := err.(type) {
	case echo_listener.NotFound:
		return c.JSON(http.StatusNotFound, map[string]any{"reason": e.Reason})
	}

	return err
}
