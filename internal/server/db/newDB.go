package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	DB *pgxpool.Pool
}

func NewDB(ctx context.Context, address string) (db *DB, err error) {
	cfg, err := pgxpool.ParseConfig(address)
	if err != nil {
		return nil, err
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		if err := conn.Ping(ctx); err != nil {
			return err
		}

		return nil
	}

	cfg.MaxConns = 1

	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB: conn,
	}, nil
}
