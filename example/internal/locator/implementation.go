package locator

import (
	"context"
	"database/sql"

	"github.com/oherych/experimental-service-kit/example/internal/migrations"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
	"github.com/oherych/experimental-service-kit/kit/features"
	"github.com/oherych/experimental-service-kit/kit/postgres"
)

type Users interface {
	All(ctx context.Context) ([]repository.User, error)
	GetByID(ctx context.Context, id string) (*repository.User, error)
}

type Implementation struct {
	Config Config
	Users  Users

	FF features.Interface

	postgresCon *sql.DB
}

func Builder(conf Config) (Implementation, error) {
	postgresCon, err := postgres.Connect(conf.Postgres)
	if err != nil {
		return Implementation{}, err
	}

	if err := migrations.Schema.Up(postgresCon); err != nil {
		return Implementation{}, err
	}

	users, err := repository.NewUsers(postgresCon)
	if err != nil {
		return Implementation{}, err
	}

	return Implementation{
		Config: conf,
		Users:  users,

		// external
		postgresCon: postgresCon,
	}, nil
}

func (l Implementation) HealthCheck(ctx context.Context) map[string]error {
	return map[string]error{
		"postgres": l.postgresCon.PingContext(ctx),
	}
}

func (l Implementation) Ready(ctx context.Context) bool {
	return true
}

func (l Implementation) Close() error {
	return l.postgresCon.Close()
}
