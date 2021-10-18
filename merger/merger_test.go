package merger

import (
	"solarwinds/pisigma/analyzer/analyzer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNilIsGiven_ThenMergeIsSafeForUse(t *testing.T) {
	arg := []analyzer.ResultText{}

	results := Merge(nil, &arg)
	assert.Equal(t, 0, len(results))

	results = Merge(&arg, nil)

	assert.Equal(t, 0, len(results))
}

func TestWhenResultsAreGiven_ThenMergeIsPerformed(t *testing.T) {

	rightResults := []analyzer.ResultText{
		{Word: "a", Count: 1},
		{Word: "b", Count: 2},
	}

	leftResults := []analyzer.ResultText{
		{Word: "a", Count: 1},
		{Word: "c", Count: 2},
	}

	result := Merge(&rightResults, &leftResults)

	assert.Equal(t, len(result), 3)
}
