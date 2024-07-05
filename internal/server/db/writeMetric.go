package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
)

// WriteGaoge writes gauge metric into db
func (d *DB) WriteGauge(ctx context.Context, name string, value float64) (float64, error) {
	name = strings.ToLower(name)

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	if _, err := tx.Exec(ctx, `UPDATE metrics SET gauge = $1 WHERE type = 'gauge' AND name = $2`, value, name); err != nil {
		d.Logger.Errorw("write gauge error", "error", err)
		return 0, err
	}

	var newVal sql.NullFloat64
	if err := tx.QueryRow(ctx, "SELECT gauge FROM metrics WHERE name = $1 AND type = 'gauge'", name).Scan(&newVal); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("now rows")
			return 0, nil
		}
		return 0, err
	}

	if !newVal.Valid {
		return 0, fmt.Errorf("gauge metric does not exist")
	}

	return newVal.Float64, nil
}

// WriteCounter writes counter metric into db
func (d *DB) WriteCounter(ctx context.Context, name string, value int64) (int64, error) {
	name = strings.ToLower(name)

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	var oldVal sql.NullInt64
	if err := tx.QueryRow(ctx, "SELECT counter FROM metrics WHERE name = $1 AND type = 'counter'", name).Scan(&oldVal); err != nil {
		return 0, err
	}

	if !oldVal.Valid {
		oldVal.Int64 = 0
	}

	value += oldVal.Int64

	if _, err := tx.Exec(ctx, `UPDATE metrics SET counter = $1 WHERE type = 'counter' AND name = $2`, value, name); err != nil {
		d.Logger.Errorw("write counter error", "error", err)
		return 0, err
	}

	return value, nil
}
