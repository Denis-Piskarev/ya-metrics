package handlers

import "net/http"

func InitHandlers() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", createMetric)

	return mux
}
