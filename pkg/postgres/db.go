package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MustNew(cfg Config) *sqlx.DB {
	return sqlx.MustConnect("postgres", cfg.ConnectionString())
}
