package textprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess_WhenTextIsGive_ThenCorrectResultIsReturned(t *testing.T) {
	originalText := "a b  c   d	d"

	processor := NewProcessor()
	stats := processor.Process(originalText)

	assert.True(t, len(stats.Stats) == 4, "Expected stats to have 4 elements")
	assert.Equal(t, 1, stats.Stats["a"])
	assert.Equal(t, 1, stats.Stats["b"])
	assert.Equal(t, 1, stats.Stats["b"])
	assert.Equal(t, 2, stats.Stats["d"])
}
