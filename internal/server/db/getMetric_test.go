package db

import (
	"context"
	"errors"
	"testing"

	"github.com/DenisquaP/ya-metrics/internal/server/db/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetCounterOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mocks.NewMockDBInterface(ctrl)
	ctx := context.Background()

	db.EXPECT().GetCounter(ctx, "test").Return(int64(1), nil)

	got, err := db.GetCounter(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, int64(1), got)
}

func TestGetCounterErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mocks.NewMockDBInterface(ctrl)
	ctx := context.Background()

	errExp := errors.New("error")

	db.EXPECT().GetCounter(ctx, "test").Return(int64(0), errExp)

	got, err := db.GetCounter(ctx, "test")
	require.Error(t, err)
	require.EqualError(t, err, errExp.Error())
	require.Equal(t, int64(0), got)
}

func TestGetGaugeOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mocks.NewMockDBInterface(ctrl)
	ctx := context.Background()

	db.EXPECT().GetGauge(ctx, "test").Return(float64(1), nil)

	got, err := db.GetGauge(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, float64(1), got)
}

func TestGetGaugeErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := mocks.NewMockDBInterface(ctrl)
	ctx := context.Background()

	errExp := errors.New("error")

	db.EXPECT().GetGauge(ctx, "test").Return(float64(0), errExp)

	got, err := db.GetGauge(ctx, "test")
	require.Error(t, err)
	require.EqualError(t, err, errExp.Error())
	require.Equal(t, float64(0), got)
}
