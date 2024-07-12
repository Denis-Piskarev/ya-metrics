package memyandex

import (
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemStatsYaSt_UpdateMetrics(t *testing.T) {
	m := MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	ctx := context.Background()

	m.UpdateMetrics(ctx)

	// Проверка что хотя бы 1 переменная обновилась
	assert.NotEmpty(t, m.RuntimeMem.Alloc)
}
