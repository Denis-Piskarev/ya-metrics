package handlers

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
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
			url:          "/update/counters/Met/2",
			method:       "POST",
			expectedCode: http.StatusBadRequest,
		},
	}
	srv := httptest.NewServer(InitRouter())
	defer srv.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tt.method
			req.URL = srv.URL + tt.url

			resp, err := req.Send()
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode())
		})
	}
}
