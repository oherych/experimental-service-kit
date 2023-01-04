package echo_listener

import "github.com/labstack/echo/v4"

type Handler func(c Context) error

func Wrap(in Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return in(Context{o: c})
	}
}
