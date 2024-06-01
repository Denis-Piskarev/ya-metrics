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
	gauge := float64(m.RuntimeMem.Alloc)
	res, err := http.Post(fmt.Sprintf(URL+"gauge/Alloc/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// BuckHashSys
	gauge = float64(m.RuntimeMem.BuckHashSys)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/BuckHashSys/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// Frees
	gauge = float64(m.RuntimeMem.Frees)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/Frees/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// GCCPUFraction
	gauge = m.RuntimeMem.GCCPUFraction
	res, err = http.Post(fmt.Sprintf(URL+"gauge/GCCPUFraction/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// GCSys
	gauge = float64(m.RuntimeMem.GCSys)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/GCSys/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// HeapAlloc
	gauge = float64(m.RuntimeMem.HeapAlloc)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/HeapAlloc/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// HeapIdle
	gauge = float64(m.RuntimeMem.HeapIdle)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/HeapIdle/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// HeapInuse
	gauge = float64(m.RuntimeMem.HeapInuse)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/HeapInuse/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// HeapObjects
	gauge = float64(m.RuntimeMem.HeapObjects)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/HeapObjects/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// HeapReleased
	gauge = float64(m.RuntimeMem.HeapReleased)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/HeapReleased/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// HeapSys
	gauge = float64(m.RuntimeMem.HeapSys)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/HeapSys/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// LastGC
	gauge = float64(m.RuntimeMem.LastGC)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/LastGC/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// Lookups
	gauge = float64(m.RuntimeMem.Lookups)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/Lookups/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// MCacheInuse
	gauge = float64(m.RuntimeMem.MCacheInuse)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/MCacheInuse/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// MSpanSys
	gauge = float64(m.RuntimeMem.MSpanSys)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/MSpanSys/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// Mallocs
	gauge = float64(m.RuntimeMem.Mallocs)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/Mallocs/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// NextGC
	gauge = float64(m.RuntimeMem.NextGC)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/NextGC/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// NumForcedGC
	gauge = float64(m.RuntimeMem.NumForcedGC)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/NumForcedGC/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// NumGC
	gauge = float64(m.RuntimeMem.NumGC)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/NumGC/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// OtherSys
	gauge = float64(m.RuntimeMem.OtherSys)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/OtherSys/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// PauseTotalNs
	gauge = float64(m.RuntimeMem.PauseTotalNs)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/PauseTotalNs/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// StackInuse
	gauge = float64(m.RuntimeMem.StackInuse)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/StackInuse/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// StackSys
	gauge = float64(m.RuntimeMem.StackSys)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/StackSys/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// Sys
	gauge = float64(m.RuntimeMem.Sys)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/Sys/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// TotalAlloc
	gauge = float64(m.RuntimeMem.TotalAlloc)
	res, err = http.Post(fmt.Sprintf(URL+"gauge/TotalAlloc/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// PollCount
	counter := m.PollCount
	res, err = http.Post(fmt.Sprintf(URL+"counter/PollCount/%d", counter), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	// RandomValue
	gauge = m.RandomValue
	res, err = http.Post(fmt.Sprintf(URL+"gauge/RandomValue/%v", gauge), "text/plain", nil)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d", res.StatusCode)
	}

	time.Sleep(10 * time.Second)
}
