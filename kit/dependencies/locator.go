package dependencies

import "context"

type Locator interface {
	HealthCheck(ctx context.Context) map[string]error
	Ready(ctx context.Context) bool
	Close() error
}

type Builder[Conf Config, Dep Locator] func(cnf Conf) (Dep, error)
