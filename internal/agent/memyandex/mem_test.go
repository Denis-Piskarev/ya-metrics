package memyandex

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestMemStatsYaSt_UpdateMetrics(t *testing.T) {
	m := MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	m.UpdateMetrics()

	// Проверка что хотя бы 1 переменная обновилась
	assert.NotEmpty(t, m.RuntimeMem.Alloc)
}
