package analyzer

import (
	"fmt"
	"solarwinds/pisigma/analyzer/textprocessor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenMOreThan50StatsIsGiven_ThenOnlyTop50IsReturned(t *testing.T) {
	stats := make(map[string]int)
	for i := 0; i < 100; i++ {
		stats[fmt.Sprintf("%d", i)] = i
	}

	mockTextStats := textprocessor.TextStatistics{Stats: stats}
	result := Analyze(mockTextStats)

	assert.Equal(t, 50, len(*result))

	for index, element := range *result {
		assert.Equal(t, 99-index, element.Count)
	}
}

func Test_WhenLessThan50StatsIsGiven_ThenAllAreREturned(t *testing.T) {
	stats := make(map[string]int)
	for i := 0; i < 23; i++ {
		stats[fmt.Sprintf("%d", i)] = i
	}

	mockTextStats := textprocessor.TextStatistics{Stats: stats}
	result := Analyze(mockTextStats)

	assert.Equal(t, 23, len(*result))

	for index, element := range *result {
		assert.Equal(t, 22-index, element.Count)
	}
}
