package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect(config Config) (*sql.DB, error) {
	bd, err := sql.Open("postgres", config.DSN)
	return bd, err
}
