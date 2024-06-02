package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMetrics(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		method       string
		expectedCode int
	}{
		{
			name:         "POST 200",
			url:          "/update/counter/Met/2",
			method:       "POST",
			expectedCode: http.StatusOK,
		},
		{
			name:         "GET 405",
			url:          "/update/counter/Met/2",
			method:       "GET",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			name:         "POST 404",
			url:          "/update/Met/2",
			method:       "POST",
			expectedCode: http.StatusNotFound,
		}, {
			name:         "POST 400",
			url:          "/set/counter/Met/2",
			method:       "POST",
			expectedCode: http.StatusBadRequest,
		},
	}
	h := NewHanler()
	srv := httptest.NewServer(InitRouter())
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, tt.url, nil)
			req.Header.Set("Content-Type", "plain/text")

			h.createMetric(r, req)
			assert.Equal(t, tt.expectedCode, r.Code)
		})
	}

}
