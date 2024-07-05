package db

import "context"

func (d *DB) Ping(ctx context.Context) error {
	if err := d.DB.Ping(ctx); err != nil {
		d.Logger.Errorw("ping error", "error", err)
		return err
	}

	return nil
}
