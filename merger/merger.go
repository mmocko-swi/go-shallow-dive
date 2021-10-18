package merger

import "solarwinds/pisigma/analyzer/analyzer"

func Merge(leftResults *[]analyzer.ResultText, rightResults *[]analyzer.ResultText) []analyzer.ResultText {
	finalMap := make(map[string]int)

	if leftResults != nil {
		for _, value := range *leftResults {
			finalMap[value.Word] += value.Count
		}
	}

	if rightResults != nil {
		for _, value := range *rightResults {
			finalMap[value.Word] += value.Count
		}
	}

	finalResults := make([]analyzer.ResultText, 0, len(finalMap))
	for key, value := range finalMap {
		finalResults = append(finalResults, analyzer.ResultText{Word: key, Count: value})
	}
	return finalResults
}
