package middlewares

import (
	"net/http"
	"strings"

	"github.com/DenisquaP/ya-metrics/internal/server/compression"
)

func Commpression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := w

		// Если принимаем gzip то сжимаем
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			cw := compression.NewCompressWriter(w)
			rw = cw
			
			defer cw.Close()
		}

		// Если содержимое содержит gzip
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			cr, err := compression.NewCompressReader(r.Body)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			r.Body = cr
			defer cr.Close()
		}

		next.ServeHTTP(rw, r)
	})
}
