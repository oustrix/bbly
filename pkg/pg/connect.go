package pg

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/ini.v1"
)

func DbConnect() (*pgxpool.Pool, error) {
	cfg, err := ini.Load("../../config/config.ini")
	if err != nil {
		return nil, errors.New("failed to read config file")
	}

	connStr := cfg.Section("").Key("DB_CON").String()
	conn, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, errors.New("failed to connect to DB")
	}

	return conn, nil
}
