package handlers

import (
	"github.com/DenisquaP/ya-metrics/internal/yaMetrics"
	"net/http"
	"strings"
)

func createMetric(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	if r.Method != "POST" {
		http.Error(rw, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	url, ok := strings.CutPrefix(r.URL.String(), "/update/")
	if !ok {
		http.Error(rw, "empty yaMetrics", http.StatusBadRequest)
		return
	}

	sliceMetric := strings.Split(url, "/")
	if len(sliceMetric) != 3 {
		http.Error(rw, "empty yaMetrics", http.StatusNotFound)
		return
	}

	typeMetric := sliceMetric[0]
	nameMetric := sliceMetric[1]
	valueMetric := sliceMetric[2]

	err := yaMetric.WriteMetric(nameMetric, typeMetric, valueMetric)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
