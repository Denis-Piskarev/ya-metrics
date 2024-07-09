package yametrics

import (
	"context"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

func (m *MemStorage) WriteMetrics(ctx context.Context, metrics []*models.Metrics) error {
	for _, metric := range metrics {
		switch metric.MType {
		case "gauge":
			m.Gauge[metric.ID] = *metric.Value
		case "counter":
			met, ok := m.Counter[metric.ID]
			if !ok {
				m.Counter[metric.ID] = *metric.Delta
			} else {
				m.Counter[metric.ID] = met + *metric.Delta
			}
		}
	}

	return nil
}
