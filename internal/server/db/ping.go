package db

import "context"

func (d Db) Ping(ctx context.Context) error {
	return d.Db.Ping(ctx)
}
