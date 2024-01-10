package pgrepo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type RepoConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewRepoPG(cfg RepoConfig) (*Postgres, error) {
	pool, err := pgxpool.Connect(context.Background(), cfg.DSN)
	if err != nil {
		return nil, err
	}

	pool.Config().MaxConnIdleTime, err = time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	pool.Config().MaxConns = int32(cfg.MaxOpenConns)

	return &Postgres{pool}, nil
}
