package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) createMetric(rw http.ResponseWriter, r *http.Request) {
	typeMetric := chi.URLParam(r, "type")
	nameMetric := chi.URLParam(r, "name")
	valueMetric := chi.URLParam(r, "value")

	log.Println(typeMetric, nameMetric, valueMetric)

	if typeMetric == "counter" {
		err := h.Metrics.WriteCounter(nameMetric, valueMetric)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	} else if typeMetric == "gauge" {
		err := h.Metrics.WriteGauge(nameMetric, valueMetric)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (h *Handler) GetMetric(rw http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	val, err := h.Metrics.GetGauge(name)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	rw.Write([]byte(val))
}

func (h *Handler) GetMetrics(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(h.Metrics.GetMetrics()))
}
