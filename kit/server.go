package kit

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/oherych/experimental-service-kit/kit/application"
	"github.com/oherych/experimental-service-kit/kit/cmd"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/oherych/experimental-service-kit/kit/logs"
	"os"
	"os/signal"
)

type ServerRunner[Conf dependencies.Config, Dep dependencies.Locator] struct {
	construct *application.Construct[Conf, Dep]
}

func Server[Conf dependencies.Config, Dep dependencies.Locator](name string, debBuilder dependencies.Builder[Conf, Dep]) *ServerRunner[Conf, Dep] {
	return &ServerRunner[Conf, Dep]{
		construct: &application.Construct[Conf, Dep]{
			Name:       name,
			DebBuilder: debBuilder,
			Log:        logs.New(os.Stdout),
		},
	}

}

func (a *ServerRunner[Conf, Dep]) WithPorts(port application.Port[Dep]) *ServerRunner[Conf, Dep] {
	a.construct.Ports = append(a.construct.Ports, port)

	return a
}

func (a *ServerRunner[Conf, Dep]) Run() {
	//a.restIsMandatory()

	ctx := a.contextWithInterrupt()

	if err := cmd.Root(*a.construct).ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (a *ServerRunner[Conf, Dep]) restIsMandatory() {
	for _, port := range a.construct.Ports {
		if _, ok := port.(HttpEcho[Dep]); ok {
			return
		}
	}

	a.construct.Ports = append(a.construct.Ports, HttpEcho[Dep]{
		Builder: func(e *echo.Echo, dep Dep) error { return nil },
	})
}

func (ServerRunner[Conf, Dep]) contextWithInterrupt() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c

		cancel()
	}()

	return ctx
}
