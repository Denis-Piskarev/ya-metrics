package main

import (
	"github.com/DenisquaP/ya-metrics/internal/agent"
	"github.com/DenisquaP/ya-metrics/internal/agent/memyandex"
	"log"
	"runtime"
)

func main() {
	cfg, err := agent.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	mem := memyandex.MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	mem.UpdateMetrics(cfg.PollInterval)
	mem.SendToServer(cfg.RunAddr, cfg.ReportInterval)
}
