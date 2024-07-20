package app

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"

	"github.com/DenisquaP/ya-metrics/internal/server/config"
	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
)

// WriteFile writes metrics to file with duration
func writeFile(ctx context.Context, logger *zap.SugaredLogger, metrics yametrics.Metric, cfg config.Config) {
	// Getting working directory
	wd, err := os.Getwd()
	if err != nil {
		logger.Fatalw("Failed to get working directory", "error", err)
	}

	// If restore is true restore metrics from file
	if cfg.Restore {
		logger.Info("Restore metrics from file")
		if err := metrics.RestoreFromFile(wd); err != nil {
			logger.Errorw("Failed to restore metrics from file", "error", err)
		}
	}

	if cfg.FileStoragePath == "" {
		return
	}

	// Getting path to folder with metrics
	ds, _ := filepath.Split(cfg.FileStoragePath)
	// Creating folder with metrics
	if err := os.MkdirAll(filepath.Join(wd, ds), 0777); err != nil {
		logger.Fatalw("Failed to create working directory", "error", errors.Unwrap(err))
	}

	for {
		select {
		case <-ctx.Done():
			logger.Info("Save metrics to file")
			metrics.SaveMetricsToFile(wd)
			return
		default:
			// Saving metrics to file with interval
			wTO, cancel := context.WithTimeout(ctx, time.Duration(cfg.StoreInterval)*time.Second)
			defer cancel()
			<-wTO.Done()

			if err := metrics.SaveMetricsToFile(wd); err != nil {
				logger.Errorw("Failed to save metrics to file", "error", err)
			}
		}
	}
}
