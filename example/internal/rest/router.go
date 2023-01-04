package rest

import (
	_ "embed"

	"github.com/labstack/echo/v4"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/kit/server"
	listener "github.com/oherych/experimental-service-kit/pkg/echo-listener"
)

//go:embed data/swagger.yaml
var swaggerContent []byte

func New() server.Listener[locator.Config, *locator.Container] {
	return listener.HttpEcho[locator.Config, *locator.Container]{
		Swagger: swaggerContent,
		InitConfig: func(conf locator.Config) listener.Config {
			return conf.Echo
		},
		Builder: func(e *echo.Echo, dep *locator.Container) error {
			e.GET("/user/", listener.Wrap(UsersController{d: dep}.List))
			e.GET("/user/:id/", listener.Wrap(UsersController{d: dep}.Get))
			e.DELETE("/user/:id/", listener.Wrap(UsersController{d: dep}.Delete))

			return nil
		},
	}
}
