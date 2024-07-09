package db

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

func (d *DB) WriteMetrics(ctx context.Context, metrics []*models.Metrics) error {
	// Beginning transaction
	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	query := `INSERT INTO metrics (type, name, counter, gauge) VALUES ($1, $2, $3, $4)`

	for _, metric := range metrics {
		if _, err := tx.Exec(ctx, query, metric.MType, metric.ID, metric.Delta, metric.Value); err != nil {
			d.Logger.Errorw("write metrics error", "error", err)
			return err
		}
	}

	return nil
}
