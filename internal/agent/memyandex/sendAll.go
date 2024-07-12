package memyandex

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

func getPointerFloat(v float64) *float64 {
	return &v
}

func getPointerInt(v int64) *int64 {
	return &v
}

func (m *MemStatsYaSt) SendAllMetricsToServer(ctx context.Context, addr string) error {
	req2 := []models.Metrics{
		{
			ID:    "Alloc",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.Alloc)),
		},
		{
			ID:    "BuckHashSys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.BuckHashSys)),
		},
		{
			ID:    "Frees",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.Frees)),
		},
		{
			ID:    "GCCPUFraction",
			MType: "gauge",
			Value: getPointerFloat(m.RuntimeMem.GCCPUFraction),
		},
		{
			ID:    "GCSys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.GCSys)),
		},
		{
			ID:    "HeapAlloc",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.HeapAlloc)),
		},
		{
			ID:    "HeapIdle",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.HeapIdle)),
		},
		{
			ID:    "HeapObjects",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.HeapObjects)),
		},
		{
			ID:    "HeapReleased",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.HeapReleased)),
		},
		{
			ID:    "HeapSys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.HeapSys)),
		},
		{
			ID:    "LastGC",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.LastGC)),
		},
		{
			ID:    "Lookups",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.Lookups)),
		},
		{
			ID:    "MCacheInuse",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.MCacheInuse)),
		},
		{
			ID:    "MCacheSys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.MCacheSys)),
		},
		{
			ID:    "MSpanInuse",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.MSpanInuse)),
		},
		{
			ID:    "MSpanSys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.MSpanSys)),
		},
		{
			ID:    "Mallocs",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.Mallocs)),
		},
		{
			ID:    "NextGC",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.NextGC)),
		},
		{
			ID:    "NumForcedGC",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.NumForcedGC)),
		},
		{
			ID:    "NumGC",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.NumGC)),
		},
		{
			ID:    "OtherSys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.OtherSys)),
		},
		{
			ID:    "PauseTotalNs",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.PauseTotalNs)),
		},
		{
			ID:    "StackInuse",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.StackInuse)),
		},
		{
			ID:    "Sys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.Sys)),
		},
		{
			ID:    "StackSys",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.StackSys)),
		},
		{
			ID:    "TotalAlloc",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.TotalAlloc)),
		},
		{
			ID:    "HeapInuse",
			MType: "gauge",
			Value: getPointerFloat(float64(m.RuntimeMem.HeapInuse)),
		},
		{
			ID:    "RandomValue",
			MType: "gauge",
			Value: getPointerFloat(m.RandomValue),
		},
		{
			ID:    "PollCount",
			MType: "counter",
			Delta: getPointerInt(m.PollCount),
		},
	}

	metrics, err := json.Marshal(req2)
	if err != nil {
		return err
	}

	// Creating a new gzip writer
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	defer gz.Close()

	// Writing the body to the gzip writer
	if _, err = gz.Write(metrics); err != nil {
		return err
	}

	if err = gz.Flush(); err != nil {
		return err
	}

	// Sending request with compressed data
	client := http.Client{Timeout: 5 * time.Second}
	reqw, err := http.NewRequest("POST", fmt.Sprintf(AllMetricsURL, addr), &buf)
	if err != nil {
		return err
	}

	if err = cliSend(client, reqw); err != nil {
		sec := 3
		tickerResend := time.NewTicker(time.Duration(sec) * time.Second)
		for {
			select {
			case <-ctx.Done():
				return fmt.Errorf("context canceled")
			case <-tickerResend.C:
				err := cliSend(client, reqw)
				if sec == 7 && err != nil {
					return err
				}
				if err == nil {
					return nil
				}

				sec += 2
			}
		}
	}

	return nil
}

func cliSend(client http.Client, reqw *http.Request) error {
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
