package locator

import (
	"context"
	"database/sql"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
)

type Implementation struct {
	Config Config
	Users  Users

	postgresCon *sql.DB
}

type Users interface {
	All(ctx context.Context) ([]repository.User, error)
	GetByID(ctx context.Context, id string) (*repository.User, error)
}

func (l Implementation) HealthCheck(ctx context.Context) map[string]any {
	return map[string]any{
		"postgres": l.postgresCon.PingContext(ctx) == nil,
	}
}

func (l Implementation) Close() error {
	return l.postgresCon.Close()
}
