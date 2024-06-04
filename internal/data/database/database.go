package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const driverName = "postgres"

type DB struct {
	*sqlx.DB
}

type Config struct {
	Host     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func NewDB(cfg Config) (*DB, error) {
	dataSourceName := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode,
	)
	sqlxDB, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{sqlxDB}, nil
}
