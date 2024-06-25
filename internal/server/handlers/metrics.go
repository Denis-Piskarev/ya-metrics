package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/DenisquaP/ya-metrics/pkg/models"
	"github.com/go-chi/chi"
)

func (h *Handler) createMetric(rw http.ResponseWriter, r *http.Request) {
	typeMetric := chi.URLParam(r, "type")
	nameMetric := chi.URLParam(r, "name")
	valueMetric := chi.URLParam(r, "value")

	if nameMetric == "" {
		http.Error(rw, "empty name", http.StatusNotFound)
		return
	}

	switch typeMetric {
	case "counter":
		val, err := strconv.ParseInt(valueMetric, 10, 64)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.Metrics.WriteCounter(nameMetric, val); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

	case "gauge":
		val, err := strconv.ParseFloat(valueMetric, 64)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.Metrics.WriteGauge(nameMetric, val); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(rw, "wrong type", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMetric(rw http.ResponseWriter, r *http.Request) {
	typeMet := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")

	switch typeMet {
	case "counter":
		val, err := h.Metrics.GetCounter(name)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}

		rw.Write([]byte(strconv.Itoa(int(val))))
	case "gauge":
		val, err := h.Metrics.GetGauge(name)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}

		resp, err := json.Marshal(val)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Write(resp)
	default:
		http.Error(rw, "wrong type", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) createMetricV2(rw http.ResponseWriter, r *http.Request) {
	var request models.Metrics

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if request.ID == "" {
		http.Error(rw, "empty name", http.StatusNotFound)
		return
	}

	switch request.MType {
	case "counter":
		old, err := h.Metrics.WriteCounter(request.ID, *request.Delta)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		request.Delta = &old

	case "gauge":
		old, err := h.Metrics.WriteGauge(request.ID, *request.Value)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		request.Value = &old
	default:
		http.Error(rw, "wrong type", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(request)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(resp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMetricV2(rw http.ResponseWriter, r *http.Request) {
	var request models.Metrics
	var response models.Metrics

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	switch request.MType {
	case "counter":
		c, err := h.Metrics.GetCounter(request.ID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}
		response.ID = request.ID
		response.MType = request.MType
		response.Delta = &c
	case "gauge":
		g, err := h.Metrics.GetGauge(request.ID)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}
		response.ID = request.ID
		response.MType = request.MType
		response.Value = &g
	default:
		http.Error(rw, "wrong type", http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	if _, err = rw.Write(res); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetMetrics(rw http.ResponseWriter, r *http.Request) {
	metrics := h.Metrics.GetMetrics()

	metHTML := strings.Replace(HTMLMet, "{{metrics}}", metrics, -1)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(metHTML))
}

var HTMLMet = `<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Metrics</title>
</head>

<body>
    Metrics:
    {{metrics}}
</body>

</html>`
