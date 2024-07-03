package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	"github.com/DenisquaP/ya-metrics/internal/server/db"
	"github.com/DenisquaP/ya-metrics/internal/server/middlewares"
	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
)

type Handler struct {
	Metrics yametrics.Metric
	Logger  *zap.SugaredLogger
	DB      db.DbInterface
}

func NewHandler(metrics yametrics.Metric, logger *zap.SugaredLogger, db db.DbInterface) *Handler {
	return &Handler{
		Metrics: metrics,
		Logger:  logger,
		DB:      db,
	}
}

func NewRouterWithMiddlewares(ctx context.Context, logger *zap.SugaredLogger, metrics yametrics.Metric, db db.DbInterface) http.Handler {
	select {
	case <-ctx.Done():
		logger.Errorw("context canceled", "error", ctx.Err())
		return nil
	default:
	}

	r := chi.NewRouter()

	r.Use(middlewares.Logging(logger))

	// Middleware for comporession
	r.Use(middlewares.Compression)

	h := NewHandler(metrics, logger, db)

	// To get all metrics in HTML
	r.Get("/", h.GetMetrics)

	// Ping database
	r.Get("/ping", h.Ping)

	r.Route("/", func(r chi.Router) {
		// Middleware for check ContentType
		r.Use(middleware.AllowContentType("application/json"))

		// Update metric
		r.Post("/update/{type}/{name}/{value}", h.createMetric)

		// Update metric JSON
		r.Post("/update/", h.createMetricJSON)

		// Get metric JSON
		r.Post("/value/", h.GetMetricJSON)
	})

	// Получение метрик v1
	r.Get("/value/{type}/{name}", h.GetMetric)

	return r
}
