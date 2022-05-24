package migrations

import (
	"embed"
	"github.com/oherych/experimental-service-kit/kit/migrate"
)

//go:embed *.sql
var embedMigrations embed.FS

var Schema = migrate.Migrate{
	Embed:   embedMigrations,
	Dialect: "postgres",
}
