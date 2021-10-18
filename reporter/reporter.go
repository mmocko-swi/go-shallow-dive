package reporter

import "solarwinds/pisigma/analyzer/analyzer"

type Reporter interface {
	Report(outputFile string, results *[]analyzer.ResultText)
}
