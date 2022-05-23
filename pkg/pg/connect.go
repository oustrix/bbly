package pg

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/ini.v1"
)

// TODO: add tests

func GetDBConnectionConfig() (*pgxpool.Config, error) {
	cfg, err := ini.Load("../../config/config.ini")
	if err != nil {
		return nil, errors.New("failed to read config file")
	}
	connStr := cfg.Section("").Key("DB_CON").String()

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, errors.New("failed to parse connection string to pgxpool.Config")
	}
	return poolConfig, nil
}

func DbPool(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, errors.New("failed to connect to DB")
	}

	return conn, nil
}
