package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitHandlers(t *testing.T) {
	mux := InitHandlers()

	assert.NotEmpty(t, mux)
}
