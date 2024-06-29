package yametrics

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (m *MemStorage) SaveToFile(wd string) error {
	file, err := os.OpenFile(filepath.Join(wd, m.FilePath), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	metrics, err := getMetricsJson(m)
	if err != nil {
		return err
	}

	_, err = file.Write(metrics)
	if err != nil {
		return err
	}

	return nil
}

func (m *MemStorage) Restore() error {
	return nil
}

func getMetricsJson(m *MemStorage) ([]byte, error) {
	type metric struct {
		Name  string      `json:"name"`
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	}

	metrics := make([]metric, 0, len(m.Gauge)+len(m.Counter))

	for k, v := range m.Gauge {
		var m metric

		m.Name = k
		m.Type = "gauge"
		m.Value = v

		metrics = append(metrics, m)
	}

	for k, v := range m.Counter {
		var m metric

		m.Name = k
		m.Type = "counter"
		m.Value = v
	}

	return json.Marshal(metrics)
}
