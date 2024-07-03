package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.DB.Ping(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		h.Logger.Errorw("ping error", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok"}`)
}
