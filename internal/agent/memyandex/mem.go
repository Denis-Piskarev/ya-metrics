package memyandex

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

const URL = "http://localhost%s/update/"

type MemStatsYaSt struct {
	RuntimeMem  *runtime.MemStats
	PollCount   int64
	RandomValue float64
}

func (m *MemStatsYaSt) UpdateMetrics(pollInterval time.Duration) {
	runtime.ReadMemStats(m.RuntimeMem)
	m.RandomValue = float64(m.RuntimeMem.Alloc) / float64(1024)
	m.PollCount++

	time.Sleep(pollInterval)
}

func (m *MemStatsYaSt) SendToServer(runAddr string, reportInterval time.Duration) {
	// Alloc
	sent(float64(m.RuntimeMem.Alloc), "Alloc", "gauge", runAddr)

	// BuckHashSys
	sent(float64(m.RuntimeMem.BuckHashSys), "BuckHashSys", "gauge", runAddr)

	// Frees
	sent(float64(m.RuntimeMem.Frees), "Frees", "gauge", runAddr)

	// GCCPUFraction
	sent(m.RuntimeMem.GCCPUFraction, "GCCPUFraction", "gauge", runAddr)

	// GCSys
	sent(float64(m.RuntimeMem.GCSys), "GCSys", "gauge", runAddr)

	// HeapAlloc
	sent(float64(m.RuntimeMem.HeapAlloc), "HeapAlloc", "gauge", runAddr)

	// HeapIdle
	sent(float64(m.RuntimeMem.HeapIdle), "HeapIdle", "gauge", runAddr)

	// HeapObjects
	sent(float64(m.RuntimeMem.HeapObjects), "HeapObjects", "gauge", runAddr)

	// HeapReleased
	sent(float64(m.RuntimeMem.HeapReleased), "HeapReleased", "gauge", runAddr)

	// HeapSys
	sent(float64(m.RuntimeMem.HeapSys), "HeapSys", "gauge", runAddr)

	// LastGC
	sent(float64(m.RuntimeMem.LastGC), "LastGC", "gauge", runAddr)

	// Lookups
	sent(float64(m.RuntimeMem.Lookups), "Lookups", "gauge", runAddr)

	// MCacheInuse
	sent(float64(m.RuntimeMem.MCacheInuse), "MCacheInuse", "gauge", runAddr)

	// MCacheSys
	sent(float64(m.RuntimeMem.MCacheSys), "MCacheSys", "gauge", runAddr)

	// MSpanInuse
	sent(float64(m.RuntimeMem.MSpanInuse), "MSpanInuse", "gauge", runAddr)

	// MSpanSys
	sent(float64(m.RuntimeMem.MSpanSys), "MSpanSys", "gauge", runAddr)

	// Mallocs
	sent(float64(m.RuntimeMem.Mallocs), "Mallocs", "gauge", runAddr)

	// NextGC
	sent(float64(m.RuntimeMem.NextGC), "NextGC", "gauge", runAddr)

	// NumForcedGC
	sent(float64(m.RuntimeMem.NumForcedGC), "NumForcedGC", "gauge", runAddr)

	// NumGC
	sent(float64(m.RuntimeMem.NumGC), "NumGC", "gauge", runAddr)

	// OtherSys
	sent(float64(m.RuntimeMem.OtherSys), "OtherSys", "gauge", runAddr)

	// PauseTotalNs
	sent(float64(m.RuntimeMem.PauseTotalNs), "PauseTotalNs", "gauge", runAddr)

	// StackInuse
	sent(float64(m.RuntimeMem.StackInuse), "StackInuse", "gauge", runAddr)

	// Sys
	sent(float64(m.RuntimeMem.Sys), "Sys", "gauge", runAddr)

	// TotalAlloc
	sent(float64(m.RuntimeMem.TotalAlloc), "TotalAlloc", "gauge", runAddr)

	// PollCount
	sent(float64(m.PollCount), "PollCount", "counter", runAddr)

	// RandomValue
	sent(m.RandomValue, "RandomValue", "gauge", runAddr)

	time.Sleep(reportInterval)
}

func sent(variable any, name, vType, addr string) {
	res, err := http.Post(fmt.Sprintf(URL+"%s/%s/%v", addr, vType, name, variable), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}
}
