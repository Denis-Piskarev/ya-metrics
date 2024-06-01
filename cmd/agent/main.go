package main

import (
	"github.com/DenisquaP/ya-metrics/internal/agent/memyandex"
	"runtime"
)

func main() {
	mem := memyandex.MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	mem.UpdateMetrics()
	mem.SendToServer()
}
