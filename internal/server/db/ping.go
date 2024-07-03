package db

import "context"

func (d *DB) Ping(ctx context.Context) error {
	return d.Db.Ping(ctx)
}
