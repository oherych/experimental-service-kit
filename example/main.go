package main

import (
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/example/internal/rest"
	"github.com/oherych/experimental-service-kit/kit"
)

func main() {
	kit.Server("DEMO", locator.Constructor).
		WithListeners(rest.New()).
		// WithListeners(grpc.Router).
		Run()
}
