package analyzer

import (
	"solarwinds/pisigma/analyzer/textprocessor"
	"sort"
)

type ResultText struct {
	word    string
	counter int
}

type keyValuePair struct {
	word    string
	counter int
}

func Analyze(stats textprocessor.TextStatistics) *[]ResultText {

	var sortedSlice []keyValuePair
	for key, value := range stats.Stats {
		sortedSlice = append(sortedSlice, keyValuePair{key, value})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].counter > sortedSlice[j].counter
	})

	var ln = len(sortedSlice)
	if ln > 50 {
		ln = 50
	}

	var results []ResultText
	for _, word := range sortedSlice[:ln] {
		results = append(results, ResultText{word.word, word.counter})
	}

	return &results
}
