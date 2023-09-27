package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	DBName   string `koanf:"dbname"`
}

type DB struct {
	db *sql.DB
}

func New(config Config) *DB {
	param := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", config.Host, config.Port, config.User, config.Password)
	db, err := sql.Open("postgres", param)
	if err != nil {
		_ = fmt.Errorf("can't connect to postgres with err: %v", err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DB{db: db}
}
