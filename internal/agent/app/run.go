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

	tickerSend := time.NewTicker(time.Duration(cfg.ReportInterval) * time.Second)
	tickerUpdate := time.NewTicker(time.Duration(cfg.PollInterval) * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-tickerSend.C:
			if err := mem.SendAllMetricsToServer(ctx, cfg.RunAddr); err != nil {
				log.Fatalf("error send metrics: %s", err)
			}

		case <-tickerUpdate.C:
			mem.UpdateMetrics(ctx)
		}

	}
}
