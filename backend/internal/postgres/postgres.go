package postgres

import (
	"context"
	"fmt"

	"github.com/404th/Ink/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const ()

type Database struct {
	Pool *pgxpool.Pool
}

func NewPostgres(cfg *config.Config) (db *Database, err error) {
	db = new(Database)

	ctx := context.Background()
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	db.Pool = pool

	return db, nil
}
