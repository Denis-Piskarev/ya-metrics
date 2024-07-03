package handlers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/golang/mock/gomock"

	"github.com/DenisquaP/ya-metrics/internal/server/db/mocks"
	yametrics "github.com/DenisquaP/ya-metrics/internal/server/yaMetrics"
)

func TestInitHandlers(t *testing.T) {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)
	defer logger.Sync()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mocks.NewMockDbInterface(ctrl)

	suggared := logger.Sugar()
	mem := yametrics.NewMemStorage("mem.json")

	r := NewRouterWithMiddlewares(ctx, suggared, mem, db)

	assert.NotEmpty(t, r)
}
