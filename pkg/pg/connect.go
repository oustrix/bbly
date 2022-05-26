package pg

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectToDB() error {
	cfg, err := getDBConnectionConfig()
	if err != nil {
		return err
	}

	cfg.MaxConns = 10

	err = createDbPool(cfg)
	if err != nil {
		return err
	}

	return nil
}

func getDBConnectionConfig() (*pgxpool.Config, error) {
	connStr := os.Getenv("CONN_STR")

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, errors.New("failed to parse connection string to pgxpool.Config")
	}
	return poolConfig, nil
}

func createDbPool(poolConfig *pgxpool.Config) error {
	var err error
	DB, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return errors.New("failed to connect to DB")
	}
	err = checkDBConnection()
	if err != nil {
		return errors.New("failed to ping server")
	}
	return nil
}

func checkDBConnection() error {
	return DB.Ping(context.Background())
}
