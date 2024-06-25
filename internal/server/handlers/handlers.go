package handlers

import (
	"net/http"

	"github.com/DenisquaP/ya-metrics/internal/server/middlewares"
	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type Handler struct {
	Metrics *yametrics.MemStorage
}

func NewHandler() *Handler {
	metrics := yametrics.NewMemStorage()
	return &Handler{
		Metrics: metrics,
	}
}

func InitRouter(logger zap.SugaredLogger) http.Handler {
	r := chi.NewRouter()
	r.Use(middlewares.Logging(logger))

	h := NewHandler()

	r.Get("/", h.GetMetrics)

	r.Route("/", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))

		r.Post("/update", h.createMetric)

		r.Post("/", h.GetMetric)
	})

	return r
}
