package rest

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/kit"
	"github.com/oherych/experimental-service-kit/kit/responses"
)

//go:embed data/swagger.yaml
var swaggerContent []byte

var Router = kit.HttpEcho[locator.Implementation]{
	Swagger:      swaggerContent,
	ErrorHandler: errorHandling,
	Builder: func(e *echo.Echo, dep locator.Implementation) error {
		e.GET("/user/", UsersController{d: dep}.List)
		e.GET("/user/:id/", UsersController{d: dep}.Get)

		return nil
	},
}

func errorHandling(err error, c echo.Context) error {
	switch e := err.(type) {
	case responses.NotFound:
		return c.JSON(http.StatusNotFound, map[string]any{"reason": e.Reason})
	}

	return err
}
