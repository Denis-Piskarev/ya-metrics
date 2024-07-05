package migrations

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInitTables, downInitTables)
}

var metrics = []string{
	"HeapAlloc",
	"MCacheInuse",
	"MSpanInuse",
	"Sys",
	"GCCPUFraction",
	"HeapIdle",
	"StackInuse",
	"Alloc",
	"Lookups",
	"Mallocs",
	"NextGC",
	"HeapInuse",
	"Frees",
	"HeapReleased",
	"LastGC",
	"NumForcedGC",
	"NumGC",
	"RandomValue",
	"HeapObjects",
	"MSpanSys",
	"OtherSys",
	"TotalAlloc",
	"MCacheSys",
	"HeapSys",
	"PauseTotalNs",
	"StackSys",
	"GCSys",
}

func upInitTables(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	log.Println("start upInitTables")

	query := `
		CREATE TABLE IF NOT EXISTS metrics (
		    id SERIAL,
		    type VARCHAR(64) NOT NULL,
		    name VARCHAR(128) NOT NULL,
		    counter BIGINT,
		    gauge DOUBLE PRECISION
		)
	`

	queryMetrics := `INSERT INTO metrics (type, name) VALUES ($1, $2);`

	// Creating metrics table
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return err
	}

	// Inserting metrics
	for _, gauge := range metrics {
		gauge = strings.ToLower(gauge)
		if _, err := tx.ExecContext(ctx, queryMetrics, "gauge", gauge); err != nil {
			return err
		}
	}

	if _, err := tx.ExecContext(ctx, queryMetrics, "counter", "pollcount"); err != nil {
		return err
	}

	return nil
}

func downInitTables(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	log.Println("start downInitTables")

	if _, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS metrics"); err != nil {
		return err
	}

	return nil
}
