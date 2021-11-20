package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	DBUser   string
	DBName   string
	Password string
	SSLMode  string
}

func NewPostgres(cfg *Config) (*sqlx.DB, error) {
	db_url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBUser, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", db_url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
