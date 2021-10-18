package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"solarwinds/pisigma/analyzer/analyzer"
)

type jsonReporter struct {
}

func NewJsonReporter() Reporter {
	return &jsonReporter{}
}

func (reporter *jsonReporter) Report(outputFile string, results *[]analyzer.ResultText) {
	js, _ := json.Marshal(results)
	f, _ := os.OpenFile(outputFile, os.O_CREATE, 0777)
	defer f.Close()

	f.Write(js)

	fmt.Printf("Report written into: %s \n", outputFile)
}
