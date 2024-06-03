package handlers

import (
	"net/http"

	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	Metrics *yametrics.MemStorage
}

func NewHanler() *Handler {
	metrics := yametrics.NewMemStorage()
	return &Handler{
		Metrics: metrics,
	}
}

func InitRouter() http.Handler {
	r := chi.NewRouter()

	h := NewHanler()

	r.Get("/", h.GetMetrics)

	r.Route("/update", func(r chi.Router) {
		r.Use(middleware.AllowContentType("text/plain"))

		r.Post("/{type}/{name}/{value}", h.createMetric)
	})

	r.Get("/value/{type}/{name}", h.GetMetric)

	return r
}
