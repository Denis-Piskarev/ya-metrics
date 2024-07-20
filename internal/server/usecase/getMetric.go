package usecase

import "context"

// GetMetrics gets all metrics. Returns string with all metrics
func (m *Metric) GetMetrics(ctx context.Context) (string, error) {
	return m.m.GetMetrics(ctx)
}

// GetGauge gets gauge metric
func (m *Metric) GetGauge(ctx context.Context, name string) (float64, error) {
	return m.m.GetGauge(ctx, name)
}

// GetCounter gets counter metric
func (m *Metric) GetCounter(ctx context.Context, name string) (int64, error) {
	return m.m.GetCounter(ctx, name)
}
