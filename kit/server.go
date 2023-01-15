package kit

import (
	"context"
	"fmt"
	"github.com/oherych/experimental-service-kit/kit/dependencies"
	"github.com/oherych/experimental-service-kit/kit/server"
	"net/http"
	"os"
	"os/signal"
)

type ServerRunner[Conf dependencies.Config, Dep dependencies.Locator] struct {
	server *server.Config[Conf, Dep]
}

func Server[Conf dependencies.Config, Dep dependencies.Locator](name string, debBuilder dependencies.Builder[Conf, Dep]) *ServerRunner[Conf, Dep] {
	return &ServerRunner[Conf, Dep]{
		server: &server.Config[Conf, Dep]{
			Name:       name,
			DebBuilder: debBuilder,
		},
	}

}

func (a *ServerRunner[Conf, Dep]) WithListeners(port server.Listener[Conf, Dep]) *ServerRunner[Conf, Dep] {
	a.server.Listeners = append(a.server.Listeners, port)

	return a
}

func (a *ServerRunner[Conf, Dep]) Run() {
	ctx := a.contextWithInterrupt()

	a.server.Args = os.Args

	if err := a.server.New(ctx); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		os.Exit(1)
	}
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
