package kit

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/oherych/experimental-service-kit/kit/application"
	"github.com/oherych/experimental-service-kit/kit/cmd"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
)

type Runner[Conf dependencies.Config, Dep dependencies.Locator] struct {
	construct *application.Construct[Conf, Dep]
}

func Server[Conf dependencies.Config, Dep dependencies.Locator](name string, debBuilder dependencies.Builder[Conf, Dep]) *Runner[Conf, Dep] {
	return &Runner[Conf, Dep]{
		construct: &application.Construct[Conf, Dep]{
			Name:       name,
			DebBuilder: debBuilder,
			Log:        zerolog.New(os.Stdout).Output(zerolog.ConsoleWriter{Out: os.Stderr}),
		},
	}

}

func (a *Runner[Conf, Dep]) WithPorts(port application.Port[Dep]) *Runner[Conf, Dep] {
	a.construct.Ports = append(a.construct.Ports, port)

	return a
}

func (a *Runner[Conf, Dep]) Run() {
	a.restIsMandatory()

	ctx := a.contextWithInterrupt()

	if err := cmd.Root(*a.construct).ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (a *Runner[Conf, Dep]) restIsMandatory() {
	for _, port := range a.construct.Ports {
		if _, ok := port.(Rest[Dep]); ok {
			return
		}
	}

	a.construct.Ports = append(a.construct.Ports, Rest[Dep]{
		Builder: func(e *echo.Echo, dep Dep) error { return nil },
	})
}

func (Runner[Conf, Dep]) contextWithInterrupt() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c

		cancel()
	}()

	return ctx
}
