package handlers

import (
	"net/http"
	"os"
	"strings"

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

	if err := h.Metrics.WriteMetric(nameMetric, typeMetric, valueMetric); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMetric(rw http.ResponseWriter, r *http.Request) {
	typeMet := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")

	val, err := h.Metrics.GetMetric(typeMet, name)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	file, err := os.Open("metrics.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	defer file.Close()

	metHtml := strings.Replace(htmlMet, "{{metrics}}", val, -1)

	rw.Write([]byte(metHtml))
}

func (h *Handler) GetMetrics(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(h.Metrics.GetMetrics()))
}

var htmlMet = `<!DOCTYPE html>
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
