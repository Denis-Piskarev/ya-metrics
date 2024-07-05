package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/DenisquaP/ya-metrics/internal/server"
	"github.com/DenisquaP/ya-metrics/internal/server/db"
	"github.com/DenisquaP/ya-metrics/internal/server/handlers"
	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
	"go.uber.org/zap"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	sugared := logger.Sugar()

	// Initiating config
	cfg, err := server.NewConfig()
	if err != nil {
		sugared.Fatalw("Failed to parse config", "error", err)
	}

	sugared.Infow("Starting server", "address", cfg.RunAddr)

	// Initiating router
	metrics := yametrics.NewMemStorage(cfg.FileStoragePath)

	// Initiating DB
	db, err := db.NewDB(ctx, cfg.DatabaseDsn)
	if err != nil {
		sugared.Fatalw("Failed to create new DB", "error", err)
	}
	defer db.DB.Close()

	router := handlers.NewRouterWithMiddlewares(ctx, sugared, metrics, db)

	go writeFile(ctx, sugared, metrics, cfg)

	if err := http.ListenAndServe(cfg.RunAddr, router); err != nil {
		sugared.Fatalw("Failed to start server", "error", err)
	}
}

// WriteFile writes metrics to file with duration
func writeFile(ctx context.Context, logger *zap.SugaredLogger, metrics yametrics.Metric, cfg server.Config) {
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
