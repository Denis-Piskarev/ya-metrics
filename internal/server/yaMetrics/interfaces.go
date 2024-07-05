package yametrics

import "context"

// Metrics interface
type Metric interface {
	MetricGetter
	MetricWriter
	MetricSaver
}

// Iterface for writing metrics
type MetricWriter interface {
	WriteGauge(ctx context.Context, name string, val float64) (float64, error)
	WriteCounter(ctx context.Context, name string, val int64) (int64, error)
}

// Interface for getting metrics
type MetricGetter interface {
	GetMetrics(ctx context.Context) (string, error)
	GetGauge(ctx context.Context, name string) (float64, error)
	GetCounter(ctx context.Context, name string) (int64, error)
}

// Interface for saving metrics
type MetricSaver interface {
	SaveMetricsToFile(wd string) error
	RestoreFromFile(wd string) error
}
