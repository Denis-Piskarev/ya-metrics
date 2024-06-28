package yametrics

import (
	"encoding/json"
	"os"
)

func (m *MemStorage) SaveToFile() error {
	file, err := os.OpenFile(m.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
