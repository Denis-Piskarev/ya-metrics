package usecase

import (
	"context"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

// Metrics interface
type MetricInterface interface {
	MetricGetter
	MetricWriter
	Ping(ctx context.Context) error
}

// Iterface for writing metrics
type MetricWriter interface {
	// WriteGauge writes gauge metric
	WriteGauge(ctx context.Context, name string, val float64) (float64, error)
	// WriteCounter writes counter metric
	WriteCounter(ctx context.Context, name string, val int64) (int64, error)
	// WriteMetrics writes multiple metrics
	WriteMetrics(context.Context, []*models.Metrics) error
}

// Interface for getting metrics
type MetricGetter interface {
	// GetMetrics get all metrics
	GetMetrics(ctx context.Context) (string, error)
	// GetGauge gets gauge metric
	GetGauge(ctx context.Context, name string) (float64, error)
	// GetCounter gets counter metric
	GetCounter(ctx context.Context, name string) (int64, error)
}
