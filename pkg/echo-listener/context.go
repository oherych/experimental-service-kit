package echo_listener

import (
	"context"
	"github.com/oherych/experimental-service-kit/kit"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Context struct {
	o echo.Context
}

func (c Context) Origin() echo.Context {
	return c.o
}

func (c Context) Bind(target any) error {
	return c.Bind(target)
}

func (c Context) Context() context.Context {
	return c.o.Request().Context()
}

// permissions

func (c Context) ShouldCan(permissions ...string) error {
	return nil
}

// params

func (c Context) ParamString(name string) string {
	return c.o.Param(name)
}

func (c Context) ParamInt(name string) (int, error) {
	str := c.o.Param(name)

	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, kit.WrongParameter{Name: name, Internal: err}
	}

	return val, nil
}

func (c Context) Pagination(max int) (kit.Pagination, error) {
	// TODO: implement me

	return kit.Pagination{}, nil
}

// responses

func (c Context) List(p kit.Pagination, i interface{}) error {
	return c.o.JSON(http.StatusOK, map[string]any{
		"from":  p.From,
		"limit": p.Limit,
		"list":  i,
	})
}

func (c Context) StatusOK(i interface{}) error {
	return c.o.JSON(http.StatusOK, i)
}

func (c Context) Deleted() error {
	return c.o.NoContent(http.StatusNoContent)
}
