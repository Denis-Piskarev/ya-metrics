package usecase

import (
	"context"

	"github.com/DenisquaP/ya-metrics/internal/models"
)

func (m *Metric) WriteGauge(ctx context.Context, name string, val float64) (float64, error) {
	return m.m.WriteGauge(ctx, name, val)
}

func (m *Metric) WriteCounter(ctx context.Context, name string, val int64) (int64, error) {
	return m.m.WriteCounter(ctx, name, val)
}

func (m *Metric) WriteMetrics(ctx context.Context, metrics []*models.Metrics) error {
	return m.m.WriteMetrics(ctx, metrics)
}
