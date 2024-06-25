package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/DenisquaP/ya-metrics/pkg/models"
)

func (h *Handler) createMetric(rw http.ResponseWriter, r *http.Request) {
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
		if err := h.Metrics.WriteCounter(request.ID, *request.Delta); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

	case "gauge":
		if err := h.Metrics.WriteGauge(request.ID, *request.Value); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(rw, "wrong type", http.StatusNotFound)
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

func (h *Handler) GetMetric(rw http.ResponseWriter, r *http.Request) {
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
		http.Error(rw, "wrong type", http.StatusNotFound)
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
