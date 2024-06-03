package main

import (
	"github.com/DenisquaP/ya-metrics/internal/agent/memyandex"
	"runtime"
)

func main() {
	parseFlags()
	mem := memyandex.MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	mem.UpdateMetrics(flagPollInterval)
	mem.SendToServer(flagRunAddr, flagReportInterval)
}
