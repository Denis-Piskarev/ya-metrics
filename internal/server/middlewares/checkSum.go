package middlewares

import (
	"github.com/DenisquaP/ya-metrics/internal/cryptography"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func GetSum(logger *zap.SugaredLogger, key string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			metrics, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Errorw("error reading body", "error", err)
			}

			sumGet := r.Header.Get("HashSHA256")
			if sumGet == "" {
				logger.Warnw("Missing hash SHA256 header")
				w.WriteHeader(http.StatusBadRequest)

				return
			}

			expectedSum := cryptography.GetSum(metrics, key)

			if sumGet != expectedSum {
				logger.Warnw("Expected hash does not match", "expected", expectedSum, "actual", sumGet)

				w.WriteHeader(http.StatusBadRequest)
			}

			next.ServeHTTP(w, r)
		})
	}
}
