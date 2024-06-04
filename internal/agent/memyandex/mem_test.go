package memyandex

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemStatsYaSt_UpdateMetrics(t *testing.T) {
	m := MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	m.UpdateMetrics(2)

	// Проверка что хотя бы 1 переменная обновилась
	assert.NotEmpty(t, m.RuntimeMem.Alloc)
}
