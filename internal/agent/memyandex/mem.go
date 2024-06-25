package memyandex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

const URL = "http://%s/update"

type MemStatsYaSt struct {
	RuntimeMem  *runtime.MemStats
	PollCount   int64
	RandomValue float64
}

func (m *MemStatsYaSt) UpdateMetrics(ctx context.Context, pollInterval int) {
	runtime.ReadMemStats(m.RuntimeMem)
	m.RandomValue = float64(m.RuntimeMem.Alloc) / float64(1024)
	m.PollCount++

	withTimeout, cancel := context.WithTimeout(ctx, time.Duration(pollInterval)*time.Second)
	defer cancel()
	<-withTimeout.Done()
}

func (m *MemStatsYaSt) SendToServer(ctx context.Context, runAddr string, reportInterval int) {
	// Alloc
	sendGauge(float64(m.RuntimeMem.Alloc), "Alloc", "gauge", runAddr)

	// BuckHashSys
	sendGauge(float64(m.RuntimeMem.BuckHashSys), "BuckHashSys", "gauge", runAddr)

	// Frees
	sendGauge(float64(m.RuntimeMem.Frees), "Frees", "gauge", runAddr)

	// GCCPUFraction
	sendGauge(m.RuntimeMem.GCCPUFraction, "GCCPUFraction", "gauge", runAddr)

	// GCSys
	sendGauge(float64(m.RuntimeMem.GCSys), "GCSys", "gauge", runAddr)

	// HeapAlloc
	sendGauge(float64(m.RuntimeMem.HeapAlloc), "HeapAlloc", "gauge", runAddr)

	// HeapIdle
	sendGauge(float64(m.RuntimeMem.HeapIdle), "HeapIdle", "gauge", runAddr)

	// HeapObjects
	sendGauge(float64(m.RuntimeMem.HeapObjects), "HeapObjects", "gauge", runAddr)

	// HeapReleased
	sendGauge(float64(m.RuntimeMem.HeapReleased), "HeapReleased", "gauge", runAddr)

	// HeapSys
	sendGauge(float64(m.RuntimeMem.HeapSys), "HeapSys", "gauge", runAddr)

	// LastGC
	sendGauge(float64(m.RuntimeMem.LastGC), "LastGC", "gauge", runAddr)

	// Lookups
	sendGauge(float64(m.RuntimeMem.Lookups), "Lookups", "gauge", runAddr)

	// MCacheInuse
	sendGauge(float64(m.RuntimeMem.MCacheInuse), "MCacheInuse", "gauge", runAddr)

	// MCacheSys
	sendGauge(float64(m.RuntimeMem.MCacheSys), "MCacheSys", "gauge", runAddr)

	// MSpanInuse
	sendGauge(float64(m.RuntimeMem.MSpanInuse), "MSpanInuse", "gauge", runAddr)

	// MSpanSys
	sendGauge(float64(m.RuntimeMem.MSpanSys), "MSpanSys", "gauge", runAddr)

	// Mallocs
	sendGauge(float64(m.RuntimeMem.Mallocs), "Mallocs", "gauge", runAddr)

	// NextGC
	sendGauge(float64(m.RuntimeMem.NextGC), "NextGC", "gauge", runAddr)

	// NumForcedGC
	sendGauge(float64(m.RuntimeMem.NumForcedGC), "NumForcedGC", "gauge", runAddr)

	// NumGC
	sendGauge(float64(m.RuntimeMem.NumGC), "NumGC", "gauge", runAddr)

	// OtherSys
	sendGauge(float64(m.RuntimeMem.OtherSys), "OtherSys", "gauge", runAddr)

	// PauseTotalNs
	sendGauge(float64(m.RuntimeMem.PauseTotalNs), "PauseTotalNs", "gauge", runAddr)

	// StackInuse
	sendGauge(float64(m.RuntimeMem.StackInuse), "StackInuse", "gauge", runAddr)

	// Sys
	sendGauge(float64(m.RuntimeMem.Sys), "Sys", "gauge", runAddr)

	// TotalAlloc
	sendGauge(float64(m.RuntimeMem.TotalAlloc), "TotalAlloc", "gauge", runAddr)

	// PollCount
	sendCounter(m.PollCount, "PollCount", "counter", runAddr)

	// RandomValue
	sendGauge(m.RandomValue, "RandomValue", "gauge", runAddr)

	WithTimeout, cancel := context.WithTimeout(ctx, time.Duration(reportInterval)*time.Second)
	defer cancel()

	<-WithTimeout.Done()
}

// For sending gouge metric to server
func sendGauge(variable float64, name, vType, addr string) {
	req := models.Metrics{
		ID:    name,
		Value: &variable,
		MType: vType,
	}

	body, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(fmt.Sprintf(URL, addr), "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}
}

// For sending counter metric to server
func sendCounter(variable int64, name, vType, addr string) {
	req := models.Metrics{
		ID:    name,
		Delta: &variable,
		MType: vType,
	}

	body, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(fmt.Sprintf(URL, addr), "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}
}
