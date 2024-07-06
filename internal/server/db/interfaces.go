// Package to operate with database

package db

import (
	"context"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/db.go -package=mocks

// DBInterface interface of database
type DBInterface interface {
	Ping(context.Context) error
	Writer
	Reader
}

// Writer interface for writing metrics
type Writer interface {
	// WriteGauge writes gauge metric into db
	WriteGauge(context.Context, string, float64) (float64, error)
	// WriteCounter writes counter metric into db
	WriteCounter(context.Context, string, int64) (int64, error)
	// WriteMetrics writes multiple metrics into db
	WriteMetrics(context.Context, []*models.Metrics) error
}

// Reader interface for reading metrics
type Reader interface {
	// GetMetrics get all metrics from DB
	GetMetrics(context.Context) (string, error)
	// GetGauge gets gauge metric from DB by name
	GetGauge(context.Context, string) (float64, error)
	// GetCounter gets counter metric from DB by name
	GetCounter(context.Context, string) (int64, error)
}
