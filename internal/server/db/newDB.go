package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	Db *pgxpool.Pool
}

func NewDB(ctx context.Context, address string) (db Db, err error) {
	cfg, err := pgxpool.ParseConfig(address)
	if err != nil {
		return Db{}, err
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
		return Db{}, err
	}

	return Db{
		Db: conn,
	}, nil
}
