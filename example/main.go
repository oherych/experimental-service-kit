package main

import (
	"github.com/oherych/experimental-service-kit/example/internal/grpc"
	"github.com/oherych/experimental-service-kit/example/internal/locator"
	"github.com/oherych/experimental-service-kit/example/internal/rest"
	"github.com/oherych/experimental-service-kit/kit"
)

func main() {
	kit.Server("DEMO", locator.Builder).
		WithPorts(rest.Router).
		WithPorts(grpc.Router).
		Run()
}
