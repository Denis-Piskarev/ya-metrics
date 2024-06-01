package memyandex

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

const URL = "http://localhost:8080/update/"

type MemStatsYaSt struct {
	RuntimeMem  *runtime.MemStats
	PollCount   int64
	RandomValue float64
}

func (m *MemStatsYaSt) UpdateMetrics() {
	runtime.ReadMemStats(m.RuntimeMem)
	m.RandomValue = float64(m.RuntimeMem.Alloc) / float64(1024)
	m.PollCount++
	time.Sleep(2 * time.Second)
}

func (m *MemStatsYaSt) SendToServer() {
	// Alloc
	sent(float64(m.RuntimeMem.Alloc), "Alloc", "gauge")

	// BuckHashSys
	sent(float64(m.RuntimeMem.BuckHashSys), "BuckHashSys", "gauge")

	// Frees
	sent(float64(m.RuntimeMem.Frees), "Frees", "gauge")

	// GCCPUFraction
	sent(m.RuntimeMem.GCCPUFraction, "GCCPUFraction", "gauge")

	// GCSys
	sent(float64(m.RuntimeMem.GCSys), "GCSys", "gauge")

	// HeapAlloc
	sent(float64(m.RuntimeMem.HeapAlloc), "HeapAlloc", "gauge")

	// HeapIdle
	sent(float64(m.RuntimeMem.HeapIdle), "HeapIdle", "gauge")

	// HeapObjects
	sent(float64(m.RuntimeMem.HeapObjects), "HeapObjects", "gauge")

	// HeapReleased
	sent(float64(m.RuntimeMem.HeapReleased), "HeapReleased", "gauge")

	// HeapSys
	sent(float64(m.RuntimeMem.HeapSys), "HeapSys", "gauge")

	// LastGC
	sent(float64(m.RuntimeMem.LastGC), "LastGC", "gauge")

	// Lookups
	sent(float64(m.RuntimeMem.Lookups), "Lookups", "gauge")

	// MCacheInuse
	sent(float64(m.RuntimeMem.MCacheInuse), "MCacheInuse", "gauge")

	// MCacheSys
	sent(float64(m.RuntimeMem.MCacheSys), "MCacheSys", "gauge")

	// MSpanInuse
	sent(float64(m.RuntimeMem.MSpanInuse), "MSpanInuse", "gauge")

	// MSpanSys
	sent(float64(m.RuntimeMem.MSpanSys), "MSpanSys", "gauge")

	// Mallocs
	sent(float64(m.RuntimeMem.Mallocs), "Mallocs", "gauge")

	// NextGC
	sent(float64(m.RuntimeMem.NextGC), "NextGC", "gauge")

	// NumForcedGC
	sent(float64(m.RuntimeMem.NumForcedGC), "NumForcedGC", "gauge")

	// NumGC
	sent(float64(m.RuntimeMem.NumGC), "NumGC", "gauge")

	// OtherSys
	sent(float64(m.RuntimeMem.OtherSys), "OtherSys", "gauge")

	// PauseTotalNs
	sent(float64(m.RuntimeMem.PauseTotalNs), "PauseTotalNs", "gauge")

	// StackInuse
	sent(float64(m.RuntimeMem.StackInuse), "StackInuse", "gauge")

	// Sys
	sent(float64(m.RuntimeMem.Sys), "Sys", "gauge")

	// TotalAlloc
	sent(float64(m.RuntimeMem.TotalAlloc), "TotalAlloc", "gauge")

	// PollCount
	sent(float64(m.PollCount), "PollCount", "counter")

	// RandomValue
	sent(m.RandomValue, "RandomValue", "gauge")

	time.Sleep(10 * time.Second)
}

func sent(variable any, name, vType string) {
	res, err := http.Post(fmt.Sprintf(URL+"%s/%s/%v", vType, name, variable), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}
}
