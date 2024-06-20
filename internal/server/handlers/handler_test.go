package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitHandlers(t *testing.T) {
	r := InitRouter()

	assert.NotEmpty(t, r)
}
