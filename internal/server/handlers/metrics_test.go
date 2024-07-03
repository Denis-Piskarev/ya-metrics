package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/DenisquaP/ya-metrics/internal/server/db/mocks"
	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
)

func TestCreateMetrics(t *testing.T) {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer logger.Sync()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mocks.NewMockDbInterface(ctrl)

	suggared := logger.Sugar()

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

	mem := yametrics.NewMemStorage("mem.json")
	srv := httptest.NewServer(NewRouterWithMiddlewares(ctx, suggared, mem, db))
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
