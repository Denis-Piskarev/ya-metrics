package app

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/DenisquaP/ya-metrics/internal/agent/config"
	"github.com/DenisquaP/ya-metrics/internal/agent/memyandex"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Creating struct for collecting metrics
	mem := memyandex.MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	ctx := context.Background()

	ticker := time.NewTicker(time.Duration(cfg.PollInterval) * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			mem.SendAllMetricsToServer(ctx)
			if err := mem.SendToServer(ctx, cfg.RunAddr, cfg.ReportInterval); err != nil {
				log.Printf("error send metrics: %s", err)
			}
		}

	}
}
