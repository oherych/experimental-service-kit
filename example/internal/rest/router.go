package rest

import (
	_ "embed"
	"github.com/labstack/echo/v4"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/kit"
)

//go:embed data/swagger.yaml
var swaggerContent []byte

var Router = kit.Rest[locator.Implementation]{
	Swagger: swaggerContent,
	Builder: func(e *echo.Echo, dep locator.Implementation) error {
		e.GET("/user/", UsersController{d: dep}.List)
		e.GET("/user/:id/", UsersController{d: dep}.Get)

		return nil
	},
}
