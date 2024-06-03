package memyandex

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

func TestMemStatsYaSt_UpdateMetrics(t *testing.T) {
	m := MemStatsYaSt{RuntimeMem: &runtime.MemStats{}}

	m.UpdateMetrics(2 * time.Second)

	// Проверка что хотя бы 1 переменная обновилась
	assert.NotEmpty(t, m.RuntimeMem.Alloc)
}
