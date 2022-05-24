package dependencies

import "context"

type Locator interface {
	HealthCheck(ctx context.Context) map[string]interface{}
	Close() error
}

type Builder[Conf Config, Dep Locator] func(cnf Conf) (Dep, error)
