package locator

import (
	"context"
	"database/sql"
	"github.com/oherych/experimental-service-kit/example/internal/migrations"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
	"github.com/oherych/experimental-service-kit/kit"
	"github.com/oherych/experimental-service-kit/pkg/postgres"
)

type Users interface {
	All(ctx context.Context, pagination kit.Pagination) ([]repository.User, error)
	GetByID(ctx context.Context, id string) (*repository.User, error)
	Delete(ctx context.Context, id string) error
}

type Container struct {
	Config Config
	Users  Users

	postgresCon *sql.DB
}

func Constructor(conf Config) (*Container, error) {
	postgresCon, err := postgres.Connect(conf.Postgres)
	if err != nil {
		return nil, err
	}

	if err := migrations.Schema.Up(postgresCon); err != nil {
		return nil, err
	}

	users, err := repository.NewUsers(postgresCon)
	if err != nil {
		return nil, err
	}

	return &Container{
		Config: conf,
		Users:  users,

		// external
		postgresCon: postgresCon,
	}, nil
}

func (l Container) HealthCheck(ctx context.Context) map[string]error {
	return map[string]error{
		"postgres": l.postgresCon.PingContext(ctx),
	}
}

func (l Container) Ready(ctx context.Context) bool {
	return true
}

func (l Container) Close() error {
	return l.postgresCon.Close()
}
