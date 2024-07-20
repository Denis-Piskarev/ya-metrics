package yametrics

import (
	"context"
	"fmt"
)

// Получение всех метрик
func (m *MemStorage) GetMetrics(ctx context.Context) (string, error) {
	res := ""

	for k, v := range m.Gauge {
		res += fmt.Sprintf("%v: %v\n", k, v)
	}
	for k, v := range m.Counter {
		res += fmt.Sprintf("%v: %v\n", k, v)
	}

	return res, nil
}

// Получение метрики типа Gauge
func (m *MemStorage) GetGauge(ctx context.Context, name string) (float64, error) {
	g, ok := m.Gauge[name]
	if !ok {
		return 0, fmt.Errorf("variable does not exist")
	}

	return g, nil
}

// Получение метрики типа Counter
func (m *MemStorage) GetCounter(ctx context.Context, name string) (int64, error) {
	c, ok := m.Counter[name]
	if !ok {
		return 0, fmt.Errorf("variable does not exists")
	}

	return c, nil
}
