package memyandex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DenisquaP/ya-metrics/internal/agent/compress"
	"github.com/DenisquaP/ya-metrics/internal/models"
)

// Send sends metric to server
func Send(sender Sender, addr, name string) error {
	return sender.Send(addr, name)
}

// Send sends counter metrics to server
func (c Counter) Send(addr, name string) error {
	intC := int64(c)
	req := models.Metrics{
		ID:    name,
		MType: "counter",
		Delta: &intC,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Creating a new gzip writer
	buf, err := compress.GetGZip(body)
	if err != nil {
		return err
	}

	// Sending request with compressed data
	client := http.Client{Timeout: 20 * time.Second}
	reqw, err := http.NewRequest("POST", fmt.Sprintf(MetricsUpdateURL, addr), buf)
	if err != nil {
		return err
	}
	reqw.Header.Set("Content-Type", "application/json")
	reqw.Header.Set("Content-Encoding", "gzip")
	reqw.Header.Set("Accept-Encoding", "gzip")

	resp, err := client.Do(reqw)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not expected status code: %d", resp.StatusCode)
	}

	return nil
}

// Send sends gauge metrics to server
func (g Gauge) Send(addr, name string) error {
	floatG := float64(g)
	req := models.Metrics{
		ID:    name,
		MType: "gauge",
		Value: &floatG,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// Creating a new gzip writer
	buf, err := compress.GetGZip(body)
	if err != nil {
		return err
	}

	client := http.Client{}

	// Sending request with compressed data
	reqw, err := http.NewRequest("POST", fmt.Sprintf(MetricsUpdateURL, addr), buf)
	if err != nil {
		return err
	}
	reqw.Header.Set("Content-Type", "application/json")
	reqw.Header.Set("Content-Encoding", "gzip")
	reqw.Header.Set("Accept-Encoding", "gzip")

	resp, err := client.Do(reqw)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not expected status code: %d", resp.StatusCode)
	}

	return nil
}
