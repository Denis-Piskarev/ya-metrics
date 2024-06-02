package yametrics

import (
	"fmt"
	"strconv"
)

type Metric interface {
	WriteMetric(name, typeMet, val string) error
	GetMetrics() string
	GetMetric(typeMet, name string) (string, error)
}

// MemStorage struct
type MemStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}
}

func (m *MemStorage) WriteMetric(name, typeMet, val string) error {
	if typeMet == "gauge" {
		g, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return fmt.Errorf("unsupported type of gauge: %s", err.Error())
		}
		m.Gauge[name] = g
		return nil
	} else if typeMet == "counter" {
		c, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("unsupported type of gauge: %s", err.Error())
		}
		m.Counter[name] = int64(c)
		return nil
	}

	return fmt.Errorf("unsupported type of metric: %s", typeMet)
}

func (m *MemStorage) GetMetrics() string {
	res := ""

	for k, v := range m.Gauge {
		res += fmt.Sprintf("%v: %v\n", k, v)
	}
	for k, v := range m.Counter {
		res += fmt.Sprintf("%v: %v\n", k, v)
	}

	return res
}

func (m *MemStorage) GetMetric(typeMet, name string) (string, error) {
	if typeMet == "gouge" {
		g, ok := m.Gauge[name]
		if !ok {
			return "", fmt.Errorf("variable does not exists")
		}
		return fmt.Sprintf("%v", g), nil
	} else if typeMet == "counter" {
		c, ok := m.Counter[name]
		if !ok {
			return "", fmt.Errorf("variable does not exists")
		}
		return fmt.Sprintf("%v", c), nil
	}

	return "", fmt.Errorf("unsupported type of metric")
}
