package main

import (
	"github.com/DenisquaP/ya-metrics/internal/agent/memyandex"
	"log"
	"runtime"
)

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	mem := memyandex.MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	mem.UpdateMetrics(cfg.pollInterval)
	mem.SendToServer(cfg.runAddr, cfg.reportInterval)
}
