package yametrics

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteGauge(t *testing.T) {
	metric := NewMemStorage("")

	val, err := metric.WriteGauge(context.Background(), "test", 1.0)
	require.NoError(t, err)

	require.Equal(t, 1.0, val)
}

func TestWriteCountrer(t *testing.T) {
	metric := NewMemStorage("")

	val, err := metric.WriteCounter(context.Background(), "test", 1)
	require.NoError(t, err)

	require.Equal(t, int64(1), val)
}

func TestGetCounter(t *testing.T) {
	metric := NewMemStorage("")

	val, err := metric.WriteCounter(context.Background(), "test", 1)
	require.NoError(t, err)

	getVal, err := metric.GetCounter(context.Background(), "test")
	require.NoError(t, err)
	require.Equal(t, val, getVal)
}

func TestGetGauge(t *testing.T) {
	metric := NewMemStorage("")

	val, err := metric.WriteGauge(context.Background(), "test", 1.2)
	require.NoError(t, err)

	getVal, err := metric.GetGauge(context.Background(), "test")
	require.NoError(t, err)
	require.Equal(t, val, getVal)
}
