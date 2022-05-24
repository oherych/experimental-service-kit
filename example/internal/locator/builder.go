package locator

import (
	_ "embed"
	"github.com/oherych/experimental-service-kit/example/internal/migrations"
	"github.com/oherych/experimental-service-kit/example/internal/repository"
	"github.com/oherych/experimental-service-kit/kit/postgres"
)

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
