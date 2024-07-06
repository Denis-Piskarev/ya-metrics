package yametrics

import (
	"context"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

func (y *MemStorage) WriteMetrics(ctx context.Context, metrics []*models.Metrics) error {
	for _, metric := range metrics {
		switch metric.MType {
		case "gauge":
			y.Gauge[metric.ID] = *metric.Value
		case "counter":
			met, ok := y.Counter[metric.ID]
			if !ok {
				y.Counter[metric.ID] = *metric.Delta
			} else {
				y.Counter[metric.ID] = met + *metric.Delta
			}
		}
	}

	return nil
}
