package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// GetMetrics get all metrics from DB
func (d *DB) GetMetrics(ctx context.Context) (string, error) {
	rows, err := d.DB.Query(ctx, `SELECT type, name, gauge, counter FROM metrics`)
	if err != nil {
		d.Logger.Errorw("error getting metrics", "error", err)
		return "", err
	}
	defer rows.Close()

	metrics := ""

	for rows.Next() {
		var typeMet string
		var name string
		var gauge sql.NullFloat64
		var counter sql.NullInt64

		if err := rows.Scan(&typeMet, &name, &gauge, &counter); err != nil {
			d.Logger.Errorw("scan error", "error", err)
			return "", err
		}

		if gauge.Valid {
			metrics += fmt.Sprintf("%s: %v\n", name, gauge.Float64)
		} else if counter.Valid {
			metrics += fmt.Sprintf("%s: %v\n", name, counter.Int64)
		} else {
			return "", fmt.Errorf("not found metrics")
		}
	}

	return metrics, nil
}

// GetCounter gets counter metric from DB by name
func (d *DB) GetCounter(ctx context.Context, name string) (int64, error) {
	var counter sql.NullInt64

	if err := d.DB.QueryRow(ctx, "SELECT counter FROM metrics WHERE name = $1 AND type = 'counter'", name).Scan(&counter); err != nil {
		return 0, err
	}

	if !counter.Valid {
		return 0, fmt.Errorf("unexpected type of metric")
	}

	return counter.Int64, nil
}

// GetGauge gets gauge metric from DB by name
func (d *DB) GetGauge(ctx context.Context, name string) (float64, error) {
	var gauge sql.NullFloat64

	if err := d.DB.QueryRow(ctx, "SELECT gauge FROM metrics WHERE name = $1 AND type = 'gauge'", name).Scan(&gauge); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("no metrics found")
			return 0, nil
		}
		return 0, err
	}

	if !gauge.Valid {
		return 0, fmt.Errorf("unexpected type of metric")
	}

	return gauge.Float64, nil
}
