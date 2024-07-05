package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type DB struct {
	DB     *pgx.Conn
	Logger *zap.SugaredLogger
}

func NewDB(ctx context.Context, logger *zap.SugaredLogger, address string) (db *DB, err error) {
	conn, err := pgx.Connect(ctx, address)
	if err != nil {
		return nil, err
	}

	return &DB{
		Logger: logger,
		DB:     conn,
	}, nil
}
