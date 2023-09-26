package postgres

import (
	"database/sql"
	"log/slog"
)

type Config struct {
}

type PostgresDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) *PostgresDB {
	param := ""
	db, err := sql.Open("postgres", param)
	if err != nil {
		slog.Error("can't connect to database: ", err)
	}
	return &PostgresDB{
		config: config,
		db:     db,
	}
}

func (d *PostgresDB) Conn() *sql.DB {
	return d.db
}
