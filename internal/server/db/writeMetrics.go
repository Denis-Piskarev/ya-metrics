package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/DenisquaP/ya-metrics/internal/models"
)

// WriteMetrics writes multiple metrics into db
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
		// If metric type is counter we need to get old value
		if metric.MType == "counter" {
			var oldCounter int64
			if err := tx.QueryRow(ctx, `SELECT counter FROM metrics WHERE name=$1 AND type = 'counter'`, metric.ID).Scan(&oldCounter); err != nil {
				// If no rows we insert new value
				if errors.Is(err, pgx.ErrNoRows) {
					if _, err := tx.Exec(ctx, query, metric.MType, metric.ID, metric.Delta, metric.Value); err != nil {
						d.Logger.Errorw("write new counter error", "error", err)
						return err
					}

					continue
				}

				d.Logger.Errorw("error get old counter", "error", err)
				return err
			}

			// Add old value
			*metric.Delta += oldCounter

			// Write new value
			if _, err := tx.Exec(ctx, `UPDATE metrics SET counter = $1 WHERE type = 'counter' AND name = $2`, metric.Delta, metric.ID); err != nil {
				d.Logger.Errorw("write counter error", "error", err)
				return err
			}
		} else if metric.MType == "gauge" {
			// Updating gauge
			ra, err := tx.Exec(ctx, `UPDATE metrics SET gauge = $1 WHERE type = 'gauge' AND name = $2`, metric.Value, metric.ID)
			if err != nil {
				d.Logger.Errorw("write gauge error", "error", err)
				return err
			}

			// If no rows updated we insert new value
			if ra.RowsAffected() == 0 {
				if _, err := tx.Exec(ctx, query, metric.MType, metric.ID, metric.Delta, metric.Value); err != nil {
					d.Logger.Errorw("write new gauge error", "error", err)
					return err
				}
			}
		}
	}

	return nil
}
