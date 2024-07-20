package yametrics

import "context"

// Запись метрики типа Gauge
func (m *MemStorage) WriteGauge(ctx context.Context, name string, val float64) (float64, error) {
	m.Gauge[name] = val
	return m.Gauge[name], nil
}

// Запись метрики типа Counter
func (m *MemStorage) WriteCounter(ctx context.Context, name string, val int64) (int64, error) {
	m.Counter[name] += val
	return m.Counter[name], nil
}
