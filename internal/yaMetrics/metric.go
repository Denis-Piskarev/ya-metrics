package yametrics

import (
	"fmt"
	"strconv"
)

// Имя метрики: метрика
var metrics = make(map[string]*MemStorage)

type MemStorage struct {
	Gauge   float64
	Counter int64
}

func WriteMetric(name, typeMetric, val string) error {
	met, ok := metrics[name]
	if !ok {
		metrics[name] = &MemStorage{}
		met = metrics[name]
	}

	switch typeMetric {
	case "counter":
		counter, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("unsupported type of counter")
		}

		met.Counter += int64(counter)
	case "gauge":
		gauge, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return fmt.Errorf("unsupported type of gauge")
		}

		met.Gauge = gauge
	default:
		return fmt.Errorf("unsupported metric type: %s", typeMetric)
	}

	return nil
}
