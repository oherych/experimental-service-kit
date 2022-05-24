package migrate

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

type Migrate struct {
	Embed   embed.FS
	Dialect string
}

func (m Migrate) Up(db *sql.DB) error {
	goose.SetBaseFS(m.Embed)

	if err := goose.SetDialect(m.Dialect); err != nil {
		return err
	}

	return goose.Up(db, ".")
}
